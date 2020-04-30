package kube

import "k8s.io/client-go/kubernetes"

type ServiceAccountInterface interface {
	Create()
}

type ServiceAccount struct {
	ServiceAccountInterface
	clientSet *kubernetes.Clientset
}

func (s *ServiceAccount) Create() {

}