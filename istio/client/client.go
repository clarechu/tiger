package istio

import (
	client "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)



func GetClient(kubeConfig, namespace string) (clientSet *client.Clientset, err error) {
	if len(kubeConfig) == 0 || len(namespace) == 0 {
		log.Fatalf("Environment variables KUBECONFIG and NAMESPACE need to be set")
	}

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		log.Fatalf("Failed to create k8s rest client: %s", err)
	}

	ic, err := client.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("Failed to create istio client: %s", err)
	}
	return ic, err
}
