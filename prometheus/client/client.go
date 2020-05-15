package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ClareChu/gorequest"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type Prometheus interface {
	New(url string) *Client
	// 指标设置
	Metric(query string) *Client
	// 过滤对象
	Query(query interface{}) *Client
	//设置offset 修饰符
	OffSet(offSet string) *Client
	// 设置PromQL
	PromQL(ql string) *Client
}

var (
	DefaultTagKey = "pro"
	MaRegex       = "\\{[^\\}]+\\}"
)

type Client struct {
	Prometheus
	request *gorequest.SuperAgent
	metric  string
	start   float64
	end     float64
	step    int
	query   interface{}
	time    bool
	promeQL string
	offset  string
}

//prometheus response result
type PromeResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Result     []Result `json:"result"`
	ResultType string   `json:"resultType"`
}
type Result struct {
	Metric interface{}   `json:"metric"`
	Value  []interface{} `json:"value"`
}

var DefaultPrometheusHost int32 = 9090

func New(url string) *Client {
	request := gorequest.New()
	request.Url = url
	return &Client{
		request: request,
	}
}

//Query 指标名称
func (client *Client) Query(query interface{}) *Client {
	client.query = query
	return client
}

//Query 指标名称
func (client *Client) Metric(query string) *Client {
	client.metric = query
	return client
}

//Range 取值范围
func (client *Client) Range(start, end float64) *Client {
	client.start = start
	client.end = end
	client.time = true
	return client
}

//Range 取值范围
func (client *Client) Start(start float64) *Client {
	client.start = start
	return client
}

//Step 间隔时长
func (client *Client) Step(step int) *Client {
	client.step = step
	return client
}

func (client *Client) Run() (pro *PromeResponse, err error) {
	url := appendApiVersion(client.request.Url, client.time)
	url = client.appendParam(url)
	fmt.Println("request url:", url)
	resp, body, errs := client.request.Get(url).End()
	if len(errs) != 0 {
		return nil, errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(body)
	}
	pro = &PromeResponse{}
	err = json.Unmarshal([]byte(body), pro)
	if err != nil {
		return
	}
	return
}

func appendApiVersion(url string, time bool) string {
	if time {
		url = url + "/api/v1/query_range"
	}
	url = url + "/api/v1/query"
	return url
}

func (client *Client) appendParam(url string) string {
	now := time.Now().Unix()
	url = url + "?query=" + client.build()
	if client.time {
		url = fmt.Sprintf("%s&start=%f&end=%f", url, client.start, client.end)
	} else {
		if client.start == 0 {
			//default 设置 当前时间一分钟之内的
			client.start = float64(now - 60)
			url = fmt.Sprintf("%s&time=%f", url, client.start)
		}
	}
	url = fmt.Sprintf("%s&step=%d&_=%d", url, client.step, now)
	return url
}

// 拼接 promeql
func Where(o interface{}) string {
	sql := ""
	value := reflect.ValueOf(o)
	t := reflect.TypeOf(o)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
		t = t.Elem()
	}
	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).Kind() == reflect.String {
			tag := t.Field(i).Tag.Get(DefaultTagKey)
			if tag == "" {
				tag = t.Field(i).Name
			}
			if value.Field(i).String() != "" {
				r, err := regexp.Compile(MaRegex)
				if err != nil {
					return ""
				}
				if r.Match([]byte(sql)) {
					temp := fmt.Sprintf(",%s=\"%s\"}", tag, value.Field(i).String())
					sql = strings.Replace(sql, "}", temp, 1)
				} else {
					sql = fmt.Sprintf("{%s=\"%s\"}", tag, value.Field(i).String())
				}

			}
		}
	}
	fmt.Println("query :", sql)
	return sql
}

func (client *Client) PromQL(ql string) *Client {
	client.promeQL = ql
	return client
}

func (client *Client) OffSet(offSet string) *Client {
	client.offset = offSet
	return client
}

func (client *Client) build() (sql string) {

	if client.promeQL != "" {
		return client.promeQL
	}
	if client.query == "" || client.query == nil {
		sql = fmt.Sprintf("%s%s", client.metric, client.offset)
	} else {
		sql = fmt.Sprintf("%s%s%s", client.metric, Where(client.query), client.offset)
	}
	return sql
}

func (client *Client) where(have string, args ...interface{}) *Client {
	return client
}
