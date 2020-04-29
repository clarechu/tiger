package client

import (
	"fmt"
	"testing"
)

func TestClient_Query(t *testing.T) {
	pro, err := New("10.10.13.39", 31002).Query("istio_request_bytes_bucket{connection_security_policy=\"none\"}").Step(7).Run()
	fmt.Println(err)
	fmt.Println(pro)
}
