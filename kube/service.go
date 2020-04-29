package kube

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Service struct {
	clientset *kubernetes.Clientset
}

func NewService(clientset *kubernetes.Clientset) *Service {
	return &Service{clientset: clientset}
}

func (c *Service) Create(namespace string, service *v1.Service) (err error) {
	service, err = c.clientset.CoreV1().Services(namespace).Create(service)
	return
}

func (c *Service) Delete(namespace, name string) (err error) {
	ops := &meta_v1.DeleteOptions{}
	err = c.clientset.CoreV1().Services(namespace).Delete(name, ops)
	return
}

func (c *Service) Update(namespace string, service *v1.Service) (err error) {
	service, err = c.clientset.CoreV1().Services(namespace).Update(service)
	return
}

func (c *Service) Get(name, namespace string) (service *v1.Service, err error) {
	ops := meta_v1.GetOptions{}
	service, err = c.clientset.CoreV1().Services(namespace).Get(name, ops)
	return
}

func (c *Service) List(namespace string, ops meta_v1.ListOptions) (services *v1.ServiceList, err error) {
	services, err = c.clientset.CoreV1().Services(namespace).List(ops)
	return
}
