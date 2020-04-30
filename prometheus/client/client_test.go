package client

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestClient_Query(t *testing.T) {
	_, err := New("http://10.10.13.39:31002").
		PromQL("istio_request_bytes_bucket{connection_security_policy=\"none\"}").
		Step(7).
		Run()
	assert.Equal(t, nil, err)
}

type Foo struct {
	Name string `json:"name" pro:"name"`
	Age  int    `json:"age" pro:"age"`
	P    string `json:"p" pro:"p"`
}

func TestPromQL(t *testing.T) {
	foo := Foo{
		Name: "a",
		Age:  1,
		P:    "cxx",
	}
	sql := Where(foo)
	assert.Equal(t, sql, "{name=\"a\",p=\"cxx\"}")
	sql = Where(&foo)
	assert.Equal(t, sql, "{name=\"a\",p=\"cxx\"}")
}

func TestClient_Build(t *testing.T) {
	foo := Foo{
		Name: "a",
		Age:  1,
		P:    "cxx",
	}
	client := &Client{}
	sql := client.Query(foo).build()
	assert.Equal(t, sql, "{name=\"a\",p=\"cxx\"}")
	client = &Client{}
	sql = client.Query(foo).Metric("istio").build()
	assert.Equal(t, sql, "istio{name=\"a\",p=\"cxx\"}")
	client = &Client{}
	sql = client.Query(foo).Metric("istio").OffSet("[5m] offset 1w").build()
	assert.Equal(t, sql, "istio{name=\"a\",p=\"cxx\"}[5m] offset 1w")
	client = &Client{}
	sql = client.Metric("istio").OffSet("[5m] offset 1w").PromQL("xxx").build()
	assert.Equal(t, sql, "xxx")
	client = &Client{}
	sql = client.Metric("istio").OffSet("[5m] offset 1w").build()
	assert.Equal(t, sql, "istio[5m] offset 1w")
}
