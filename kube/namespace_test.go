package kube

import (
	"github.com/ClareChu/tiger/kube/client"
	"gotest.tools/assert"
	"testing"
)

func TestNewNamespace(t *testing.T) {
	ns := "fp20051302054771600000025422610"
	clientSet, err := client.GetDefaultK8sClientSet()
	assert.Equal(t, err, nil)
	n := NewNamespace(clientSet)
/*	namespace := &v1.Namespace{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: ns,
		},
		Spec: v1.NamespaceSpec{
			Finalizers: []v1.FinalizerName{
			},
		},
	}
	_, err = n.Create(namespace)
	assert.Equal(t, err, nil)*/
	nss, err := n.Get(ns)
	nss.ObjectMeta.Annotations["openshift.io/sa.scc.uid-range"] = "0/0"
	err = n.Update(nss)
	assert.Equal(t, err, nil)
}
