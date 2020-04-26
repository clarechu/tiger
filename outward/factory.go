package outward

import (
	"github.com/ClareChu/tiger/kube"
	"github.com/ClareChu/tiger/kube/client"
	k8s "k8s.io/client-go/kubernetes"
	"log"
)

//工厂 初始化kube istio client
type Kube struct {
	clientset *k8s.Clientset
}

type kubes struct {
	clients map[string][]*k8s.Clientset
}

type istio struct {
}

func NewDefaultKube() *Kube {
	clientset, err := client.GetDefaultK8sClientSet()
	if err != nil {
		log.Printf("get default client err:%v", err)
		return nil
	}
	return &Kube{clientset: clientset}
}

func GetKube() {

}

func SetKube() {

}

func (k *Kube) ConfigMap() *kube.ConfigMap {
	return kube.NewConfigMap(k.clientset)
}
