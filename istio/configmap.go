package istio

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v2"
	meshconfig "istio.io/api/mesh/v1alpha1"
	"istio.io/istio/pkg/config/mesh"
	istio_inject "istio.io/istio/pkg/kube/inject"
	"istio.io/pkg/log"
	"istio.io/pkg/version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	istioNamespace             = "istio-system"
	defaultInjectConfigMapName = "istio-sidecar-injector"
	injectConfigMapKey         = "config"
	valuesConfigMapKey         = "values"
	defaultMeshConfigMapName   = "istio"
	configMapKey               = "mesh"
)

type ConfigMap struct {
	client *kubernetes.Clientset
}

func NewConfigMap(client *kubernetes.Clientset) *ConfigMap {
	return &ConfigMap{
		client: client,
	}
}

func (c *ConfigMap) GetMeshConfigFromConfigMap(name string, injectConfigMapNames ...string) (meshConfig *meshconfig.MeshConfig, err error) {
	meshConfigMapName := defaultMeshConfigMapName
	if len(injectConfigMapNames) != 0 {
		meshConfigMapName = injectConfigMapNames[0]
	}
	meshConfigMap, err := c.client.CoreV1().ConfigMaps(istioNamespace).Get(meshConfigMapName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not read valid configmap %q from namespace %q: %v - "+
			"Use --meshConfigFile or re-run "+name+" with `-i <istioSystemNamespace> and ensure valid MeshConfig exists",
			meshConfigMapName, istioNamespace, err)
	}
	// values in the data are strings, while proto might use a
	// different data type.  therefore, we have to get a value by a
	// key
	configYaml, exists := meshConfigMap.Data[configMapKey]
	if !exists {
		return nil, fmt.Errorf("missing configuration map key %q", configMapKey)
	}
	cfg, err := mesh.ApplyMeshConfigDefaults(configYaml)
	if err != nil {
		err = multierror.Append(err, fmt.Errorf("istioctl version %s cannot parse mesh config.  Install istioctl from the latest Istio release",
			version.Info.Version))
	}
	return cfg, err
}

//getInjectConfigFromConfigMap get inject configMap
func (c *ConfigMap) GetInjectConfigFromConfigMap(injectConfigMapNames ...string) (sidecarTemplate string, err error) {
	injectConfigMapName := defaultInjectConfigMapName
	if len(injectConfigMapNames) != 0 {
		injectConfigMapName = injectConfigMapNames[0]
	}
	meshConfigMap, err := c.client.CoreV1().ConfigMaps(istioNamespace).Get(injectConfigMapName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("could not find valid configmap %q from namespace  %q: %v - "+
			"Use --injectConfigFile or re-run kube-inject with `-i <istioSystemNamespace> and ensure istio-sidecar-injector configmap exists",
			injectConfigMapName, istioNamespace, err)
	}
	// values in the data are strings, while proto might use a
	// different data type.  therefore, we have to get a value by a
	// key
	injectData, exists := meshConfigMap.Data[injectConfigMapKey]
	if !exists {
		return "", fmt.Errorf("missing configuration map key %q in %q",
			injectConfigMapKey, injectConfigMapName)
	}
	var injectConfig istio_inject.Config
	if err := yaml.Unmarshal([]byte(injectData), &injectConfig); err != nil {
		return "", fmt.Errorf("unable to convert data from configmap %q: %v",
			injectConfigMapName, err)
	}
	log.Debugf("using inject template from configmap %q", injectConfigMapName)
	return injectConfig.Template, nil
}

func (c *ConfigMap) GetValuesFromConfigMap(injectConfigMapNames ...string) (valuesConfig string, err error) {
	injectConfigMapName := defaultInjectConfigMapName
	if len(injectConfigMapNames) != 0 {
		injectConfigMapName = injectConfigMapNames[0]
	}
	meshConfigMap, err := c.client.CoreV1().ConfigMaps(istioNamespace).Get(injectConfigMapName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("could not find valid configmap %q from namespace  %q: %v - "+
			"Use --valuesFile or re-run kube-inject with `-i <istioSystemNamespace> and ensure istio-sidecar-injector configmap exists",
			injectConfigMapName, istioNamespace, err)
	}

	valuesData, exists := meshConfigMap.Data[valuesConfigMapKey]
	if !exists {
		return "", fmt.Errorf("missing configuration map key %q in %q",
			valuesConfigMapKey, injectConfigMapName)
	}

	return valuesData, nil
}

func (c *ConfigMap) GetInjectConfig(injectConfigMapNames ...string) (sidecarTemplate *istio_inject.Config, err error) {
	injectConfigMapName := defaultInjectConfigMapName
	if len(injectConfigMapNames) != 0 {
		injectConfigMapName = injectConfigMapNames[0]
	}
	meshConfigMap, err := c.client.CoreV1().ConfigMaps(istioNamespace).Get(injectConfigMapName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not find valid configmap %q from namespace  %q: %v - "+
			"Use --injectConfigFile or re-run kube-inject with `-i <istioSystemNamespace> and ensure istio-sidecar-injector configmap exists",
			injectConfigMapName, istioNamespace, err)
	}
	// values in the data are strings, while proto might use a
	// different data type.  therefore, we have to get a value by a
	// key
	injectData, exists := meshConfigMap.Data[injectConfigMapKey]
	if !exists {
		return nil, fmt.Errorf("missing configuration map key %q in %q",
			injectConfigMapKey, injectConfigMapName)
	}
	var injectConfig istio_inject.Config
	if err := yaml.Unmarshal([]byte(injectData), &injectConfig); err != nil {
		return nil, fmt.Errorf("unable to convert data from configmap %q: %v",
			injectConfigMapName, err)
	}
	log.Debugf("using inject template from configmap %q", injectConfigMapName)
	return &injectConfig, nil
}