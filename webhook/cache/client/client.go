package client

import (
	"errors"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

type Client interface {
	GetDefaultK8sClientSet() (clientSet *kubernetes.Clientset, err error)
	GetDefaultApiExtensionClientSet() (clientSet *apiextension.Clientset, err error)
	GetK8sClientSet(kConfig string) (clientSet *kubernetes.Clientset, err error)
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

	clientSet, err = apiextension.NewForConfig(config)
	return
}

func GetApiExtensionClientSet(kubeConfigBytes []byte) (clientSet *apiextension.Clientset, err error) {
	defer func() {
		err := recover()
		if err != nil {
			return
		}
	}()
	cf, err := LoadFromFile(kubeConfigBytes)
	if err != nil {
		return nil, err
	}
	authInfo := ""
	cluster := ""
	for _, v := range cf.Contexts {
		authInfo = v.AuthInfo
		cluster = v.Cluster
	}
	config := &rest.Config{
		Host: cf.Clusters[cluster].Server,
		TLSClientConfig: rest.TLSClientConfig{
			CAData:   cf.Clusters[cluster].CertificateAuthorityData,
			CertData: cf.AuthInfos[authInfo].ClientCertificateData,
			KeyData:  cf.AuthInfos[authInfo].ClientKeyData,
		},
	}
	clientSet, err = apiextension.NewForConfig(config)
	return
}

func GetK8sClientSet(kubeConfigBytes []byte) (clientSet *kubernetes.Clientset, err error) {
	defer func() {
		err := recover()
		if err != nil {
			err = errors.New("")
			return
		}
	}()
	cf, err := LoadFromFile(kubeConfigBytes)
	if err != nil {
		return nil, err
	}
	config := &rest.Config{
		Host: cf.Clusters[cf.Contexts[cf.CurrentContext].Cluster].Server,
		TLSClientConfig: rest.TLSClientConfig{
			CAData:   cf.Clusters[cf.Contexts[cf.CurrentContext].Cluster].CertificateAuthorityData,
			CertData: cf.AuthInfos[cf.Contexts[cf.CurrentContext].AuthInfo].ClientCertificateData,
			KeyData:  cf.AuthInfos[cf.Contexts[cf.CurrentContext].AuthInfo].ClientKeyData,
		},
	}
	clientSet, err = kubernetes.NewForConfig(config)
	return
}

// LoadFromFile takes a filename and deserializes the contents into Config object
func LoadFromFile(kubeconfigBytes []byte) (*clientcmdapi.Config, error) {
	config, err := clientcmd.Load(kubeconfigBytes)
	if err != nil {
		return nil, err
	}

	// set LocationOfOrigin on every Cluster, User, and Context
	for key, obj := range config.AuthInfos {
		obj.LocationOfOrigin = "default"
		config.AuthInfos[key] = obj
	}
	for key, obj := range config.Clusters {
		obj.LocationOfOrigin = "default"
		config.Clusters[key] = obj
	}
	for key, obj := range config.Contexts {
		obj.LocationOfOrigin = "default"
		config.Contexts[key] = obj
	}

	if config.AuthInfos == nil {
		config.AuthInfos = map[string]*clientcmdapi.AuthInfo{}
	}
	if config.Clusters == nil {
		config.Clusters = map[string]*clientcmdapi.Cluster{}
	}
	if config.Contexts == nil {
		config.Contexts = map[string]*clientcmdapi.Context{}
	}

	return config, nil
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
