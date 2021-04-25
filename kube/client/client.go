package client

import (
	"flag"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

type Client interface {
	GetDefaultK8sClientSet() (clientset *kubernetes.Clientset, err error)
	GetDefaultApiExtensionClientSet() (clientSet *apiextension.Clientset, err error)
	GetK8sClientSet(kConfig string) (clientset *kubernetes.Clientset, err error)
	GetApiExtensionClientSet(kConfig string) (clientSet *apiextension.Clientset, err error)
}

func GetDefaultK8sClientSet() (clientSet *kubernetes.Clientset, err error) {
	var config *rest.Config
	if os.Getenv("KUBERNETES_SERVICE_HOST") == "" {
		kubeConfig := GetKubeConfig()
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			return
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			return
		}
	}
	clientSet, err = kubernetes.NewForConfig(config)
	return
}

func GetDefaultApiExtensionClientSet() (clientSet *apiextension.Clientset, err error) {
	var config *rest.Config
	if os.Getenv("KUBERNETES_SERVICE_HOST") == "" {
		var kubeConfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeConfig file")
		} else {
			kubeConfig = flag.String("kubeConfig", "", "absolute path to the kubeConfig file")
		}
		config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
		if err != nil {
			return
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			return
		}
	}
	clientSet, err = apiextension.NewForConfig(config)
	return
}

func GetK8sClientSet(kConfig string) (clientset *kubernetes.Clientset, err error) {
	config, err := clientcmd.BuildConfigFromFlags("", kConfig)
	if err != nil {
		return
	}
	clientset, err = kubernetes.NewForConfig(config)
	return
}

func GetApiExtensionClientSet(kConfig string) (clientSet *apiextension.Clientset, err error) {
	config, err := clientcmd.BuildConfigFromFlags("", kConfig)
	if err != nil {
		return
	}
	clientSet, err = apiextension.NewForConfig(config)
	return
}

//GetKubeConfig
func GetKubeConfig() (kubeConfig string) {
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeConfig = os.Getenv("KUBECONFIG")
	}
	return
}
