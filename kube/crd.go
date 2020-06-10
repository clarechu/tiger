package kube

import (
	"context"
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	api_err "k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	DefaultCRDVersion = "v1alpha1"
)

type CustomResourceDefinitionInterface interface {
	Create(resource *CustomResourceDefinition) (crd *v1beta1.CustomResourceDefinition, err error)
	Get(name string) (crd *v1beta1.CustomResourceDefinition, err error)
	Delete(name string) (err error)
	Update(resource *CustomResourceDefinition) (crd *v1beta1.CustomResourceDefinition, err error)
}

type Resource struct {
	CustomResourceDefinitionInterface
	clientSet *apiextension.Clientset
}

func NewResource(clientSet *apiextension.Clientset) (resource *Resource) {
	return &Resource{clientSet: clientSet}
}
// +optional
func (c *Resource) Create(resource *CustomResourceDefinition) (crd *v1beta1.CustomResourceDefinition, err error) {
	crd = &v1beta1.CustomResourceDefinition{
		ObjectMeta: meta_v1.ObjectMeta{Name: resource.FullCRDName},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   resource.CRDGroup,
			Version: DefaultCRDVersion,
			Scope:   v1beta1.NamespaceScoped,
			Names: v1beta1.CustomResourceDefinitionNames{
				Plural: resource.CRDPlural,
				Kind:   resource.Kind,
			},
		},
	}
	crd, err = c.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Create(context.TODO(), crd, meta_v1.CreateOptions{})
	if api_err.IsAlreadyExists(err) {
		return crd, nil
	}
	return
}

func (c *Resource) Get(name string) (crd *v1beta1.CustomResourceDefinition, err error) {
	ops := meta_v1.GetOptions{}
	crd, err = c.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Get(context.TODO(), name, ops)
	return
}

func (c *Resource) Delete(name string) (err error) {
	ops := meta_v1.DeleteOptions{}
	err = c.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(context.TODO(), name, ops)
	if api_err.IsAlreadyExists(err) {
		return nil
	}
	return
}

func (c *Resource) Update(resource *CustomResourceDefinition) (crd *v1beta1.CustomResourceDefinition, err error) {
	crd = &v1beta1.CustomResourceDefinition{
		ObjectMeta: meta_v1.ObjectMeta{Name: resource.FullCRDName},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   resource.CRDGroup,
			Version: DefaultCRDVersion,
			Scope:   v1beta1.NamespaceScoped,
			Names: v1beta1.CustomResourceDefinitionNames{
				Plural: resource.CRDPlural,
				Kind:   resource.Kind,
			},
		},
	}
	crd, err = c.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Update(crd)
	if api_err.IsAlreadyExists(err) {
		return crd, nil
	}
	return
}
