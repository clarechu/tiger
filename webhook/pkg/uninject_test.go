package pkg

import (
	"github.com/ClareChu/tiger/kube/client"
	"github.com/bmizerany/assert"
	"testing"
)

func TestAutoInject(t *testing.T) {
	clientSet, err := client.GetDefaultK8sClientSet()
	assert.Equal(t, nil, err)
	uninject := &Uninject{
		ClientSet: clientSet,
	}
	err = uninject.unInjectDeployment("tcp-echo", "istio-io-tcp-traffic-shifting")
}
