package pkg

import (
	"context"
	"fmt"
	"github.com/ClareChu/tiger/webhook/cache"
	"gotest.tools/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"testing"
)

func TestIntoResourceFile(t *testing.T) {
	conf, err := cache.NewConfig()
	assert.Equal(t, nil, err)
	cache.Set(conf)
	inject := &Inject{
	}
	restConfig := cache.Get().Clusters["cluster02"].KubeConfig
	clientSet, err := kubernetes.NewForConfig(restConfig)
	assert.Equal(t, nil, err)
	deployment, err := clientSet.AppsV1().Deployments("default").Get(context.TODO(), "helloworld", metav1.GetOptions{})
	assert.Equal(t, nil, err)
	//deployment.Kind = "Deployment"
	//deployment.APIVersion = "extensions/v1beta1"
	values, err := inject.ManualInject("cluster02", deployment)
	assert.Equal(t, nil, err)
	fmt.Println(values)
}
