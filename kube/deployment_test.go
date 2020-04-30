package kube

import (
	"github.com/ClareChu/tiger/kube/client"
	"github.com/bmizerany/assert"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestNewDeployment(t *testing.T) {
	namespace := "default"
	name := "demo"
	var replicase int32
	replicase = 1
	clientSet, err := client.GetDefaultK8sClientSet()
	assert.Equal(t, err, nil)
	d := NewDeployment(clientSet)
	svc := &v1.Deployment{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels: map[string]string{
				"app": "demo",
			},
		},
		Spec: v1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Replicas: &replicase,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx-v1",
							Image: "nginx",
						},
					},
				},
			},
		},
	}
	err = d.Create(namespace, svc)
	assert.Equal(t, nil, err)
	svc, err = d.Get(name, namespace)
	err = d.Update(namespace, svc)
	assert.Equal(t, nil, err)
	ops := metav1.ListOptions{}
	_, err = d.List(namespace, ops)
	err = d.Delete(name, namespace)
	assert.Equal(t, nil, err)
}
