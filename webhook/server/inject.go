package server

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ClareChu/tiger/istio"
	"github.com/ClareChu/tiger/kube/client"
	"io/ioutil"
	"istio.io/api/annotation"
	meshconfig "istio.io/api/mesh/v1alpha1"
	"istio.io/istio/pilot/cmd/pilot-agent/status"
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pkg/kube/inject"
	"istio.io/pkg/log"
	"k8s.io/api/admission/v1beta1"
	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"net/http"
	"sort"
	"strings"
)

var (
	runtimeScheme = runtime.NewScheme()
	codecs        = serializer.NewCodecFactory(runtimeScheme)
	deserializer  = codecs.UniversalDeserializer()

	// (https://github.com/kubernetes/kubernetes/issues/57982)
	defaulter = runtime.ObjectDefaulter(runtimeScheme)
)

// Retain deprecated hardcoded container and volumes names to aid in
// backwards compatible migration to the new SidecarInjectionStatus.
var (
	legacyInitContainerNames = []string{"istio-init", "enable-core-dump"}
	legacyContainerNames     = []string{"istio-proxy"}
	legacyVolumeNames        = []string{"istio-certs", "istio-envoy"}
)

// Webhook Server parameters
type WhSvrParameters struct {
	port           int    // webhook server port
	certFile       string // path to the x509 certificate for https
	keyFile        string // path to the x509 private key matching `CertFile`
	sidecarCfgFile string // path to sidecar injector configuration file
}

type WebhookServer struct {
	//SidecarConfig *Config
	Server                 *http.Server
	ValuesConfig           string
	MeshConfig             *meshconfig.MeshConfig
	SidecarTemplateVersion string
	Config                 *inject.Config
	ClientSet              *kubernetes.Clientset
}

func init() {
	_ = corev1.AddToScheme(runtimeScheme)
	_ = admissionregistrationv1beta1.AddToScheme(runtimeScheme)
	// defaulting with webhooks:
	// https://github.com/kubernetes/kubernetes/issues/57982
	_ = v1beta1.AddToScheme(runtimeScheme)
}

// Config specifies the sidecar injection configuration This includes
// the sidecar template and cluster-side injection policy. It is used
// by kube-inject, sidecar injector, and http endpoint.
type Config struct {
	Policy inject.InjectionPolicy `json:"policy"`

	// Template is the templated version of `SidecarInjectionSpec` prior to
	// expansion over the `SidecarTemplateData`.
	Template string `json:"template"`

	// NeverInjectSelector: Refuses the injection on pods whose labels match this selector.
	// It's an array of label selectors, that will be OR'ed, meaning we will iterate
	// over it and stop at the first match
	// Takes precedence over AlwaysInjectSelector.
	NeverInjectSelector []metav1.LabelSelector `json:"neverInjectSelector"`

	// AlwaysInjectSelector: Forces the injection on pods whose labels match this selector.
	// It's an array of label selectors, that will be OR'ed, meaning we will iterate
	// over it and stop at the first match
	AlwaysInjectSelector []metav1.LabelSelector `json:"alwaysInjectSelector"`

	// InjectedAnnotations are additional annotations that will be added to the pod spec after injection
	// This is primarily to support PSP annotations.
	InjectedAnnotations map[string]string `json:"injectedAnnotations"`
}

func (wh *WebhookServer) Inject(resp http.ResponseWriter, req *http.Request) {
	log.Info("sidecar inject start...")
	var body []byte
	if req.Body != nil {
		if data, err := ioutil.ReadAll(req.Body); err == nil {
			body = data
		}
	}
	if len(body) == 0 {
		log.Error("body nil")
		http.Error(resp, "no body found", http.StatusBadRequest)
		return
	}
	// verify the content type is accurate
	contentType := req.Header.Get("Content-Type")
	if contentType != "application/json" {
		log.Infof("contentType :%v", contentType)
		http.Error(resp, "invalid Content-Type, want `application/json`", http.StatusUnsupportedMediaType)
		return
	}
	wh.SidecarTemplateVersion = sidecarTemplateVersionHash(wh.Config.Template)
	var reviewResponse *v1beta1.AdmissionResponse
	ar := v1beta1.AdmissionReview{}
	if _, _, err := deserializer.Decode(body, nil, &ar); err != nil {
		//handleError(fmt.Sprintf("Could not decode body: %v", err))
		log.Infof("body decode error:%v", err)
		reviewResponse = toAdmissionResponse(err)
	} else {
		reviewResponse = wh.inject(&ar)
	}
	response := v1beta1.AdmissionReview{}
	if reviewResponse != nil {
		response.Response = reviewResponse
		if ar.Request != nil {
			response.Response.UID = ar.Request.UID
		}
	}
	re, err := json.Marshal(response)
	if err != nil {
		log.Infof("Could not encode response: %v", err)
		http.Error(resp, fmt.Sprintf("could not encode response: %v", err), http.StatusInternalServerError)
	}
	if _, err := resp.Write(re); err != nil {
		log.Infof("Could not write response: %v", err)
		http.Error(resp, fmt.Sprintf("could not write response: %v", err), http.StatusInternalServerError)
	}
}

// helper function to generate a template version identifier from a
// hash of the un-executed template contents.
func sidecarTemplateVersionHash(in string) string {
	hash := sha256.Sum256([]byte(in))
	return hex.EncodeToString(hash[:])
}

func (wh *WebhookServer) DefaultBuild() (err error) {
	clientSet, err := client.GetDefaultK8sClientSet()
	if err != nil {
		log.Infof("get client set err:%v", err)
		return err
	}
	configmap := istio.NewConfigMap(clientSet)
	//configmap get meshConfig
	meshConfig, err := configmap.GetMeshConfigFromConfigMap("kube-inject")
	if err != nil {
		log.Infof("get client set err:%v", err)
		return
	}
	//configmap get valueConfig
	valueConfig, err := configmap.GetValuesFromConfigMap()
	if err != nil {
		log.Infof("get valueConfig err:%v", err)
		return
	}
	wh.MeshConfig = meshConfig
	wh.ValuesConfig = valueConfig
	c, err := configmap.GetInjectConfig()
	if err != nil {
		log.Infof("GetInjectConfig err:%v", err)
		return
	}
	wh.Config = c
	return nil
}

func (wh *WebhookServer) inject(ar *v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	log.Info("web hook inject")
	req := ar.Request
	var pod corev1.Pod
	if err := json.Unmarshal(req.Object.Raw, &pod); err != nil {
		return toAdmissionResponse(err)
	}
	log.Infof("pod describe: %v", pod)
	// Deal with potential empty fields, e.g., when the pod is created by a deployment
	podName := pod.Name
	if pod.ObjectMeta.Namespace == "" {
		pod.ObjectMeta.Namespace = req.Namespace
	}
	// pod Labels istio-injection disable
	if pod.Labels["istio-injection"] != "enabled" {
		log.Info("this pod not found labels key : istio-injection, value: enabled")
		log.Info("*** continue ***")
		return &v1beta1.AdmissionResponse{
			Allowed: true,
		}
	}
	// pod Already sidecar
	for _, container := range pod.Spec.InitContainers {
		if container.Name == "istio-init" {
			log.Info("pod Already injection sidecar")
			log.Info("*** continue ***")
			return &v1beta1.AdmissionResponse{
				Allowed: true,
			}
		}
	}
	log.Infof("AdmissionReview for Kind=%v Namespace=%v Name=%v (%v) UID=%v Rfc6902PatchOperation=%v UserInfo=%v",
		req.Kind, req.Namespace, req.Name, podName, req.UID, req.Operation, req.UserInfo)
	log.Infof("Object: %v", string(req.Object.Raw))
	log.Infof("OldObject: %v", string(req.OldObject.Raw))

	// due to bug https://github.com/kubernetes/kubernetes/issues/57923,
	// k8s sa jwt token volume mount file is only accessible to root user, not istio-proxy(the user that istio proxy runs as).
	// workaround by https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-pod
	if wh.MeshConfig.SdsUdsPath != "" {
		var grp = int64(1337)
		if pod.Spec.SecurityContext == nil {
			pod.Spec.SecurityContext = &corev1.PodSecurityContext{
				FSGroup: &grp,
			}
		} else {
			pod.Spec.SecurityContext.FSGroup = &grp
		}
	}

	// try to capture more useful namespace/name info for deployments, etc.
	// TODO(dougreid): expand to enable lookup of OWNERs recursively a la kubernetesenv
	deployMeta := pod.ObjectMeta.DeepCopy()
	deployMeta.Namespace = req.Namespace

	typeMetadata := &metav1.TypeMeta{
		Kind:       "Pod",
		APIVersion: "v1",
	}

	if len(pod.GenerateName) > 0 {
		// if the pod name was generated (or is scheduled for generation), we can begin an investigation into the controlling reference for the pod.
		var controllerRef metav1.OwnerReference
		controllerFound := false
		for _, ref := range pod.GetOwnerReferences() {
			if *ref.Controller {
				controllerRef = ref
				controllerFound = true
				break
			}
		}
		if controllerFound {
			typeMetadata.APIVersion = controllerRef.APIVersion
			typeMetadata.Kind = controllerRef.Kind

			// heuristic for deployment detection
			if typeMetadata.Kind == "ReplicaSet" && strings.HasSuffix(controllerRef.Name, pod.Labels["pod-template-hash"]) {
				name := strings.TrimSuffix(controllerRef.Name, "-"+pod.Labels["pod-template-hash"])
				deployMeta.Name = name
				typeMetadata.Kind = "Deployment"
			} else {
				deployMeta.Name = controllerRef.Name
			}
		}
	}

	if deployMeta.Name == "" {
		// if we haven't been able to extract a deployment name, then just give it the pod name
		deployMeta.Name = pod.Name
	}

	spec, iStatus, err := inject.InjectionData(wh.Config.Template,
		wh.ValuesConfig,
		wh.SidecarTemplateVersion,
		typeMetadata,
		deployMeta,
		&pod.Spec,
		&pod.ObjectMeta,
		wh.MeshConfig,
		"") // nolint: lll
	if err != nil {
		return toAdmissionResponse(err)
	}

	annotations := map[string]string{annotation.SidecarStatus.Name: iStatus}

	// Add all additional injected annotations
	for k, v := range wh.Config.InjectedAnnotations {
		annotations[k] = v
	}

	patchBytes, err := createPatch(&pod,
		injectionStatus(&pod),
		annotations,
		spec,
		deployMeta.Name)
	if err != nil {
		return toAdmissionResponse(err)
	}

	log.Infof("AdmissionResponse: patch=%v\n", string(patchBytes))

	reviewResponse := v1beta1.AdmissionResponse{
		Allowed: true,
				Patch:   patchBytes,
				PatchType: func() *v1beta1.PatchType {
					pt := v1beta1.PatchTypeJSONPatch
					return &pt
				}(),
	}
	return &reviewResponse
}

func toAdmissionResponse(err error) *v1beta1.AdmissionResponse {
	return &v1beta1.AdmissionResponse{Result: &metav1.Status{Message: err.Error()}}
}

// It would be great to use https://github.com/mattbaird/jsonpatch to
// generate RFC6902 JSON patches. Unfortunately, it doesn't produce
// correct patches for object removal. Fortunately, our patching needs
// are fairly simple so generating them manually isn't horrible (yet).
type rfc6902PatchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

func createPatch(pod *corev1.Pod,
	prevStatus *inject.SidecarInjectionStatus,
	annotations map[string]string,
	sic *inject.SidecarInjectionSpec,
	workloadName string) ([]byte, error) {

	var patch []rfc6902PatchOperation

	// Remove any containers previously injected by kube-inject using
	// container and volume name as unique key for removal.
	patch = append(patch, removeContainers(pod.Spec.InitContainers, prevStatus.InitContainers, "/spec/initContainers")...)
	patch = append(patch, removeContainers(pod.Spec.Containers, prevStatus.Containers, "/spec/containers")...)
	patch = append(patch, removeVolumes(pod.Spec.Volumes, prevStatus.Volumes, "/spec/volumes")...)
	patch = append(patch, removeImagePullSecrets(pod.Spec.ImagePullSecrets, prevStatus.ImagePullSecrets, "/spec/imagePullSecrets")...)

	rewrite := inject.ShouldRewriteAppHTTPProbers(pod.Annotations, sic)
	addAppProberCmd := func() {
		if !rewrite {
			return
		}
		sidecar := inject.FindSidecar(sic.Containers)
		if sidecar == nil {
			log.Errorf("sidecar not found in the template, skip addAppProberCmd")
			return
		}
		// We don't have to escape json encoding here when using golang libraries.
		if prober := inject.DumpAppProbers(&pod.Spec); prober != "" {
			sidecar.Env = append(sidecar.Env, corev1.EnvVar{Name: status.KubeAppProberEnvName, Value: prober})
		}
	}
	addAppProberCmd()

	patch = append(patch, addContainer(pod.Spec.InitContainers, sic.InitContainers, "/spec/initContainers")...)
	patch = append(patch, addContainer(pod.Spec.Containers, sic.Containers, "/spec/containers")...)
	patch = append(patch, addVolume(pod.Spec.Volumes, sic.Volumes, "/spec/volumes")...)
	patch = append(patch, addImagePullSecrets(pod.Spec.ImagePullSecrets, sic.ImagePullSecrets, "/spec/imagePullSecrets")...)

	if sic.DNSConfig != nil {
		patch = append(patch, addPodDNSConfig(sic.DNSConfig, "/spec/dnsConfig")...)
	}

	if pod.Spec.SecurityContext != nil {
		patch = append(patch, addSecurityContext(pod.Spec.SecurityContext, "/spec/securityContext")...)
	}

	patch = append(patch, updateAnnotation(pod.Annotations, annotations)...)

	canonicalSvc, canonicalRev := extractCanonicalServiceLabels(pod.Labels, workloadName)
	patch = append(patch, addLabels(pod.Labels, map[string]string{
		//model.TLSModeLabelName:                       model.IstioMutualTLSModeLabel,
		model.IstioCanonicalServiceLabelName:         canonicalSvc,
		model.IstioCanonicalServiceRevisionLabelName: canonicalRev})...)

	if rewrite {
		patch = append(patch, createProbeRewritePatch(pod.Annotations, &pod.Spec, sic)...)
	}

	return json.Marshal(patch)
}

func extractCanonicalServiceLabels(podLabels map[string]string, workloadName string) (string, string) {
	return extractCanonicalServiceLabel(podLabels, workloadName), extractCanonicalServiceRevision(podLabels)
}

func extractCanonicalServiceLabel(podLabels map[string]string, workloadName string) string {
	if svc, ok := podLabels[model.IstioCanonicalServiceLabelName]; ok {
		return svc
	}

	if svc, ok := podLabels["app.kubernetes.io/name"]; ok {
		return svc
	}

	if svc, ok := podLabels["app"]; ok {
		return svc
	}

	return workloadName
}

func extractCanonicalServiceRevision(podLabels map[string]string) string {
	if rev, ok := podLabels[model.IstioCanonicalServiceRevisionLabelName]; ok {
		return rev
	}

	if rev, ok := podLabels["app.kubernetes.io/version"]; ok {
		return rev
	}

	if rev, ok := podLabels["version"]; ok {
		return rev
	}

	return "latest"
}

func injectionStatus(pod *corev1.Pod) *inject.SidecarInjectionStatus {
	var statusBytes []byte
	if pod.ObjectMeta.Annotations != nil {
		if value, ok := pod.ObjectMeta.Annotations[annotation.SidecarStatus.Name]; ok {
			statusBytes = []byte(value)
		}
	}

	// default case when injected pod has explicit status
	var iStatus inject.SidecarInjectionStatus
	if err := json.Unmarshal(statusBytes, &iStatus); err == nil {
		// heuristic assumes status is valid if any of the resource
		// lists is non-empty.
		if len(iStatus.InitContainers) != 0 ||
			len(iStatus.Containers) != 0 ||
			len(iStatus.Volumes) != 0 ||
			len(iStatus.ImagePullSecrets) != 0 {
			return &iStatus
		}
	}

	// backwards compatibility case when injected pod has legacy
	// status. Infer status from the list of legacy hardcoded
	// container and volume names.
	return &inject.SidecarInjectionStatus{
		InitContainers: legacyInitContainerNames,
		Containers:     legacyContainerNames,
		Volumes:        legacyVolumeNames,
	}
}

// JSONPatch `remove` is applied sequentially. Remove items in reverse
// order to avoid renumbering indices.
func removeContainers(containers []corev1.Container, removed []string, path string) (patch []rfc6902PatchOperation) {
	names := map[string]bool{}
	for _, name := range removed {
		names[name] = true
	}
	for i := len(containers) - 1; i >= 0; i-- {
		if _, ok := names[containers[i].Name]; ok {
			patch = append(patch, rfc6902PatchOperation{
				Op:   "remove",
				Path: fmt.Sprintf("%v/%v", path, i),
			})
		}
	}
	return patch
}

func removeVolumes(volumes []corev1.Volume, removed []string, path string) (patch []rfc6902PatchOperation) {
	names := map[string]bool{}
	for _, name := range removed {
		names[name] = true
	}
	for i := len(volumes) - 1; i >= 0; i-- {
		if _, ok := names[volumes[i].Name]; ok {
			patch = append(patch, rfc6902PatchOperation{
				Op:   "remove",
				Path: fmt.Sprintf("%v/%v", path, i),
			})
		}
	}
	return patch
}

func removeImagePullSecrets(imagePullSecrets []corev1.LocalObjectReference, removed []string, path string) (patch []rfc6902PatchOperation) {
	names := map[string]bool{}
	for _, name := range removed {
		names[name] = true
	}
	for i := len(imagePullSecrets) - 1; i >= 0; i-- {
		if _, ok := names[imagePullSecrets[i].Name]; ok {
			patch = append(patch, rfc6902PatchOperation{
				Op:   "remove",
				Path: fmt.Sprintf("%v/%v", path, i),
			})
		}
	}
	return patch
}

func addContainer(target, added []corev1.Container, basePath string) (patch []rfc6902PatchOperation) {
	saJwtSecretMountName := ""
	var saJwtSecretMount corev1.VolumeMount
	// find service account secret volume mount(/var/run/secrets/kubernetes.io/serviceaccount,
	// https://kubernetes.io/docs/reference/access-authn-authz/service-accounts-admin/#service-account-automation) from app container
	for _, add := range target {
		for _, vmount := range add.VolumeMounts {
			if vmount.MountPath == "/var/run/secrets/kubernetes.io/serviceaccount" {
				saJwtSecretMountName = vmount.Name
				saJwtSecretMount = vmount
			}
		}
	}
	first := len(target) == 0
	var value interface{}
	for _, add := range added {
		if add.Name == "istio-proxy" && saJwtSecretMountName != "" {
			// add service account secret volume mount(/var/run/secrets/kubernetes.io/serviceaccount,
			// https://kubernetes.io/docs/reference/access-authn-authz/service-accounts-admin/#service-account-automation) to istio-proxy container,
			// so that envoy could fetch/pass k8s sa jwt and pass to sds server, which will be used to request workload identity for the pod.
			add.VolumeMounts = append(add.VolumeMounts, saJwtSecretMount)
		}
		value = add
		path := basePath
		if first {
			first = false
			value = []corev1.Container{add}
		} else {
			path += "/-"
		}
		patch = append(patch, rfc6902PatchOperation{
			Op:    "add",
			Path:  path,
			Value: value,
		})
	}
	return patch
}

func addSecurityContext(target *corev1.PodSecurityContext, basePath string) (patch []rfc6902PatchOperation) {
	patch = append(patch, rfc6902PatchOperation{
		Op:    "add",
		Path:  basePath,
		Value: target,
	})
	return patch
}

func addVolume(target, added []corev1.Volume, basePath string) (patch []rfc6902PatchOperation) {
	first := len(target) == 0
	var value interface{}
	for _, add := range added {
		value = add
		path := basePath
		if first {
			first = false
			value = []corev1.Volume{add}
		} else {
			path += "/-"
		}
		patch = append(patch, rfc6902PatchOperation{
			Op:    "add",
			Path:  path,
			Value: value,
		})
	}
	return patch
}

func addImagePullSecrets(target, added []corev1.LocalObjectReference, basePath string) (patch []rfc6902PatchOperation) {
	first := len(target) == 0
	var value interface{}
	for _, add := range added {
		value = add
		path := basePath
		if first {
			first = false
			value = []corev1.LocalObjectReference{add}
		} else {
			path += "/-"
		}
		patch = append(patch, rfc6902PatchOperation{
			Op:    "add",
			Path:  path,
			Value: value,
		})
	}
	return patch
}

func addPodDNSConfig(target *corev1.PodDNSConfig, basePath string) (patch []rfc6902PatchOperation) {
	patch = append(patch, rfc6902PatchOperation{
		Op:    "add",
		Path:  basePath,
		Value: target,
	})
	return patch
}

// escape JSON Pointer value per https://tools.ietf.org/html/rfc6901
func escapeJSONPointerValue(in string) string {
	step := strings.Replace(in, "~", "~0", -1)
	return strings.Replace(step, "/", "~1", -1)
}

// adds labels to the target spec, will not overwrite label's value if it already exists
func addLabels(target map[string]string, added map[string]string) []rfc6902PatchOperation {
	patches := []rfc6902PatchOperation{}

	addedKeys := make([]string, 0, len(added))
	for key := range added {
		addedKeys = append(addedKeys, key)
	}
	sort.Strings(addedKeys)

	for _, key := range addedKeys {
		value := added[key]
		patch := rfc6902PatchOperation{
			Op:    "add",
			Path:  "/metadata/labels/" + escapeJSONPointerValue(key),
			Value: value,
		}

		if target == nil {
			target = map[string]string{}
			patch.Path = "/metadata/labels"
			patch.Value = map[string]string{
				key: value,
			}
		}

		if target[key] == "" {
			patches = append(patches, patch)
		}
	}

	return patches
}

func updateAnnotation(target map[string]string, added map[string]string) (patch []rfc6902PatchOperation) {
	// To ensure deterministic patches, we sort the keys
	var keys []string
	for k := range added {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := added[key]
		if target == nil {
			target = map[string]string{}
			patch = append(patch, rfc6902PatchOperation{
				Op:   "add",
				Path: "/metadata/annotations",
				Value: map[string]string{
					key: value,
				},
			})
		} else {
			op := "add"
			if target[key] != "" {
				op = "replace"
			}
			patch = append(patch, rfc6902PatchOperation{
				Op:    op,
				Path:  "/metadata/annotations/" + escapeJSONPointerValue(key),
				Value: value,
			})
		}
	}
	return patch
}
