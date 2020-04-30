package kube

import (
	"github.com/ClareChu/tiger/kube/client"
	"github.com/bmizerany/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestService_Create(t *testing.T) {
	namespace := "default"
	name := "demo"
	clientSet, err := client.GetDefaultK8sClientSet()
	assert.Equal(t, err, nil)
	service := NewService(clientSet)
	svc := &v1.Service{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app": "v1",
			},
			Ports: []v1.ServicePort{
				{
					Name: "8080-http",
					Port: 8080,
				},
			},
		},
	}
	err = service.Create(namespace, svc)
	assert.Equal(t, nil, err)
	svc, err = service.Get(name, namespace)
	svc.Spec.Selector["app"] = "v2"
	err = service.Update(namespace, svc)
	assert.Equal(t, nil, err)
	ops := metav1.ListOptions{}
	_, err = service.List(namespace, ops)
	err = service.Delete(namespace, name)
	assert.Equal(t, nil, err)
}
