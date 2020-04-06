package kube

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMap struct {
	clientset *kubernetes.Clientset
}

func NewConfigMap(clientset *kubernetes.Clientset) *ConfigMap {
	return &ConfigMap{clientset: clientset}
}

func (c *ConfigMap) Create(namespace string, service *v1.ConfigMap) (err error) {
	service, err = c.clientset.CoreV1().ConfigMaps(namespace).Create(service)
	return
}

func (c *ConfigMap) Delete(namespace, name string) (err error) {
	ops := &meta_v1.DeleteOptions{}
	err = c.clientset.CoreV1().ConfigMaps(namespace).Delete(name, ops)
	return
}

func (c *ConfigMap) Update(namespace string, service *v1.ConfigMap) (err error) {
	service, err = c.clientset.CoreV1().ConfigMaps(namespace).Update(service)
	return
}

func (c *ConfigMap) Get(name, namespace string) (service *v1.ConfigMap, err error) {
	ops := meta_v1.GetOptions{}
	service, err = c.clientset.CoreV1().ConfigMaps(namespace).Get(name, ops)
	return
}

func (c *ConfigMap) List(namespace string) (services *v1.ConfigMapList, err error) {
	ops := meta_v1.ListOptions{}
	services, err = c.clientset.CoreV1().ConfigMaps(namespace).List(ops)
	return
}
