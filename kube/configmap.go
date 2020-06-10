package kube

import (
	"context"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMap struct {
	clientSet *kubernetes.Clientset
}

func NewConfigMap(clientSet *kubernetes.Clientset) *ConfigMap {
	return &ConfigMap{clientSet: clientSet}
}

func (c *ConfigMap) Create(namespace string, service *v1.ConfigMap) (err error) {
	service, err = c.clientSet.CoreV1().ConfigMaps(namespace).Create(context.TODO(), service, meta_v1.CreateOptions{})
	return
}

func (c *ConfigMap) Delete(namespace, name string) (err error) {
	ops := meta_v1.DeleteOptions{}
	err = c.clientSet.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), name, ops)
	return
}

func (c *ConfigMap) Update(namespace string, service *v1.ConfigMap) (err error) {
	service, err = c.clientSet.CoreV1().ConfigMaps(namespace).Update(context.TODO(), service, meta_v1.UpdateOptions{})
	return
}

func (c *ConfigMap) Get(name, namespace string) (service *v1.ConfigMap, err error) {
	ops := meta_v1.GetOptions{}
	service, err = c.clientSet.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, ops)
	return
}

func (c *ConfigMap) List(namespace string, ops meta_v1.ListOptions) (services *v1.ConfigMapList, err error) {
	services, err = c.clientSet.CoreV1().ConfigMaps(namespace).List(context.TODO(), ops)
	return
}
