package pkg

import (
	"context"
	"errors"
	"fmt"
	deploymentutil "github.com/ClareChu/tiger/webhook/util"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	"sort"
)

type Uninject struct {
	ClientSet *kubernetes.Clientset
}

//Because of the removal of /rollback endpoint in apps/v1.Deployments, the example and kubectl, if switched to apps/v1.Deployments, need to do the rollback logic themselves. That includes:
//
//List all ReplicaSets the Deployment owns
//Find the ReplicaSet of a specific revision
//Copy that ReplicaSet's template back to the Deployment's template
//The rollback logic currently lives in Deployment controller code, which still uses extensions/v1beta1 Deployment client:
func (un *Uninject) rollback(name, namespace string) (err error) {
	deployments, err := un.ClientSet.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", name)})
	if err != nil {
		klog.Errorf("get deployment error:%v", err)
		return
	}

	for _, deployment := range deployments.Items {
		list, err := un.ClientSet.AppsV1().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{
			LabelSelector: fmt.Sprintf("app=%s,version=%s", name, deployment.Labels["version"]),
		})
		if err != nil {
			klog.Errorf("get replicasets error:%v", err)
			return err
		}
		_, err = un.rollbackToTemplate(&deployment, list.Items)
	}
	return err
}

func (un *Uninject) unInjectDeployment(name, namespace string) (err error) {
	klog.Errorf("unInject deployment name:%v, namespace:%v", name, namespace)
	deployments, err := un.ClientSet.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", name)})
	if err != nil {
		klog.Errorf("get deployment error:%v", err)
		return
	}
	for _, deployment := range deployments.Items {
		_, err = un.rollbackDeployment(&deployment)
		if err != nil {
			klog.Errorf("rollback deployment error:%v", err)
			return err
		}
	}
	return err
}

func (un *Uninject) unInjectDeploymentName(name, namespace string) (err error) {
	klog.Errorf("unInject deployment name:%v, namespace:%v", name, namespace)
	deployment, err := un.ClientSet.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		klog.Errorf("get deployment error:%v", err)
		return
	}
	_, err = un.rollbackDeployment(deployment)
	if err != nil {
		klog.Errorf("rollback deployment error:%v", err)
		return err
	}
	return err
}

func (un *Uninject) rollbackDeployment(d *v1.Deployment) (bool, error) {
	d.Spec.Template.Annotations = nil
	// delete Containers
	for index, container := range d.Spec.Template.Spec.Containers {
		if container.Name == "istio-proxy" {
			d.Spec.Template.Spec.Containers = append(
				d.Spec.Template.Spec.Containers[:index],
				d.Spec.Template.Spec.Containers[index+1:]...)
		}
	}
	// delete initContainers
	for index, container := range d.Spec.Template.Spec.InitContainers {
		if container.Name == "istio-init" {
			d.Spec.Template.Spec.InitContainers = append(
				d.Spec.Template.Spec.InitContainers[:index],
				d.Spec.Template.Spec.InitContainers[index+1:]...)
		}
	}
	// delete volume
	for index, volume := range d.Spec.Template.Spec.Volumes {
		if volume.Name == "istio-envoy" {
			d.Spec.Template.Spec.Volumes = append(
				d.Spec.Template.Spec.Volumes[:index],
				d.Spec.Template.Spec.Volumes[index+1:]...)
		}
	}
	for index, volume := range d.Spec.Template.Spec.Volumes {
		if volume.Name == "istio-data" {
			d.Spec.Template.Spec.Volumes = append(
				d.Spec.Template.Spec.Volumes[:index],
				d.Spec.Template.Spec.Volumes[index+1:]...)
		}
	}
	for index, volume := range d.Spec.Template.Spec.Volumes {

		if volume.Name == "istio-podinfo" {
			d.Spec.Template.Spec.Volumes = append(
				d.Spec.Template.Spec.Volumes[:index],
				d.Spec.Template.Spec.Volumes[index+1:]...)
		}
	}
	for index, volume := range d.Spec.Template.Spec.Volumes {

		if volume.Name == "istiod-ca-cert" {
			d.Spec.Template.Spec.Volumes = append(
				d.Spec.Template.Spec.Volumes[:index],
				d.Spec.Template.Spec.Volumes[index+1:]...)
		}
	}
	d.Spec.Template.Spec.DNSConfig = nil
	d.Spec.Template.Spec.SecurityContext = nil
	return true, un.updateDeploymentAndClearRollbackTo(d)
}

// rollbackToTemplate compares the templates of the provided deployment and replica set and
// updates the deployment with the replica set template in case they are different. It also
// cleans up the rollback spec so subsequent requeues of the deployment won't end up in here.
func (un *Uninject) rollbackToTemplate(d *v1.Deployment, rs []v1.ReplicaSet) (bool, error) {
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].CreationTimestamp.UnixNano() > rs[j].CreationTimestamp.UnixNano()
	})
	performedRollback := false
	replicaSet := v1.ReplicaSet{}
	for _, r := range rs {
		for _, initContainer := range r.Spec.Template.Spec.InitContainers {
			if initContainer.Name == "istio-init" {
				performedRollback = true
				break
			} else {
				klog.V(4).Infof("Rolling back to a revision that contains the same template as current deployment %q, skipping rollback...", d.Name)
			}
		}
		if len(r.Spec.Template.Spec.InitContainers) == 0 || !performedRollback {
			replicaSet = r
			//replicaSet.Spec.Template.ObjectMeta.Labels = replicaSet.Spec.Selector.MatchLabels
			performedRollback = true
			break
		}
	}
	if performedRollback {
		klog.V(4).Infof("Rolling back deployment %q to template spec %+v", d.Name, replicaSet.Spec.Template.Spec)
		deploymentutil.SetFromReplicaSetTemplate(d, replicaSet.Spec.Template)
		deploymentutil.SetDeploymentAnnotationsTo(d, &replicaSet)
	} else {
		return false, errors.New("current deployment not found sidecar")
	}
	return performedRollback, un.updateDeploymentAndClearRollbackTo(d)
}

// updateDeploymentAndClearRollbackTo sets .spec.rollbackTo to nil and update the input deployment
// It is assumed that the caller will have updated the deployment template appropriately (in case
// we want to rollback).
func (un *Uninject) updateDeploymentAndClearRollbackTo(d *v1.Deployment) error {
	klog.Infof("Cleans up rollbackTo of deployment %q", d.Name)
	_, err := un.ClientSet.AppsV1().Deployments(d.Namespace).Update(context.TODO(), d, metav1.UpdateOptions{})
	return err
}
