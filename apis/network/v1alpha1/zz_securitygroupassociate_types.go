// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecurityGroupAssociateInitParameters struct {

	// optional boolean →  Whether to replace or append the list of security groups, specified in the security_group_ids. Defaults to false.
	// Whether to replace or append the list of security groups, specified in the `security_group_ids`. Defaults to `false`.
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// optional string →  The region in which to obtain the networking client. A networking client is needed to manage a port. If omitted, the region argument of the provider is used. Changing this creates a new resource.
	// The region in which to obtain the networking client. A networking client is needed to manage a port. If omitted, the `region` argument of the provider is used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is project's default SDN.
	// SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is project's default SDN.
	Sdn *string `json:"sdn,omitempty" tf:"sdn,omitempty"`
}

type SecurityGroupAssociateObservation struct {

	// set of string →  The collection of Security Group IDs on the port which have been explicitly and implicitly added.
	// The collection of Security Group IDs on the port which have been explicitly and implicitly added.
	AllSecurityGroupIds []*string `json:"allSecurityGroupIds,omitempty" tf:"all_security_group_ids,omitempty"`

	// optional boolean →  Whether to replace or append the list of security groups, specified in the security_group_ids. Defaults to false.
	// Whether to replace or append the list of security groups, specified in the `security_group_ids`. Defaults to `false`.
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// string →  ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// required string →  An UUID of the port to apply security groups to.
	// An UUID of the port to apply security groups to.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// optional string →  The region in which to obtain the networking client. A networking client is needed to manage a port. If omitted, the region argument of the provider is used. Changing this creates a new resource.
	// The region in which to obtain the networking client. A networking client is needed to manage a port. If omitted, the `region` argument of the provider is used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is project's default SDN.
	// SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is project's default SDN.
	Sdn *string `json:"sdn,omitempty" tf:"sdn,omitempty"`

	// required set of string →  A list of security group IDs to apply to the port. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance).
	// A list of security group IDs to apply to the port. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance).
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`
}

type SecurityGroupAssociateParameters struct {

	// optional boolean →  Whether to replace or append the list of security groups, specified in the security_group_ids. Defaults to false.
	// Whether to replace or append the list of security groups, specified in the `security_group_ids`. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// required string →  An UUID of the port to apply security groups to.
	// An UUID of the port to apply security groups to.
	// +crossplane:generate:reference:type=Port
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Reference to a Port to populate portId.
	// +kubebuilder:validation:Optional
	PortIDRef *v1.Reference `json:"portIdRef,omitempty" tf:"-"`

	// Selector for a Port to populate portId.
	// +kubebuilder:validation:Optional
	PortIDSelector *v1.Selector `json:"portIdSelector,omitempty" tf:"-"`

	// optional string →  The region in which to obtain the networking client. A networking client is needed to manage a port. If omitted, the region argument of the provider is used. Changing this creates a new resource.
	// The region in which to obtain the networking client. A networking client is needed to manage a port. If omitted, the `region` argument of the provider is used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is project's default SDN.
	// SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is project's default SDN.
	// +kubebuilder:validation:Optional
	Sdn *string `json:"sdn,omitempty" tf:"sdn,omitempty"`

	// required set of string →  A list of security group IDs to apply to the port. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance).
	// A list of security group IDs to apply to the port. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance).
	// +crossplane:generate:reference:type=SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`
}

// SecurityGroupAssociateSpec defines the desired state of SecurityGroupAssociate
type SecurityGroupAssociateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupAssociateParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SecurityGroupAssociateInitParameters `json:"initProvider,omitempty"`
}

// SecurityGroupAssociateStatus defines the observed state of SecurityGroupAssociate.
type SecurityGroupAssociateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupAssociateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupAssociate is the Schema for the SecurityGroupAssociates API. Manages a port's security groups within VKCS.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vkcs}
type SecurityGroupAssociate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupAssociateSpec   `json:"spec"`
	Status            SecurityGroupAssociateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupAssociateList contains a list of SecurityGroupAssociates
type SecurityGroupAssociateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroupAssociate `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroupAssociate_Kind             = "SecurityGroupAssociate"
	SecurityGroupAssociate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroupAssociate_Kind}.String()
	SecurityGroupAssociate_KindAPIVersion   = SecurityGroupAssociate_Kind + "." + CRDGroupVersion.String()
	SecurityGroupAssociate_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroupAssociate_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroupAssociate{}, &SecurityGroupAssociateList{})
}
