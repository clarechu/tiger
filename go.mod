module github.com/ClareChu/tiger

go 1.14

require (
	github.com/ClareChu/gorequest v0.2.19
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869
	github.com/ghodss/yaml v1.0.0
	github.com/go-cmd/cmd v1.2.0
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
	istio.io/api v0.0.0-20200609235057-2f6a9b136356
	istio.io/client-go v0.0.0-20200610000813-8e1d4bd9cbca
	istio.io/istio v0.0.0-20200610054140-4489e711b34a
	istio.io/pkg v0.0.0-20200511212725-7bfbbf968c23
	k8s.io/api v0.18.3
	k8s.io/apiextensions-apiserver v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v0.18.3
	k8s.io/klog v1.0.0
)

replace (
	istio.io/api => istio.io/api v0.0.0-20200529165953-72dad51d4ffc
	istio.io/istio => github.com/istio/istio v0.0.0-20200609030148-aa3826a8b2d0
	istio.io/pkg => istio.io/pkg v0.0.0-20200511212725-7bfbbf968c23
	k8s.io/api => k8s.io/api v0.18.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.1
	k8s.io/client-go => k8s.io/client-go v0.18.0
)
