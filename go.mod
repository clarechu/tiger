module github.com/ClareChu/tiger

go 1.14

require (
	cloud.google.com/go v0.64.0 // indirect
	github.com/ClareChu/gorequest v0.2.19
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869
	github.com/confluentinc/confluent-kafka-go v1.6.1 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-cmd/cmd v1.2.0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/martian v2.1.0+incompatible
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/hashicorp/go-multierror v1.1.0
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/kr/pretty v0.2.0
	github.com/magiconair/properties v1.8.1
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/prometheus/common v0.6.0
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/spf13/cobra v1.0.0
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	golang.org/x/net v0.0.0-20201029221708-28c70e62bb1d // indirect
	golang.org/x/sys v0.0.0-20201029080932-201ba4db2418 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/genproto v0.0.0-20201030142918-24207fddd1c3 // indirect
	google.golang.org/grpc v1.33.1 // indirect
	gopkg.in/confluentinc/confluent-kafka-go.v1 v1.6.1
	gopkg.in/yaml.v2 v2.2.8
	gorm.io/driver/mysql v1.0.6
	gorm.io/gorm v1.21.9
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

replace github.com/spf13/viper => github.com/istio/viper v1.3.3-0.20190515210538-2789fed3109c

// Old version had no license
replace github.com/chzyer/logex => github.com/chzyer/logex v1.1.11-0.20170329064859-445be9e134b2

// Avoid pulling in incompatible libraries
replace github.com/docker/distribution => github.com/docker/distribution v2.7.1+incompatible

// Avoid pulling in kubernetes/kubernetes
replace github.com/Microsoft/hcsshim => github.com/Microsoft/hcsshim v0.8.8-0.20200421182805-c3e488f0d815

// Client-go does not handle different versions of mergo due to some breaking changes - use the matching version
replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.5

// See https://github.com/kubernetes/kubernetes/issues/92867, there is a bug in the library
replace github.com/evanphx/json-patch => github.com/evanphx/json-patch v0.0.0-20190815234213-e83c0a1c26c8
