package kube

//CustomResourceDefinition 添加crd 资源
type CustomResourceDefinition struct {
	//crd全名
	FullCRDName string `json:"fullCrdName, omitempty"`
	//crd 组名
	CRDGroup    string `json:"crdGroup, omitempty"`
	CRDPlural   string `json:"crdPlural, omitempty"`
	Kind        string `json:"kind, omitempty"`
	CRDVersion  string `json:"crdVersion, omitempty"`
}
