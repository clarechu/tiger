package istio

import (
	"bufio"
	"fmt"
	"github.com/ClareChu/tiger/kube/client"
	"github.com/bmizerany/assert"
	"github.com/ghodss/yaml"
	istio_inject "istio.io/istio/pkg/kube/inject"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yamlDecoder "k8s.io/apimachinery/pkg/util/yaml"
	"os"
	"testing"
)

// 通过文件inject
func TestInject_Manual(t *testing.T) {
	clientSet, err := client.GetDefaultK8sClientSet()
	assert.Equal(t, err, nil)
	in, err := os.Open("./test/nginx.yml")
	assert.Equal(t, err, nil)
	reader := yamlDecoder.NewYAMLReader(bufio.NewReaderSize(in, 4096))
	raw, err := reader.Read()
	obj, err := istio_inject.FromRawToObject(raw)
	out, err := New(clientSet).Manual(obj)
	assert.Equal(t, nil, err)
	fmt.Print(out)
}

func TestInject_Manual2(t *testing.T) {
	var replicas int32
	replicas = 1
	deploy := v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind: "Deployment",
			APIVersion: "extensions/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: "default",
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name: "nginx",
							Image: "nginx",
							Env: []corev1.EnvVar{
								{
									Name: "zx",
									Value: "qs",
								},
							},
						},
					},
				},
			},
		},
	}
	raw, err := yaml.Marshal(deploy)
	clientSet, err := client.GetDefaultK8sClientSet()
	assert.Equal(t, err, nil)
	obj, err := istio_inject.FromRawToObject(raw)
	out, err := New(clientSet).Manual(obj)
	assert.Equal(t, nil, err)
	fmt.Print(out)
}
