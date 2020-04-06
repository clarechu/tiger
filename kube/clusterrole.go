package kube

import (
	"k8s.io/api/rbac/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ClusterRole struct {
	clientset *kubernetes.Clientset
}

func NewClusterRole(clientset *kubernetes.Clientset) *ClusterRole {
	return &ClusterRole{clientset: clientset}
}

func (c *ClusterRole) Create() (clusterRoleBind *v1beta1.ClusterRole, err error) {
	clusterRoleBind = &v1beta1.ClusterRole{
		ObjectMeta: v1.ObjectMeta{
			Name: "",
		},
		Rules: []v1beta1.PolicyRule{
			{
			},
		},
	}
	clusterRoleBind, err = c.clientset.RbacV1beta1().ClusterRoles().Create(clusterRoleBind)
	return
}

func (c *ClusterRole) Delete(name string) (err error) {
	ops := &meta_v1.DeleteOptions{}
	err = c.clientset.RbacV1beta1().ClusterRoles().Delete(name, ops)
	return
}

func (c *ClusterRole) Update() (clusterRoleBind *v1beta1.ClusterRole, err error) {
	clusterRoleBind = &v1beta1.ClusterRole{
		ObjectMeta: v1.ObjectMeta{
			Name: "",
		},
		Rules: []v1beta1.PolicyRule{
			{
			},
		},
	}
	clusterRoleBind, err = c.clientset.RbacV1beta1().ClusterRoles().Update(clusterRoleBind)
	return
}

func (c *ClusterRole) Get(name string) (clusterRoleBind *v1beta1.ClusterRole, err error) {
	ops := meta_v1.GetOptions{}
	clusterRoleBind, err = c.clientset.RbacV1beta1().ClusterRoles().Get(name, ops)
	return
}

func (c *ClusterRole) List(name string) (clusterRoleBinds *v1beta1.ClusterRoleList, err error) {
	ops := meta_v1.ListOptions{}
	clusterRoleBinds, err = c.clientset.RbacV1beta1().ClusterRoles().List(ops)
	return
}
