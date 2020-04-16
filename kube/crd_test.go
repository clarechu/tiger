package kube_test

import (
	"fmt"
	"github.com/ClareChu/tiger/kube"
	"github.com/ClareChu/tiger/kube/client"
	"testing"
)

func TestResource_CURD(t *testing.T) {
	clientSet, err := client.GetDefaultApiExtensionClientSet()
	if err != nil {
		fmt.Println("client set error", err)
		return
	}
	customResourceDefinition := &kube.CustomResourceDefinition{
		FullCRDName: "demo.demo.io",
		CRDGroup:    "demo.io",
		CRDPlural:   "demo",
		Kind:        "Demo",
	}
	_, err = kube.NewResource(clientSet).Create(customResourceDefinition)
	_, err = kube.NewResource(clientSet).Get("demo.demo.io")
	_, err = kube.NewResource(clientSet).Update(customResourceDefinition)
	err = kube.NewResource(clientSet).Delete("demo.demo.io")
}
