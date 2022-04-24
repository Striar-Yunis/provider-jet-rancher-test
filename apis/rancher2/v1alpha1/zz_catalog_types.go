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

type CatalogObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CatalogParameters struct {

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Optional
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	Refresh *bool `json:"refresh,omitempty" tf:"refresh,omitempty"`

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`

	// +kubebuilder:validation:Optional
	UsernameSecretRef *v1.SecretKeySelector `json:"usernameSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// CatalogSpec defines the desired state of Catalog
type CatalogSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatalogParameters `json:"forProvider"`
}

// CatalogStatus defines the observed state of Catalog.
type CatalogStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatalogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Catalog is the Schema for the Catalogs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ranchertestjet}
type Catalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CatalogSpec   `json:"spec"`
	Status            CatalogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogList contains a list of Catalogs
type CatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Catalog `json:"items"`
}

// Repository type metadata.
var (
	Catalog_Kind             = "Catalog"
	Catalog_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Catalog_Kind}.String()
	Catalog_KindAPIVersion   = Catalog_Kind + "." + CRDGroupVersion.String()
	Catalog_GroupVersionKind = CRDGroupVersion.WithKind(Catalog_Kind)
)

func init() {
	SchemeBuilder.Register(&Catalog{}, &CatalogList{})
}
