package istio

import (
	"context"
	"github.com/google/martian/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	meshconfig "istio.io/api/mesh/v1alpha1"
	"istio.io/istio/pkg/config/mesh"
	istio_inject "istio.io/istio/pkg/kube/inject"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

// istio 手动注入或者自动注入
type Injection interface {
	//给namespace添加 sidecar
	AutoInject(name string) error
	//给namespace删除 sidecar
	DeleteInject(name string) error
	//给pod添加 sidecar
	Manual(obj runtime.Object) (out interface{}, err error)
}

type Inject struct {
	Injection
	clientSet *kubernetes.Clientset
	configMap *ConfigMap
}

func New(clientSet *kubernetes.Clientset) *Inject {
	configMap := &ConfigMap{
		client: clientSet,
	}
	return &Inject{clientSet: clientSet, configMap: configMap}
}

//Auto  edit namespace
func (i *Inject) AutoInject(name string) error {
	namespace, err := i.clientSet.CoreV1().Namespaces().Get(context.TODO(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	namespace.Labels["istio-injection"] = "enable"
	_, err = i.clientSet.CoreV1().Namespaces().Update(context.TODO(), namespace, meta_v1.UpdateOptions{})
	return err
}

//Auto  edit namespace
func (i *Inject) DeleteInject(name string) error {
	namespace, err := i.clientSet.CoreV1().Namespaces().Get(context.TODO(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	delete(namespace.Labels, "istio-injection")
	_, err = i.clientSet.CoreV1().Namespaces().Update(context.TODO(), namespace, meta_v1.UpdateOptions{})
	return err
}

// 手动注入默认参数
func (i *Inject) Manual(obj runtime.Object) (out interface{}, err error) {
	meshConfig, err := i.getMeshConfig("")
	if err != nil {
		return
	}
	sidecarTemplate, err := i.sidecarTemplate("")
	if err != nil {
		return
	}
	valueConfig, err := i.valuesConfig("")
	if err != nil {
		return
	}
	return istio_inject.IntoObject(sidecarTemplate,
		valueConfig,
		"",
		meshConfig,
		obj)
}

func (i *Inject) ManualFile(obj runtime.Object, injectConfigFile, meshconfigFile, valuesConfig string) (out interface{}, err error) {
	meshConfig, err := i.getMeshConfig(meshconfigFile)
	if err != nil {
		return
	}
	sidecarTemplate, err := i.sidecarTemplate(injectConfigFile)
	if err != nil {
		return
	}
	valueConfig, err := i.valuesConfig(valuesConfig)
	if err != nil {
		return
	}
	return istio_inject.IntoObject(sidecarTemplate,
		valueConfig,
		"",
		meshConfig,
		obj)
}

func (i *Inject) getMeshConfig(meshConfigFile string) (meshConfig *meshconfig.MeshConfig, err error) {
	if meshConfigFile != "" {
		if meshConfig, err = mesh.ReadMeshConfig(meshConfigFile); err != nil {
			return nil, err
		}
	}

	if meshConfig, err = i.configMap.GetMeshConfigFromConfigMap("kube-inject"); err != nil {
		return nil, err
	}
	return
}

func (i *Inject) sidecarTemplate(injectConfigFile string) (sidecarTemplate string, err error) {
	if injectConfigFile != "" {
		injectionConfig, err := ioutil.ReadFile(injectConfigFile) // nolint: vetshadow
		if err != nil {
			return "", err
		}
		var injectConfig istio_inject.Config
		if err := yaml.Unmarshal(injectionConfig, &injectConfig); err != nil {
			log.Errorf("yaml unmarshal err:%v", err)
			return "", err
		}
		sidecarTemplate = injectConfig.Template
	} else if sidecarTemplate, err = i.configMap.GetInjectConfigFromConfigMap(); err != nil {
		return "", err
	}
	return sidecarTemplate, nil
}

func (i *Inject) valuesConfig(valuesFile string) (valuesConfig string, err error) {
	if valuesFile != "" {
		valuesConfigBytes, err := ioutil.ReadFile(valuesFile) // nolint: vetshadow
		if err != nil {
			return "", err
		}

		valuesConfig = string(valuesConfigBytes)
	} else if valuesConfig, err = i.configMap.GetValuesFromConfigMap(); err != nil {
		return "", err
	}
	return valuesConfig, nil
}
