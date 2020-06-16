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
	list, err := un.ClientSet.AppsV1().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", name),
	})
	if err != nil {
		klog.Errorf("get replicasets error:%v", err)
		return
	}
	deployment, err := un.ClientSet.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		klog.Errorf("get deployment error:%v", err)
		return
	}
	_, err = un.rollbackToTemplate(deployment, list.Items)
	return err
}

/*
// rollback the deployment to the specified revision. In any case cleanup the rollback spec.
func (dc *DeploymentController) rollback(d *extensions.Deployment, rsList []*extensions.ReplicaSet, podMap map[types.UID]*v1.PodList) error {
	newRS, allOldRSs, err := dc.getAllReplicaSetsAndSyncRevision(d, rsList, podMap, true)
	if err != nil {
		return err
	}

	allRSs := append(allOldRSs, newRS)
	toRevision := &d.Spec.RollbackTo.Revision
	// If rollback revision is 0, rollback to the last revision
	if *toRevision == 0 {
		if *toRevision = deploymentutil.LastRevision(allRSs); *toRevision == 0 {
			// If we still can't find the last revision, gives up rollback
			dc.emitRollbackWarningEvent(d, deploymentutil.RollbackRevisionNotFound, "Unable to find last revision.")
			// Gives up rollback
			return dc.updateDeploymentAndClearRollbackTo(d)
		}
	}
	for _, rs := range allRSs {
		v, err := deploymentutil.Revision(rs)
		if err != nil {
			glog.V(4).Infof("Unable to extract revision from deployment's replica set %q: %v", rs.Name, err)
			continue
		}
		if v == *toRevision {
			glog.V(4).Infof("Found replica set %q with desired revision %d", rs.Name, v)
			// rollback by copying podTemplate.Spec from the replica set
			// revision number will be incremented during the next getAllReplicaSetsAndSyncRevision call
			// no-op if the spec matches current deployment's podTemplate.Spec
			performedRollback, err := dc.rollbackToTemplate(d, rs)
			if performedRollback && err == nil {
				dc.emitRollbackNormalEvent(d, fmt.Sprintf("Rolled back deployment %q to revision %d", d.Name, *toRevision))
			}
			return err
		}
	}
	dc.emitRollbackWarningEvent(d, deploymentutil.RollbackRevisionNotFound, "Unable to find the revision to rollback to.")
	// Gives up rollback
	return dc.updateDeploymentAndClearRollbackTo(d)
}*/

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
