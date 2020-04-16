module github.com/ClareChu/tiger

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/go-cmd/cmd v1.2.0
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/kr/pretty v0.2.0
	github.com/magiconair/properties v1.8.0
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/prometheus/common v0.4.0
	github.com/spf13/cobra v0.0.6
	gopkg.in/inf.v0 v0.9.1 // indirect
	istio.io/client-go v0.0.0-20200325170329-dc00bbff4229
	k8s.io/api v0.17.4
	k8s.io/apiextensions-apiserver v0.17.0
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20191004102349-159aefb8556b
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191004105649-b14e3c49469a
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191004074956-c5d2f014d689
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.3.0
)
