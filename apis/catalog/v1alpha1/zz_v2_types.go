/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type V2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ResourceVersion *string `json:"resourceVersion,omitempty" tf:"resource_version,omitempty"`
}

type V2Parameters struct {

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// PEM encoded CA bundle which will be used to validate the repo's certificate
	// +kubebuilder:validation:Optional
	CABundle *string `json:"caBundle,omitempty" tf:"ca_bundle,omitempty"`

	// K8s cluster ID
	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// If disabled the repo clone will not be updated or allowed to be installed from
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Git Repository branch containing Helm chart definitions
	// +kubebuilder:validation:Optional
	GitBranch *string `json:"gitBranch,omitempty" tf:"git_branch,omitempty"`

	// Git Repository containing Helm chart definitions
	// +kubebuilder:validation:Optional
	GitRepo *string `json:"gitRepo,omitempty" tf:"git_repo,omitempty"`

	// Use insecure HTTPS to download the repo's index
	// +kubebuilder:validation:Optional
	Insecure *bool `json:"insecure,omitempty" tf:"insecure,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// K8s secret name to be used to connect to the repo
	// +kubebuilder:validation:Optional
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`

	// K8s secret namespace
	// +kubebuilder:validation:Optional
	SecretNamespace *string `json:"secretNamespace,omitempty" tf:"secret_namespace,omitempty"`

	// K8s service account used to deploy charts instead of the end users credentials
	// +kubebuilder:validation:Optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// K8s namsepace of the service account
	// +kubebuilder:validation:Optional
	ServiceAccountNamespace *string `json:"serviceAccountNamespace,omitempty" tf:"service_account_namespace,omitempty"`

	// URL to an index generated by Helm
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// V2Spec defines the desired state of V2
type V2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     V2Parameters `json:"forProvider"`
}

// V2Status defines the observed state of V2.
type V2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        V2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// V2 is the Schema for the V2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ranchertestjet}
type V2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              V2Spec   `json:"spec"`
	Status            V2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// V2List contains a list of V2s
type V2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []V2 `json:"items"`
}

// Repository type metadata.
var (
	V2_Kind             = "V2"
	V2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: V2_Kind}.String()
	V2_KindAPIVersion   = V2_Kind + "." + CRDGroupVersion.String()
	V2_GroupVersionKind = CRDGroupVersion.WithKind(V2_Kind)
)

func init() {
	SchemeBuilder.Register(&V2{}, &V2List{})
}
