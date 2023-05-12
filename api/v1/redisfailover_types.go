/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RedisCommandRename defines the specification of a "rename-command" configuration option
type RedisCommandRename struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// RedisSettings defines the specification of the redis cluster
type RedisSettings struct {
	Image                         string                            `json:"image,omitempty"`
	ImagePullPolicy               corev1.PullPolicy                 `json:"imagePullPolicy,omitempty"`
	Replicas                      int32                             `json:"replicas,omitempty"`
	Port                          int32                             `json:"port,omitempty"`
	Resources                     corev1.ResourceRequirements       `json:"resources,omitempty"`
	CustomConfig                  []string                          `json:"customConfig,omitempty"`
	CustomCommandRenames          []RedisCommandRename              `json:"customCommandRenames,omitempty"`
	Command                       []string                          `json:"command,omitempty"`
	ShutdownConfigMap             string                            `json:"shutdownConfigMap,omitempty"`
	StartupConfigMap              string                            `json:"startupConfigMap,omitempty"`
	Storage                       RedisStorage                      `json:"storage,omitempty"`
	InitContainers                []corev1.Container                `json:"initContainers,omitempty"`
	Exporter                      Exporter                          `json:"exporter,omitempty"`
	ExtraContainers               []corev1.Container                `json:"extraContainers,omitempty"`
	Affinity                      *corev1.Affinity                  `json:"affinity,omitempty"`
	SecurityContext               *corev1.PodSecurityContext        `json:"securityContext,omitempty"`
	ContainerSecurityContext      *corev1.SecurityContext           `json:"containerSecurityContext,omitempty"`
	ImagePullSecrets              []corev1.LocalObjectReference     `json:"imagePullSecrets,omitempty"`
	Tolerations                   []corev1.Toleration               `json:"tolerations,omitempty"`
	TopologySpreadConstraints     []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
	NodeSelector                  map[string]string                 `json:"nodeSelector,omitempty"`
	PodAnnotations                map[string]string                 `json:"podAnnotations,omitempty"`
	ServiceAnnotations            map[string]string                 `json:"serviceAnnotations,omitempty"`
	HostNetwork                   bool                              `json:"hostNetwork,omitempty"`
	DNSPolicy                     corev1.DNSPolicy                  `json:"dnsPolicy,omitempty"`
	PriorityClassName             string                            `json:"priorityClassName,omitempty"`
	ServiceAccountName            string                            `json:"serviceAccountName,omitempty"`
	TerminationGracePeriodSeconds int64                             `json:"terminationGracePeriod,omitempty"`
	ExtraVolumes                  []corev1.Volume                   `json:"extraVolumes,omitempty"`
	ExtraVolumeMounts             []corev1.VolumeMount              `json:"extraVolumeMounts,omitempty"`
	CustomLivenessProbe           *corev1.Probe                     `json:"customLivenessProbe,omitempty"`
	CustomReadinessProbe          *corev1.Probe                     `json:"customReadinessProbe,omitempty"`
	CustomStartupProbe            *corev1.Probe                     `json:"customStartupProbe,omitempty"`
}

// SentinelSettings defines the specification of the sentinel cluster
type SentinelSettings struct {
	Image                     string                            `json:"image,omitempty"`
	ImagePullPolicy           corev1.PullPolicy                 `json:"imagePullPolicy,omitempty"`
	Replicas                  int32                             `json:"replicas,omitempty"`
	Resources                 corev1.ResourceRequirements       `json:"resources,omitempty"`
	CustomConfig              []string                          `json:"customConfig,omitempty"`
	Command                   []string                          `json:"command,omitempty"`
	StartupConfigMap          string                            `json:"startupConfigMap,omitempty"`
	Affinity                  *corev1.Affinity                  `json:"affinity,omitempty"`
	SecurityContext           *corev1.PodSecurityContext        `json:"securityContext,omitempty"`
	ContainerSecurityContext  *corev1.SecurityContext           `json:"containerSecurityContext,omitempty"`
	ImagePullSecrets          []corev1.LocalObjectReference     `json:"imagePullSecrets,omitempty"`
	Tolerations               []corev1.Toleration               `json:"tolerations,omitempty"`
	TopologySpreadConstraints []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
	NodeSelector              map[string]string                 `json:"nodeSelector,omitempty"`
	PodAnnotations            map[string]string                 `json:"podAnnotations,omitempty"`
	ServiceAnnotations        map[string]string                 `json:"serviceAnnotations,omitempty"`
	InitContainers            []corev1.Container                `json:"initContainers,omitempty"`
	Exporter                  Exporter                          `json:"exporter,omitempty"`
	ExtraContainers           []corev1.Container                `json:"extraContainers,omitempty"`
	ConfigCopy                SentinelConfigCopy                `json:"configCopy,omitempty"`
	HostNetwork               bool                              `json:"hostNetwork,omitempty"`
	DNSPolicy                 corev1.DNSPolicy                  `json:"dnsPolicy,omitempty"`
	PriorityClassName         string                            `json:"priorityClassName,omitempty"`
	ServiceAccountName        string                            `json:"serviceAccountName,omitempty"`
	ExtraVolumes              []corev1.Volume                   `json:"extraVolumes,omitempty"`
	ExtraVolumeMounts         []corev1.VolumeMount              `json:"extraVolumeMounts,omitempty"`
	CustomLivenessProbe       *corev1.Probe                     `json:"customLivenessProbe,omitempty"`
	CustomReadinessProbe      *corev1.Probe                     `json:"customReadinessProbe,omitempty"`
	CustomStartupProbe        *corev1.Probe                     `json:"customStartupProbe,omitempty"`
}

// AuthSettings contains settings about auth
type AuthSettings struct {
	SecretPath string `json:"secretPath,omitempty"`
}

// BootstrapSettings contains settings about a potential bootstrap node
type BootstrapSettings struct {
	Host           string `json:"host,omitempty"`
	Port           string `json:"port,omitempty"`
	AllowSentinels bool   `json:"allowSentinels,omitempty"`
}

// Exporter defines the specification for the redis/sentinel exporter
type Exporter struct {
	Enabled                  bool                         `json:"enabled,omitempty"`
	Image                    string                       `json:"image,omitempty"`
	ImagePullPolicy          corev1.PullPolicy            `json:"imagePullPolicy,omitempty"`
	ContainerSecurityContext *corev1.SecurityContext      `json:"containerSecurityContext,omitempty"`
	Args                     []string                     `json:"args,omitempty"`
	Env                      []corev1.EnvVar              `json:"env,omitempty"`
	Resources                *corev1.ResourceRequirements `json:"resources,omitempty"`
}

// SentinelConfigCopy defines the specification for the sentinel exporter
type SentinelConfigCopy struct {
	ContainerSecurityContext *corev1.SecurityContext `json:"containerSecurityContext,omitempty"`
}

// RedisStorage defines the structure used to store the Redis Data
type RedisStorage struct {
	KeepAfterDeletion     bool                           `json:"keepAfterDeletion,omitempty"`
	EmptyDir              *corev1.EmptyDirVolumeSource   `json:"emptyDir,omitempty"`
	PersistentVolumeClaim *EmbeddedPersistentVolumeClaim `json:"persistentVolumeClaim,omitempty"`
}

// EmbeddedPersistentVolumeClaim is an embedded version of k8s.io/api/core/v1.PersistentVolumeClaim.
// It contains TypeMeta and a reduced ObjectMeta.
type EmbeddedPersistentVolumeClaim struct {
	metav1.TypeMeta `json:",inline"`

	// EmbeddedMetadata contains metadata relevant to an EmbeddedResource.
	EmbeddedObjectMetadata `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the desired characteristics of a volume requested by a pod author.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// +optional
	Spec corev1.PersistentVolumeClaimSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the current information/status of a persistent volume claim.
	// Read-only.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// +optional
	Status corev1.PersistentVolumeClaimStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// EmbeddedObjectMetadata contains a subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta
// Only fields which are relevant to embedded resources are included.
type EmbeddedObjectMetadata struct {
	// Name must be unique within a namespace. Is required when creating resources, although
	// some resources may allow a client to request the generation of an appropriate name
	// automatically. Name is primarily intended for creation idempotence and configuration
	// definition.
	// Cannot be updated.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services.
	// More info: http://kubernetes.io/docs/user-guide/labels
	// +optional
	Labels map[string]string `json:"labels,omitempty" protobuf:"bytes,11,rep,name=labels"`

	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	// More info: http://kubernetes.io/docs/user-guide/annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,12,rep,name=annotations"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RedisFailoverSpec defines the desired state of RedisFailover
type RedisFailoverSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RedisFailover. Edit redisfailover_types.go to remove/update
	Redis          RedisSettings      `json:"redis,omitempty"`
	Sentinel       SentinelSettings   `json:"sentinel,omitempty"`
	Auth           AuthSettings       `json:"auth,omitempty"`
	LabelWhitelist []string           `json:"labelWhitelist,omitempty"`
	BootstrapNode  *BootstrapSettings `json:"bootstrapNode,omitempty"`
}

// RedisFailoverStatus defines the observed state of RedisFailover
type RedisFailoverStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RedisFailover is the Schema for the redisfailovers API
type RedisFailover struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisFailoverSpec   `json:"spec,omitempty"`
	Status RedisFailoverStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RedisFailoverList contains a list of RedisFailover
type RedisFailoverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisFailover `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedisFailover{}, &RedisFailoverList{})
}
