package kube

import (
	"k8s.io/api/rbac/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ClusterRoleBinding struct {
	clientSet *kubernetes.Clientset
}

func NewClusterRoleBinding(clientSet *kubernetes.Clientset) *ClusterRoleBinding {
	return &ClusterRoleBinding{clientSet: clientSet}
}

func (c *ClusterRoleBinding) Create(clusterRoleBind *v1beta1.ClusterRoleBinding) (err error) {
	clusterRoleBind, err = c.clientSet.RbacV1beta1().ClusterRoleBindings().Create(clusterRoleBind)
	return
}

func (c *ClusterRoleBinding) Delete(name string) (err error) {
	ops := &meta_v1.DeleteOptions{}
	err = c.clientSet.RbacV1beta1().ClusterRoleBindings().Delete(name, ops)
	return
}

func (c *ClusterRoleBinding) Update(clusterRoleBind *v1beta1.ClusterRoleBinding) (err error) {
	clusterRoleBind, err = c.clientSet.RbacV1beta1().ClusterRoleBindings().Update(clusterRoleBind)
	return
}

func (c *ClusterRoleBinding) Get(name string) (clusterRoleBind *v1beta1.ClusterRoleBinding, err error) {
	ops := meta_v1.GetOptions{}
	clusterRoleBind, err = c.clientSet.RbacV1beta1().ClusterRoleBindings().Get(name, ops)
	return
}

func (c *ClusterRoleBinding) List(name string) (clusterRoleBinds *v1beta1.ClusterRoleBindingList, err error) {
	ops := meta_v1.ListOptions{}
	clusterRoleBinds, err = c.clientSet.RbacV1beta1().ClusterRoleBindings().List(ops)
	return
}
