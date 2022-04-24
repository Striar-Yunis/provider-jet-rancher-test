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

type ConfigGithubObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConfigGithubParameters struct {

	// +kubebuilder:validation:Optional
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedPrincipalIds []*string `json:"allowedPrincipalIds,omitempty" tf:"allowed_principal_ids,omitempty"`

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	ClientIDSecretRef v1.SecretKeySelector `json:"clientIdSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	TLS *bool `json:"tls,omitempty" tf:"tls,omitempty"`
}

// ConfigGithubSpec defines the desired state of ConfigGithub
type ConfigGithubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigGithubParameters `json:"forProvider"`
}

// ConfigGithubStatus defines the observed state of ConfigGithub.
type ConfigGithubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigGithubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigGithub is the Schema for the ConfigGithubs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ranchertestjet}
type ConfigGithub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigGithubSpec   `json:"spec"`
	Status            ConfigGithubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigGithubList contains a list of ConfigGithubs
type ConfigGithubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigGithub `json:"items"`
}

// Repository type metadata.
var (
	ConfigGithub_Kind             = "ConfigGithub"
	ConfigGithub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigGithub_Kind}.String()
	ConfigGithub_KindAPIVersion   = ConfigGithub_Kind + "." + CRDGroupVersion.String()
	ConfigGithub_GroupVersionKind = CRDGroupVersion.WithKind(ConfigGithub_Kind)
)

func init() {
	SchemeBuilder.Register(&ConfigGithub{}, &ConfigGithubList{})
}