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

type MemberInitParameters struct {

	// required string →  The IP address of the member to receive traffic from the load balancer. Changing this creates a new member.
	// The IP address of the member to receive traffic from the load balancer. Changing this creates a new member.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// optional boolean →  The administrative state of the member. A valid value is true (UP) or false (DOWN). Defaults to true.
	// The administrative state of the member. A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// readable name for the member.
	// Human-readable name for the member.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// required number →  The port on which to listen for client traffic. Changing this creates a new member.
	// The port on which to listen for client traffic. Changing this creates a new member.
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// optional string →  The region in which to obtain the Loadbalancer client. If omitted, the region argument of the provider is used. Changing this creates a new member.
	// The region in which to obtain the Loadbalancer client. If omitted, the `region` argument of the provider is used. Changing this creates a new member.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional number →  A positive integer value that indicates the relative portion of traffic that this member should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a member with a weight of 2. Defaults to 1.
	// A positive integer value that indicates the relative portion of traffic that this member should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a member with a weight of 2. Defaults to 1.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type MemberObservation struct {

	// required string →  The IP address of the member to receive traffic from the load balancer. Changing this creates a new member.
	// The IP address of the member to receive traffic from the load balancer. Changing this creates a new member.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// optional boolean →  The administrative state of the member. A valid value is true (UP) or false (DOWN). Defaults to true.
	// The administrative state of the member. A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// string →  ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// readable name for the member.
	// Human-readable name for the member.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// required string →  The id of the pool that this member will be assigned to. Changing this creates a new member.
	// The id of the pool that this member will be assigned to. Changing this creates a new member.
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// required number →  The port on which to listen for client traffic. Changing this creates a new member.
	// The port on which to listen for client traffic. Changing this creates a new member.
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// optional string →  The region in which to obtain the Loadbalancer client. If omitted, the region argument of the provider is used. Changing this creates a new member.
	// The region in which to obtain the Loadbalancer client. If omitted, the `region` argument of the provider is used. Changing this creates a new member.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  The subnet in which to access the member. Changing this creates a new member.
	// The subnet in which to access the member. Changing this creates a new member.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// optional number →  A positive integer value that indicates the relative portion of traffic that this member should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a member with a weight of 2. Defaults to 1.
	// A positive integer value that indicates the relative portion of traffic that this member should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a member with a weight of 2. Defaults to 1.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type MemberParameters struct {

	// required string →  The IP address of the member to receive traffic from the load balancer. Changing this creates a new member.
	// The IP address of the member to receive traffic from the load balancer. Changing this creates a new member.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// optional boolean →  The administrative state of the member. A valid value is true (UP) or false (DOWN). Defaults to true.
	// The administrative state of the member. A valid value is true (UP) or false (DOWN). Defaults to true.
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// readable name for the member.
	// Human-readable name for the member.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// required string →  The id of the pool that this member will be assigned to. Changing this creates a new member.
	// The id of the pool that this member will be assigned to. Changing this creates a new member.
	// +crossplane:generate:reference:type=Pool
	// +kubebuilder:validation:Optional
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// Reference to a Pool to populate poolId.
	// +kubebuilder:validation:Optional
	PoolIDRef *v1.Reference `json:"poolIdRef,omitempty" tf:"-"`

	// Selector for a Pool to populate poolId.
	// +kubebuilder:validation:Optional
	PoolIDSelector *v1.Selector `json:"poolIdSelector,omitempty" tf:"-"`

	// required number →  The port on which to listen for client traffic. Changing this creates a new member.
	// The port on which to listen for client traffic. Changing this creates a new member.
	// +kubebuilder:validation:Optional
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// optional string →  The region in which to obtain the Loadbalancer client. If omitted, the region argument of the provider is used. Changing this creates a new member.
	// The region in which to obtain the Loadbalancer client. If omitted, the `region` argument of the provider is used. Changing this creates a new member.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  The subnet in which to access the member. Changing this creates a new member.
	// The subnet in which to access the member. Changing this creates a new member.
	// +crossplane:generate:reference:type=github.com/viletay/provider-vkcs/apis/network/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// optional number →  A positive integer value that indicates the relative portion of traffic that this member should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a member with a weight of 2. Defaults to 1.
	// A positive integer value that indicates the relative portion of traffic that this member should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a member with a weight of 2. Defaults to 1.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

// MemberSpec defines the desired state of Member
type MemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberParameters `json:"forProvider"`
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
	InitProvider MemberInitParameters `json:"initProvider,omitempty"`
}

// MemberStatus defines the observed state of Member.
type MemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Member is the Schema for the Members API. Manages a member resource within VKCS.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vkcs}
type Member struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.address) || (has(self.initProvider) && has(self.initProvider.address))",message="spec.forProvider.address is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocolPort) || (has(self.initProvider) && has(self.initProvider.protocolPort))",message="spec.forProvider.protocolPort is a required parameter"
	Spec   MemberSpec   `json:"spec"`
	Status MemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemberList contains a list of Members
type MemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Member `json:"items"`
}

// Repository type metadata.
var (
	Member_Kind             = "Member"
	Member_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Member_Kind}.String()
	Member_KindAPIVersion   = Member_Kind + "." + CRDGroupVersion.String()
	Member_GroupVersionKind = CRDGroupVersion.WithKind(Member_Kind)
)

func init() {
	SchemeBuilder.Register(&Member{}, &MemberList{})
}