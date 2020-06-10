package cache

import (
	"context"
	"fmt"
	"github.com/ClareChu/tiger/kube"
	"github.com/ClareChu/tiger/webhook/cache/client"
	"istio.io/pkg/log"
	corev1 "k8s.io/api/core/v1"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strings"
	"sync"
)

// The valid auth strategies and values for cookie handling
const (
	// These constants are used for external services auth (Prometheus, Grafana ...) ; not for Kiali auth
	MesherLabelSelector = "app=mesher"
	MeshNamespace       = "service-mesh"
	AuthTypeBasic       = "basic"
	AuthTypeBearer      = "bearer"
	AuthTypeNone        = "none"
)

type Config struct {
	Name                  string                  `json:"name" yaml:"name,omitempty"`
	Version               string                  `json:"version,omitempty"`
	IstiodName            string                  `json:"istiodName,omitempty"`
	InCluster             bool                    `json:"inCluster,omitempty" yaml:"in_cluster"`
	MeshNamespace         string                  `json:"meshNamespace" yaml:"mesh_namespace,omitempty"`
	Clusters              map[string]*Cluster     `json:"clusters" yaml:"clusters,omitempty"`
	API                   ApiConfig               `json:"apiConfig" yaml:"api,omitempty"`
	Prometheus            PrometheusConfig        `yaml:"prometheus,omitempty"`
	ApiExtensionClientSet *apiextension.Clientset `json:"apiExtensionClientSet,omitempty" yaml:"api_extension_client_set,omitempty"`
	K8sClientSet          *kubernetes.Clientset   `json:"k8SClientSet, omitempty" yaml:"k8s_client_set, omitempty"`
	IstioNamespace        string                  `yaml:"istio_namespace,omitempty"`
	IstioLabels           IstioLabels             `yaml:"istio_labels,omitempty"`
}

// IstioLabels holds configuration about the labels required by Istio
type IstioLabels struct {
	AppLabelName     string `yaml:"app_label_name,omitempty" json:"appLabelName"`
	VersionLabelName string `yaml:"version_label_name,omitempty" json:"versionLabelName"`
}

// PrometheusConfig describes configuration of the Prometheus component
type PrometheusConfig struct {
	Auth             Auth   `yaml:"auth,omitempty"`
	CustomMetricsURL string `yaml:"custom_metrics_url,omitempty"`
	URL              string `yaml:"url,omitempty"`
}

type KubeClientConfig struct {
}

// Auth provides authentication data for external services
type Auth struct {
	CAFile             string `yaml:"ca_file"`
	InsecureSkipVerify bool   `yaml:"insecure_skip_verify"`
	Password           string `yaml:"password"`
	Token              string `yaml:"token"`
	Type               string `yaml:"type"`
	UseKialiToken      bool   `yaml:"use_kiali_token"`
	Username           string `yaml:"username"`
}

type Cluster struct {
	Prometheus PrometheusConfig `yaml:"prometheus,omitempty"`
	Installed  bool             `json:"installed"`
	Kiali      KialiConfig      `json:"kiali,omitempty"`
	KubeConfig *rest.Config     `json:"kubeConfig,omitempty" yaml:"kube_config,omitempty"`
}

type KialiConfig struct {
}

// ApiConfig contains API specific configuration.
type ApiConfig struct {
	Namespaces ApiNamespacesConfig
}

// ApiNamespacesConfig provides a list of regex strings defining namespaces to blacklist.
type ApiNamespacesConfig struct {
	Exclude       []string `yaml:"exclude, omitempty" json:"exclude,omitempty"`
	LabelSelector string   `yaml:"label_selector,omitempty" json:"labelSelector"`
}

var configuration Config

var rwMutex sync.RWMutex

//update the global conf
func Set(conf *Config) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	configuration = *conf
}

// Get the global Config
func Get() (conf *Config) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	config := configuration
	return &config
}

//NewConfig init global config
func NewConfig() (c *Config, err error) {
	kubeClientSet, err := client.GetDefaultK8sClientSet()
	if err != nil {
		return c, err
	}
	apiExtensionClientSet, err := client.GetDefaultApiExtensionClientSet()
	if err != nil {
		return c, err
	}
	ops := metav1.ListOptions{LabelSelector: MesherLabelSelector}
	list, err := kubeClientSet.CoreV1().ConfigMaps(MeshNamespace).List(context.TODO(), ops)
	if err != nil {
		return c, err
	}
	clusters := map[string]*Cluster{}
	for _, cm := range list.Items {
		conf := cm.BinaryData["kubeConfig"]
		api, err := client.LoadFromFile(conf)
		if err != nil {
			return c, err
		}
		config := &rest.Config{
			Host: api.Clusters[api.Contexts[api.CurrentContext].Cluster].Server,
			TLSClientConfig: rest.TLSClientConfig{
				CAData:   api.Clusters[api.Contexts[api.CurrentContext].Cluster].CertificateAuthorityData,
				CertData: api.AuthInfos[api.Contexts[api.CurrentContext].AuthInfo].ClientCertificateData,
				KeyData:  api.AuthInfos[api.Contexts[api.CurrentContext].AuthInfo].ClientKeyData,
			},
			//Timeout: 5 * time.Second,
		}
		installed := true
		address, err := GetPromAddress(config)
		if err != nil {
			installed = false
		}
		cluster := &Cluster{
			KubeConfig: config,
			Prometheus: PrometheusConfig{
				URL: address,
			},
			Installed: installed,
		}
		clusters[cm.Name] = cluster
	}
	c = &Config{
		Name:           "mesher",
		IstioNamespace: "istio-system",
		MeshNamespace:  "service-mesh",
		InCluster:      true,
		IstiodName:     "discovery",
		Prometheus: PrometheusConfig{
			Auth: Auth{
				Type: AuthTypeNone,
			},
			CustomMetricsURL: "http://prometheus.istio-system:9090",
			URL:              "http://prometheus.istio-system:9090",
		},
		K8sClientSet:          kubeClientSet,
		ApiExtensionClientSet: apiExtensionClientSet,
		IstioLabels: IstioLabels{
			AppLabelName:     "app",
			VersionLabelName: "version",
		},
		API: ApiConfig{
			Namespaces: ApiNamespacesConfig{
				Exclude: []string{
					"istio-operator",
					"kube.*",
					"openshift.*",
					"ibm.*",
					"kial-operator",
				},
			},
		},
		Clusters: clusters,
	}
	return c, nil
}

func AddCluster(configMaps *corev1.ConfigMap) (err error) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	conf := configMaps.BinaryData["kubeConfig"]
	api, err := client.LoadFromFile(conf)
	if err != nil {
		return err
	}
	config := &rest.Config{
		Host: api.Clusters[api.Contexts[api.CurrentContext].Cluster].Server,
		TLSClientConfig: rest.TLSClientConfig{
			CAData:   api.Clusters[api.Contexts[api.CurrentContext].Cluster].CertificateAuthorityData,
			CertData: api.AuthInfos[api.Contexts[api.CurrentContext].AuthInfo].ClientCertificateData,
			KeyData:  api.AuthInfos[api.Contexts[api.CurrentContext].AuthInfo].ClientKeyData,
		},
		//Timeout: 5 * time.Second,
	}
	installed := true
	address, err := GetPromAddress(config)
	if err != nil {
		installed = false
		log.Error("prometheus address not found")
	}
	cluster := &Cluster{
		Installed:  installed,
		KubeConfig: config,
		Prometheus: PrometheusConfig{
			URL: address,
		},
	}
	configuration.Clusters[configMaps.Name] = cluster
	return nil
}

func UpdateClusterPrometheus(name string) (address string, err error) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	installed := true
	address, err = GetPromAddress(configuration.Clusters[name].KubeConfig)
	if err != nil {
		installed = false
		return address, err
	}
	cluster := &Cluster{
		KubeConfig: configuration.Clusters[name].KubeConfig,
		Installed:  installed,
		Prometheus: PrometheusConfig{
			URL: address,
		},
	}
	configuration.Clusters[name] = cluster
	return address, nil
}

func DeleteCluster(name string) (err error) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	delete(configuration.Clusters, name)
	return nil
}

//GetKialiAddress 获取集群kiali的集群地址
func GetPromAddress(config *rest.Config) (address string, err error) {
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", err
	}
	service, err := clientSet.CoreV1().Services("istio-system").Get(context.TODO(), "prometheus", metav1.GetOptions{})
	if err != nil {
		return
	}
	//ip, _, err := net.SplitHostPort(config.Host)
	ip := strings.Split(strings.Split(config.Host, "://")[1], ":")[0]
	if service.Spec.Type == corev1.ServiceTypeNodePort {
		return fmt.Sprintf("http://%s:%d", ip, service.Spec.Ports[0].NodePort), nil
	}
	service.Spec.Type = corev1.ServiceTypeNodePort
	service, err = kube.NewService(clientSet).Update("istio-system", service)
	if err != nil {
		return
	}
	return fmt.Sprintf("http://%s:%d", ip, service.Spec.Ports[0].NodePort), nil
}
