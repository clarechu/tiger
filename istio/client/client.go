package client

import (
	"errors"
	"flag"
	client "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type Client interface {
	GetDefaultClient() (clientSet *client.Clientset, err error)
	NewClient(kubeConfig, namespace string) (clientSet *client.Clientset, err error)
}

func GetDefaultClient() (clientSet *client.Clientset, err error) {
	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeConfig file")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "absolute path to the kubeConfig file")
	}
	restConfig, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		return nil, err
	}

	ic, err := client.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return ic, err
}

func NewClient(kubeConfig, namespace string) (clientSet *client.Clientset, err error) {
	if len(kubeConfig) == 0 || len(namespace) == 0 {
		return nil, errors.New("Environment variables KUBECONFIG and NAMESPACE need to be set")
	}

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, err
	}

	ic, err := client.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return ic, err
}
