package kube

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Namespace struct {
	clientSet *kubernetes.Clientset
}

func NewNamespace(clientSet *kubernetes.Clientset) *Namespace {
	return &Namespace{clientSet: clientSet}
}

func (e *Namespace) Create(namespace *v1.Namespace) (ns *v1.Namespace, err error) {
	ns, err = e.clientSet.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	return
}

func (e *Namespace) Delete(name string) (err error) {
	ops := metav1.DeleteOptions{}
	err = e.clientSet.CoreV1().Namespaces().Delete(context.TODO(), name, ops)
	return
}

func (e *Namespace) Update(namespace *v1.Namespace) (err error) {
	namespace, err = e.clientSet.CoreV1().Namespaces().Update(context.TODO(), namespace, metav1.UpdateOptions{})
	return
}

func (e *Namespace) Get(name string) (namespace *v1.Namespace, err error) {
	ops := metav1.GetOptions{}
	namespace, err = e.clientSet.CoreV1().Namespaces().Get(context.TODO(), name, ops)
	return
}

func (e *Namespace) List(ops metav1.ListOptions) (namespaces *v1.NamespaceList, err error) {
	namespaces, err = e.clientSet.CoreV1().Namespaces().List(context.TODO(), ops)
	return
}

