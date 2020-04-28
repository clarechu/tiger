# 变态的控制器 （Dynamic Admission Control）

让我们先来看官网是如何解释的，[Using Admission Controllers](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/)

ValidatingWebhookConfiguration

### Mutation
```bash
$ kc get mutatingwebhookconfiguration --all-namespaces -o yaml
```
Admission 的变更功能可以在资源生成之前进行修改。在 Admission 链条上可以多次修改同一个字段，因此插件不是无序的。

PodNodeSelector就是一个例子，
他使用 Namespace 的一个注解
namespace.annotations[“scheduler.alpha.kubernetes.io/node-selector”]，
在其中查找标签选择器，并将其加入pod.spec.nodeselector字段。
这一功能限制某个 Namespace 的 Pod，只能运行在指定节点上，同 Taint 的功能正好相反（也是 Admission 插件）。

### Validate

```bash
$ kc get validatingwebhookconfiguration istiod-istio-system -o yaml
```
Admission 的验证阶段用来对特定 API 资源进行验证。这一阶段会在所有变更执行结束之后，API 资源不再发生变化的情况下进行。

PodNodeSelector插件同样演示了这一活动，他确保 Pod 的spec.nodeSelector字段符合该 Namespace 的限制。如果变更链上有其他 Admission 尝试在PodNodeSelector之后修改了资源的spec.nodeSelector，就会在验证阶段因为不符合限制而被拒绝创建。

```yaml
apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  name: namespacereservations.admission.online.openshift.io
webhooks:
- name: namespacereservations.admission.online.openshift.io
  clientConfig:
    service:
      namespace: default
     name: kubernetes
    path: /apis/admission.online.openshift.io/v1alpha1/namespacereservations
   caBundle: KUBE_CA_HERE
 rules:
 - operations:
   - CREATE
   apiGroups:
   - ""
   apiVersions:
   - "*"
   resources:
   - namespaces
 failurePolicy: Fail
```
07 dispatcher.go:136] Failed calling webhook, failing open sidecar-injector.istio.io: failed calling webhook "sidecar-injector.istio.io": Post https://sidecar-injector.istio-system.svc:443/mutate?timeout=30s: http: server gave HTTP response to HTTPS client
4月 28 01:14:13 localhost.localdomain kube-apiserver[1107]: E0428 01:14:13.539384    1107 dispatcher.go:137] failed calling webhook "sidecar-injector.istio.io": Post https://sidecar-injector.istio-system.svc:443/mutate?timeout=30s: http: server gave HTTP response to HTTPS client