module github.com/ClareChu/tiger

go 1.14

require (
	github.com/ClareChu/gorequest v0.2.19
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869
	github.com/ghodss/yaml v1.0.0
	github.com/go-cmd/cmd v1.2.0
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/martian v2.1.0+incompatible
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/hashicorp/go-multierror v1.0.0
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/kr/pretty v0.2.0
	github.com/magiconair/properties v1.8.1
	github.com/prometheus/common v0.6.0
	github.com/spf13/cobra v1.0.0
	gopkg.in/yaml.v2 v2.2.8
	gotest.tools v2.2.0+incompatible
	istio.io/api v0.0.0-20200325005357-8217d7225b6d
	istio.io/client-go v0.0.0-20200316192452-065c59267750
	istio.io/istio v0.0.0-20200323201801-9d07e185b0dd
	istio.io/pkg v0.0.0-20200204185554-47b6d38ec784
	k8s.io/api v0.17.4
	k8s.io/apiextensions-apiserver v0.17.2
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v0.17.4
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	istio.io/api => istio.io/api v0.0.0-20200316215140-da46fe8e25be
	istio.io/client-go => istio.io/client-go v0.0.0-20200316192452-065c59267750
	istio.io/istio => github.com/istio/istio v0.0.0-20200323201801-9d07e185b0dd
	k8s.io/api => k8s.io/api v0.17.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.2
	k8s.io/utils => k8s.io/utils v0.0.0-20191114184206-e782cd3c129f
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.3.0
)
