package istio

import (
	"github.com/google/martian/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	meshconfig "istio.io/api/mesh/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
)

// istio 手动注入或者自动注入
type Injection interface {
}

type Inject struct {
	clientSet *kubernetes.Clientset
}

func New(clientSet *kubernetes.Clientset) *Inject {
	return &Inject{clientSet: clientSet}
}

//Auto  edit namespace
func (i *Inject) AutoInject(name string) error {
	namespace, err := i.clientSet.CoreV1().Namespaces().Get(name, v1.GetOptions{})
	if err != nil {
		return err
	}
	namespace.Labels["istio-injection"] = "enable"
	_, err = i.clientSet.CoreV1().Namespaces().Update(namespace)
	return err
}

//Auto  edit namespace
func (i *Inject) DeleteInject(name string) error {
	namespace, err := i.clientSet.CoreV1().Namespaces().Get(name, v1.GetOptions{})
	if err != nil {
		return err
	}
	delete(namespace.Labels, "istio-injection")
	_, err = i.clientSet.CoreV1().Namespaces().Update(namespace)
	return err
}

// 手动注入默认参数
func Manual() {
	outObject, err := IntoObject(sidecarTemplate, valuesConfig, meshconfig, obj)
}

func getMeshConfig(meshConfigFile string) (meshConfig *meshconfig.MeshConfig, err error) {
	if meshConfigFile != "" {
		if meshConfig, err = mesh.ReadMeshConfig(meshConfigFile); err != nil {
			return nil, err
		}
	}

	if meshConfig, err = getMeshConfigFromConfigMap(kubeconfig, "kube-inject"); err != nil {
		return nil, err
	}
	return
}

func sidecarTemplate(injectConfigFile string) (sidecarTemplate *string, err error) {
	if injectConfigFile != "" {
		injectionConfig, err := ioutil.ReadFile(injectConfigFile) // nolint: vetshadow
		if err != nil {
			return nil, err
		}
		var injectConfig inject.Config
		if err := yaml.Unmarshal(injectionConfig, &injectConfig); err != nil {
			log.Errorf("yaml unmarshal err:%v", err)
			return nil, err
		}
		sidecarTemplate = injectConfig.Template
	} else if sidecarTemplate, err = getInjectConfigFromConfigMap(kubeconfig); err != nil {
		return nil, err
	}
	return sidecarTemplate, nil
}

func valuesConfig(valuesFile string) (valuesConfig string, err error) {
	if valuesFile != "" {
		valuesConfigBytes, err := ioutil.ReadFile(valuesFile) // nolint: vetshadow
		if err != nil {
			return "", err
		}

		valuesConfig = string(valuesConfigBytes)
	} else if valuesConfig, err := getValuesFromConfigMap(kubeconfig); err != nil {
		return "", err
	}
	return valuesConfig, nil
}