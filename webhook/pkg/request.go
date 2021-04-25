package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"istio.io/pkg/log"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type Request struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

//IntoResourceFile 自动注入接口
func (i *Inject) Start(resp http.ResponseWriter, req *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			http.Error(resp, "no body found", http.StatusBadRequest)
			return
		}
	}()
	var body []byte
	if req.Body != nil {
		if data, err := ioutil.ReadAll(req.Body); err == nil {
			body = data
		}
	}
	if len(body) == 0 {
		log.Error("body nil")
		http.Error(resp, "no body found", http.StatusBadRequest)
		return
	}
	// verify the content type is accurate
	contentType := req.Header.Get("Content-Type")
	if contentType != "application/json" {
		log.Infof("contentType :%v", contentType)
		http.Error(resp, "invalid Content-Type, want `application/json`", http.StatusUnsupportedMediaType)
		return
	}
	r := Request{}
	err := json.Unmarshal(body, &r)
	deployment, err := i.ClientSet.AppsV1().Deployments(r.Namespace).Get(context.TODO(), r.Name, metav1.GetOptions{})
	if err != nil {
		http.Error(resp, fmt.Sprintf("get deployment name:%v, namespace:%v, err: %v", r.Namespace,
			r.Name, err),
			http.StatusUnsupportedMediaType)
		return
	}
	updates, err := i.ManualInject(deployment)
	if err != nil {
		http.Error(resp, fmt.Sprintf("sidecar updates error :%v", err), http.StatusUnsupportedMediaType)
		return
	}
	de := updates.(*v1.Deployment)
	_, err = i.ClientSet.AppsV1().Deployments(r.Namespace).Update(context.TODO(), de, metav1.UpdateOptions{})
	if err != nil {
		http.Error(resp, fmt.Sprintf("update deployment is error:%v", err), http.StatusUnsupportedMediaType)
		return
	}
}

//IntoResourceFile 自动注入接口
func (un *Uninject) Start(resp http.ResponseWriter, req *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			http.Error(resp, "no body found", http.StatusBadRequest)
			return
		}
	}()
	var body []byte
	if req.Body != nil {
		if data, err := ioutil.ReadAll(req.Body); err == nil {
			body = data
		}
	}
	if len(body) == 0 {
		log.Error("body nil")
		http.Error(resp, "no body found", http.StatusBadRequest)
		return
	}
	// verify the content type is accurate
	contentType := req.Header.Get("Content-Type")
	if contentType != "application/json" {
		log.Infof("contentType :%v", contentType)
		http.Error(resp, "invalid Content-Type, want `application/json`", http.StatusUnsupportedMediaType)
		return
	}
	r := Request{}
	err := json.Unmarshal(body, &r)
	if err != nil {
		http.Error(resp, fmt.Sprintf("unmarshal json error:%v", err), http.StatusUnsupportedMediaType)
		return
	}
	err = un.unInjectDeploymentName(r.Name, r.Namespace)
	if err != nil {
		http.Error(resp, fmt.Sprintf("update deployment is error:%v", err), http.StatusUnsupportedMediaType)
		return
	}
}
