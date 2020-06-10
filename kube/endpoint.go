package kube

import (
	"context"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type EndPoint struct {
	clientSet *kubernetes.Clientset
}

func NewEndPoint(clientSet *kubernetes.Clientset) *EndPoint {
	return &EndPoint{clientSet: clientSet}
}

func (e *EndPoint) Create(namespace string, endpoints *v1.Endpoints) (err error) {
	endpoints, err = e.clientSet.CoreV1().Endpoints(namespace).Create(context.TODO(), endpoints, meta_v1.CreateOptions{})
	return
}

func (e *EndPoint) Delete(namespace, name string) (err error) {
	ops := meta_v1.DeleteOptions{}
	err = e.clientSet.CoreV1().Endpoints(namespace).Delete(context.TODO(), name, ops)
	return
}

func (e *EndPoint) Update(namespace string, endpoints *v1.Endpoints) (err error) {
	endpoints, err = e.clientSet.CoreV1().Endpoints(namespace).Update(context.TODO(), endpoints, meta_v1.UpdateOptions{})
	return
}

func (e *EndPoint) Get(name, namespace string) (endpoints *v1.Endpoints, err error) {
	ops := meta_v1.GetOptions{}
	endpoints, err = e.clientSet.CoreV1().Endpoints(namespace).Get(context.TODO(), name, ops)
	return
}

func (e *EndPoint) List(namespace string, ops meta_v1.ListOptions) (endpoints *v1.EndpointsList, err error) {
	endpoints, err = e.clientSet.CoreV1().Endpoints(namespace).List(context.TODO(), ops)
	return
}
