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

type AlertRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AlertRuleParameters struct {

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Alert rule cluster ID
	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// Alert event rule
	// +kubebuilder:validation:Optional
	EventRule []EventRuleParameters `json:"eventRule,omitempty" tf:"event_rule,omitempty"`

	// Alert rule group ID
	// +kubebuilder:validation:Required
	GroupID *string `json:"groupId" tf:"group_id,omitempty"`

	// Alert rule interval seconds
	// +kubebuilder:validation:Optional
	GroupIntervalSeconds *float64 `json:"groupIntervalSeconds,omitempty" tf:"group_interval_seconds,omitempty"`

	// Alert rule wait seconds
	// +kubebuilder:validation:Optional
	GroupWaitSeconds *float64 `json:"groupWaitSeconds,omitempty" tf:"group_wait_seconds,omitempty"`

	// Alert rule inherited
	// +kubebuilder:validation:Optional
	Inherited *bool `json:"inherited,omitempty" tf:"inherited,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Alert metric rule
	// +kubebuilder:validation:Optional
	MetricRule []MetricRuleParameters `json:"metricRule,omitempty" tf:"metric_rule,omitempty"`

	// Alert node rule
	// +kubebuilder:validation:Optional
	NodeRule []NodeRuleParameters `json:"nodeRule,omitempty" tf:"node_rule,omitempty"`

	// Alert rule repeat interval seconds
	// +kubebuilder:validation:Optional
	RepeatIntervalSeconds *float64 `json:"repeatIntervalSeconds,omitempty" tf:"repeat_interval_seconds,omitempty"`

	// Alert rule severity
	// +kubebuilder:validation:Optional
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// Alert system service rule
	// +kubebuilder:validation:Optional
	SystemServiceRule []SystemServiceRuleParameters `json:"systemServiceRule,omitempty" tf:"system_service_rule,omitempty"`
}

type EventRuleObservation struct {
}

type EventRuleParameters struct {

	// Event type
	// +kubebuilder:validation:Optional
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// Resource kind
	// +kubebuilder:validation:Required
	ResourceKind *string `json:"resourceKind" tf:"resource_kind,omitempty"`
}

type MetricRuleObservation struct {
}

type MetricRuleParameters struct {

	// Metric rule comparison
	// +kubebuilder:validation:Optional
	Comparison *string `json:"comparison,omitempty" tf:"comparison,omitempty"`

	// Metric rule description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Metric rule duration
	// +kubebuilder:validation:Required
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// Metric rule expression
	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// Metric rule threshold value
	// +kubebuilder:validation:Required
	ThresholdValue *float64 `json:"thresholdValue" tf:"threshold_value,omitempty"`
}

type NodeRuleObservation struct {
}

type NodeRuleParameters struct {

	// Node rule cpu threshold
	// +kubebuilder:validation:Optional
	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	// Node rule condition
	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Node rule mem threshold
	// +kubebuilder:validation:Optional
	MemThreshold *float64 `json:"memThreshold,omitempty" tf:"mem_threshold,omitempty"`

	// Node ID
	// +kubebuilder:validation:Optional
	NodeID *string `json:"nodeId,omitempty" tf:"node_id,omitempty"`

	// Node rule selector
	// +kubebuilder:validation:Optional
	Selector map[string]*string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type SystemServiceRuleObservation struct {
}

type SystemServiceRuleParameters struct {

	// System service rule condition
	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`
}

// AlertRuleSpec defines the desired state of AlertRule
type AlertRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlertRuleParameters `json:"forProvider"`
}

// AlertRuleStatus defines the observed state of AlertRule.
type AlertRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlertRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlertRule is the Schema for the AlertRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ranchertestjet}
type AlertRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlertRuleSpec   `json:"spec"`
	Status            AlertRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertRuleList contains a list of AlertRules
type AlertRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertRule `json:"items"`
}

// Repository type metadata.
var (
	AlertRule_Kind             = "AlertRule"
	AlertRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlertRule_Kind}.String()
	AlertRule_KindAPIVersion   = AlertRule_Kind + "." + CRDGroupVersion.String()
	AlertRule_GroupVersionKind = CRDGroupVersion.WithKind(AlertRule_Kind)
)

func init() {
	SchemeBuilder.Register(&AlertRule{}, &AlertRuleList{})
}