package istio

import (
	"context"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	client "istio.io/client-go/pkg/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VirtualService interface {
}

type Virtual struct {
	clientset *client.Clientset
	VirtualService
}

func NewVirtualService(clientset *client.Clientset) *Virtual {
	return &Virtual{
		clientset: clientset,
	}
}

func (v *Virtual) Create(namespace string, vs *v1alpha3.VirtualService) {
	vs, err := v.clientset.NetworkingV1alpha3().VirtualServices(namespace).Create(context.TODO(), vs, v1.CreateOptions{})
	if err != nil {

	}
	return
}

func (v *Virtual) Delete(name, namespace string) error {
	option := v1.DeleteOptions{}
	err := v.clientset.NetworkingV1alpha3().VirtualServices(namespace).Delete(context.TODO(), name, option)
	return err
}
