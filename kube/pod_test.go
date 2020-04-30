package kube

import (
	"github.com/ClareChu/tiger/kube/client"
	"github.com/bmizerany/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestNewPod(t *testing.T) {
	namespace := "default"
	name := "details-v1-74f858558f-rmmrb"
	clientSet, err := client.GetDefaultK8sClientSet()
	assert.Equal(t, err, nil)
	pod := NewPod(clientSet)
	ops := metav1.ListOptions{}
	_, err = pod.List(namespace, ops)
	assert.Equal(t, nil, err)
	_, err = pod.Get(name, namespace)
	assert.Equal(t, nil, err)
	err = pod.Delete(name, namespace)
	assert.Equal(t, nil, err)
}