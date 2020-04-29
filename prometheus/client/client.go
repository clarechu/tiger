package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ClareChu/gorequest"
	"time"
)

type Prometheus interface {
	Query(query string) *Client
	Range(start, end float64) *Client
	Step(step int) *Client
	Run() (pro *ProRequest, err error)
}

type Client struct {
	Prometheus
	host  string  `json:"host"`
	port  int32   `json:"port"`
	start float64 `json:"start"`
	end   float64 `json:"end"`
	step  int     `json:"step"`
	query string  `json:"query"`
	time  bool    `json:"time"`
	url   string  `json:"url"`
}

//prometheus response result
type ProRequest struct {
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

func New(host string, ports ...int32) *Client {
	var port int32
	if len(ports) == 0 {
		port = DefaultPrometheusHost
	} else {
		port = ports[0]
	}
	return &Client{
		host: host,
		port: port,
	}
}

//Query 指标名称
func (client *Client) Query(query string) *Client {
	client.query = query
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

func (client *Client) Run() (pro *ProRequest, err error) {
	client.setUrl().appendRange().appendParam()
	fmt.Println("request url:", client.url)
	resp, body, errs := gorequest.New().Get(client.url).End()
	if len(errs) != 0 {
		return nil, errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(body)
	}
	pro = &ProRequest{}
	err = json.Unmarshal([]byte(body), pro)
	if err != nil {
		return
	}
	return
}

//1588150253687
//1588150337
//1588150254.061
func (client *Client) setUrl() *Client {
	url := fmt.Sprintf("http://%s:%d", client.host, client.port)
	client.url = url
	return client
}

func (client *Client) appendRange() *Client {
	url := client.url
	if client.time {
		url = url + "/api/v1/query_range"
	}
	url = url + "/api/v1/query"
	client.url = url
	return client
}

func (client *Client) appendParam() *Client {
	now := time.Now().Unix()
	url := client.url
	url = url + "?query=" + client.query
	if client.time {
		url = fmt.Sprintf("%s&start=%f&end=%f", url, client.start, client.end)
	} else {
		if client.start == 0 {
			client.start = float64(now - 60)
			url = fmt.Sprintf("%s&time=%f", url, client.start)
		}

	}
	url = fmt.Sprintf("%s&step=%d&_=%d", url, client.step, now)
	client.url = url
	return client
}

func PromQL(o interface{})  {

}