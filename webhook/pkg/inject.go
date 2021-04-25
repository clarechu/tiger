package pkg

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	meshconfig "istio.io/api/mesh/v1alpha1"
	"istio.io/istio/pkg/config/mesh"
	"istio.io/istio/pkg/kube/inject"
	"istio.io/pkg/log"
	"istio.io/pkg/version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

//Inject In order to take advantage of all of Istio’s features, pods in the mesh must be running an Istio sidecar proxy.
type Inject struct {
	ClientSet        *kubernetes.Clientset
	MeshConfigFile   string
	InjectConfigFile string
	ValuesFile       string
}

//todo Auto inject
// kubectl label namespace default istio-injection=enabled
func AutoInject() {

}

const (
	defaultMeshConfigMapName   = "istio"
	defaultIstioNamespace      = "istio-system"
	defaultInjectConfigMapName = "istio-sidecar-injector"
)

const (
	configMapKey       = "mesh"
	injectConfigMapKey = "config"
	valuesConfigMapKey = "values"
)

//todo Manual operation  inject
// istioctl kube-inject -f samples/sleep/sleep.yaml | kubectl apply -f -
func (i Inject) ManualInject(obj runtime.Object) (updates interface{}, err error) {
	var meshConfig *meshconfig.MeshConfig
	if i.MeshConfigFile != "" {
		if meshConfig, err = mesh.ReadMeshConfig(i.MeshConfigFile); err != nil {
			return nil, err
		}
	} else {
		if meshConfig, err = getMeshConfigFromConfigMap(i.ClientSet, "kube-inject"); err != nil {
			return nil, err
		}
	}

	var sidecarTemplate string
	if i.InjectConfigFile != "" {
		injectionConfig, err := ioutil.ReadFile(i.InjectConfigFile) // nolint: vetshadow
		if err != nil {
			return nil, err
		}
		var injectConfig inject.Config
		if err := yaml.Unmarshal(injectionConfig, &injectConfig); err != nil {
			return nil, multierror.Append(err, fmt.Errorf("loading --injectConfigFile"))
		}
		sidecarTemplate = injectConfig.Template
	} else if sidecarTemplate, err = getInjectConfigFromConfigMap(i.ClientSet); err != nil {
		return nil, err
	}

	var valuesConfig string
	if i.ValuesFile != "" {
		valuesConfigBytes, err := ioutil.ReadFile(i.ValuesFile) // nolint: vetshadow
		if err != nil {
			return nil, err
		}
		valuesConfig = string(valuesConfigBytes)
	} else if valuesConfig, err = getValuesFromConfigMap(i.ClientSet); err != nil {
		return nil, err
	}
	return IntoResourceFile(sidecarTemplate, valuesConfig, meshConfig, obj)
}

// IntoResourceFile injects the istio proxy into the specified
// kubernetes YAML file.
//手动 inject
func IntoResourceFile(sidecarTemplate string, valuesConfig string, meshconfig *meshconfig.MeshConfig, raw runtime.Object) (updated interface{}, err error) {
	return inject.IntoObject(sidecarTemplate, valuesConfig, "", meshconfig, raw) // nolint: vetshadow
}

func getInjectConfigFromConfigMap(clientSet *kubernetes.Clientset) (string, error) {
	meshConfigMap, err := clientSet.CoreV1().ConfigMaps(defaultIstioNamespace).Get(context.TODO(), defaultInjectConfigMapName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("could not find valid configmap %q from namespace  %q: %v - "+
			"Use --injectConfigFile or re-run kube-inject with `-i <istioSystemNamespace> and ensure istio-sidecar-injector configmap exists",
			defaultInjectConfigMapName, defaultIstioNamespace, err)
	}
	// values in the data are strings, while proto might use a
	// different data type.  therefore, we have to get a value by a
	// key
	injectData, exists := meshConfigMap.Data[injectConfigMapKey]
	if !exists {
		return "", fmt.Errorf("missing configuration map key %q in %q",
			injectConfigMapKey, defaultInjectConfigMapName)
	}
	var injectConfig inject.Config
	if err := yaml.Unmarshal([]byte(injectData), &injectConfig); err != nil {
		return "", fmt.Errorf("unable to convert data from configmap %q: %v",
			defaultInjectConfigMapName, err)
	}
	log.Debugf("using inject template from configmap %q", defaultInjectConfigMapName)
	return injectConfig.Template, nil
}

func getMeshConfigFromConfigMap(clientSet *kubernetes.Clientset, command string) (*meshconfig.MeshConfig, error) {
	meshConfigMap, err := clientSet.CoreV1().ConfigMaps(defaultIstioNamespace).Get(context.TODO(), defaultMeshConfigMapName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not read valid configmap %q from namespace %q: %v - "+
			"Use --meshConfigFile or re-run "+command+" with `-i <istioSystemNamespace> and ensure valid MeshConfig exists",
			defaultMeshConfigMapName, defaultIstioNamespace, err)
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

// grabs the raw values from the ConfigMap. These are encoded as JSON.
func getValuesFromConfigMap(clientSet *kubernetes.Clientset) (string, error) {

	meshConfigMap, err := clientSet.CoreV1().ConfigMaps(defaultIstioNamespace).Get(context.TODO(), defaultInjectConfigMapName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("could not find valid configmap %q from namespace  %q: %v - "+
			"Use --valuesFile or re-run kube-inject with `-i <istioSystemNamespace> and ensure istio-sidecar-injector configmap exists",
			defaultInjectConfigMapName, defaultIstioNamespace, err)
	}

	valuesData, exists := meshConfigMap.Data[valuesConfigMapKey]
	if !exists {
		return "", fmt.Errorf("missing configuration map key %q in %q",
			valuesConfigMapKey, defaultInjectConfigMapName)
	}

	return valuesData, nil
}

//todo deployment set labels
// edit deployment labels istio-injection=enabled
func SetLabelsInject() {

}
