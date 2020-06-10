package kube

import (
	"context"
	"k8s.io/api/rbac/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ClusterRole struct {
	clientSet *kubernetes.Clientset
}

func NewClusterRole(clientSet *kubernetes.Clientset) *ClusterRole {
	return &ClusterRole{clientSet: clientSet}
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
	clusterRoleBind, err = c.clientSet.RbacV1beta1().ClusterRoles().Create(context.TODO(), clusterRoleBind, meta_v1.CreateOptions{})
	return
}

func (c *ClusterRole) Delete(name string) (err error) {
	ops := meta_v1.DeleteOptions{}
	err = c.clientSet.RbacV1beta1().ClusterRoles().Delete(context.TODO(), name, ops)
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
	clusterRoleBind, err = c.clientSet.RbacV1beta1().ClusterRoles().Update(context.TODO(), clusterRoleBind, meta_v1.UpdateOptions{})
	return
}

func (c *ClusterRole) Get(name string) (clusterRoleBind *v1beta1.ClusterRole, err error) {
	ops := meta_v1.GetOptions{}
	clusterRoleBind, err = c.clientSet.RbacV1beta1().ClusterRoles().Get(context.TODO(), name, ops)
	return
}

func (c *ClusterRole) List(name string) (clusterRoleBinds *v1beta1.ClusterRoleList, err error) {
	ops := meta_v1.ListOptions{}
	clusterRoleBinds, err = c.clientSet.RbacV1beta1().ClusterRoles().List(context.TODO(), ops)
	return
}
