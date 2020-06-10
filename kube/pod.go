package kube

import (
	"context"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Pod struct {
	clientSet *kubernetes.Clientset
}

func NewPod(clientSet *kubernetes.Clientset) *Pod {
	return &Pod{clientSet: clientSet}
}

func (p *Pod) Delete(name, namespace string) (err error) {
	ops := meta_v1.DeleteOptions{}
	err = p.clientSet.CoreV1().Pods(namespace).Delete(context.TODO(), name, ops)
	return
}

func (p *Pod) Get(name, namespace string) (pod *v1.Pod, err error) {
	ops := meta_v1.GetOptions{}
	pod, err = p.clientSet.CoreV1().Pods(namespace).Get(context.TODO(), name, ops)
	return
}

func (p *Pod) List(namespace string, ops meta_v1.ListOptions) (pod *v1.PodList, err error) {
	pod, err = p.clientSet.CoreV1().Pods(namespace).List(context.TODO(), ops)
	return
}
