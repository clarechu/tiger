package pkg

import (
	"context"
	"fmt"
	"github.com/ClareChu/tiger/kube/client"
	"gotest.tools/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestIntoResourceFile(t *testing.T) {
	clientSet, err := client.GetDefaultK8sClientSet()
	assert.Equal(t, nil, err)
	deployment, err := clientSet.AppsV1().Deployments("default").Get(context.TODO(), "helloworld", metav1.GetOptions{})
	assert.Equal(t, nil, err)
	inject := Inject{}
	values, err := inject.ManualInject(deployment)
	assert.Equal(t, nil, err)
	fmt.Println(values)
}
