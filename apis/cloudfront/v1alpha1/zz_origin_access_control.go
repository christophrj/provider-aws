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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// OriginAccessControlParameters defines the desired state of OriginAccessControl
type OriginAccessControlParameters struct {
	// Region is which region the OriginAccessControl will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Contains the origin access control.
	// +kubebuilder:validation:Required
	OriginAccessControlConfig           *OriginAccessControlConfig `json:"originAccessControlConfig"`
	CustomOriginAccessControlParameters `json:",inline"`
}

// OriginAccessControlSpec defines the desired state of OriginAccessControl
type OriginAccessControlSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       OriginAccessControlParameters `json:"forProvider"`
}

// OriginAccessControlObservation defines the observed state of OriginAccessControl
type OriginAccessControlObservation struct {
	// The version identifier for the current version of the origin access control.
	ETag *string `json:"eTag,omitempty"`
	// The URL of the origin access control.
	Location *string `json:"location,omitempty"`
	// Contains an origin access control.
	OriginAccessControl *OriginAccessControl_SDK `json:"originAccessControl,omitempty"`
}

// OriginAccessControlStatus defines the observed state of OriginAccessControl.
type OriginAccessControlStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          OriginAccessControlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OriginAccessControl is the Schema for the OriginAccessControls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OriginAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OriginAccessControlSpec   `json:"spec"`
	Status            OriginAccessControlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OriginAccessControlList contains a list of OriginAccessControls
type OriginAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OriginAccessControl `json:"items"`
}

// Repository type metadata.
var (
	OriginAccessControlKind             = "OriginAccessControl"
	OriginAccessControlGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OriginAccessControlKind}.String()
	OriginAccessControlKindAPIVersion   = OriginAccessControlKind + "." + GroupVersion.String()
	OriginAccessControlGroupVersionKind = GroupVersion.WithKind(OriginAccessControlKind)
)

func init() {
	SchemeBuilder.Register(&OriginAccessControl{}, &OriginAccessControlList{})
}