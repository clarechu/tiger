package pkg

import (
	"context"
	"encoding/json"
	"github.com/ClareChu/tiger/webhook/cache"
	"io/ioutil"
	"istio.io/pkg/log"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

type Request struct {
	Context   string `json:"context"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

//IntoResourceFile 自动注入接口
func (i *Inject) IntoResourceFile(resp http.ResponseWriter, req *http.Request) {
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
	restConfig := cache.Get().Clusters[r.Context].KubeConfig
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		http.Error(resp, "get client set err", http.StatusUnsupportedMediaType)
		return
	}
	deployment, err := clientSet.AppsV1().Deployments(r.Namespace).Get(context.TODO(), r.Name, metav1.GetOptions{})
	updates, err := i.ManualInject(r.Context, deployment)
	if err != nil {
		http.Error(resp, "get updates", http.StatusUnsupportedMediaType)
		return
	}
	de := updates.(*v1.Deployment)
	_, err = clientSet.AppsV1().Deployments(r.Namespace).Update(context.TODO(), de, metav1.UpdateOptions{})
	if err != nil {
		http.Error(resp, "get updates", http.StatusUnsupportedMediaType)
		return
	}
}
