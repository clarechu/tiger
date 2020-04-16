package client

import (
	"flag"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type Client interface {
	GetDefaultK8sClientSet() (clientset *kubernetes.Clientset, err error)
	GetDefaultApiExtensionClientSet() (clientSet *apiextension.Clientset, err error)
	GetK8sClientSet(kubeConfig string) (clientset *kubernetes.Clientset, err error)
	GetApiExtensionClientSet(kubeConfig string) (clientSet *apiextension.Clientset, err error)
}

func GetDefaultK8sClientSet() (clientset *kubernetes.Clientset, err error) {
	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeConfig file")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "absolute path to the kubeConfig file")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		return
	}
	clientset, err = kubernetes.NewForConfig(config)
	return
}

func GetDefaultApiExtensionClientSet() (clientSet *apiextension.Clientset, err error) {
	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeConfig file")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "absolute path to the kubeConfig file")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		return
	}
	clientSet, err = apiextension.NewForConfig(config)
	return
}

func GetK8sClientSet(kubeConfig string) (clientset *kubernetes.Clientset, err error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return
	}
	clientset, err = kubernetes.NewForConfig(config)
	return
}

func GetApiExtensionClientSet(kubeConfig string) (clientSet *apiextension.Clientset, err error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return
	}
	clientSet, err = apiextension.NewForConfig(config)
	return
}
