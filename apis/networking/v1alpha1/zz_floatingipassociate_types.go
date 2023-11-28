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

type FloatingipAssociateInitParameters struct {

	// optional string →  One of the port's IP addresses.
	// One of the port's IP addresses.
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// optional string →  The region in which to obtain the Networking client. A Networking client is needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted, the region argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).
	// The region in which to obtain the Networking client. A Networking client is needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted, the `region` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is project's default SDN.
	// SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is project's default SDN.
	Sdn *string `json:"sdn,omitempty" tf:"sdn,omitempty"`
}

type FloatingipAssociateObservation struct {

	// optional string →  One of the port's IP addresses.
	// One of the port's IP addresses.
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// required string →  IP Address of an existing floating IP.
	// IP Address of an existing floating IP.
	FloatingIP *string `json:"floatingIp,omitempty" tf:"floating_ip,omitempty"`

	// string →  ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// required string →  ID of an existing port with at least one IP address to associate with this floating IP.
	// ID of an existing port with at least one IP address to associate with this floating IP.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// optional string →  The region in which to obtain the Networking client. A Networking client is needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted, the region argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).
	// The region in which to obtain the Networking client. A Networking client is needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted, the `region` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is project's default SDN.
	// SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is project's default SDN.
	Sdn *string `json:"sdn,omitempty" tf:"sdn,omitempty"`
}

type FloatingipAssociateParameters struct {

	// optional string →  One of the port's IP addresses.
	// One of the port's IP addresses.
	// +kubebuilder:validation:Optional
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// required string →  IP Address of an existing floating IP.
	// IP Address of an existing floating IP.
	// +crossplane:generate:reference:type=Floatingip
	// +crossplane:generate:reference:extractor=github.com/viletay/provider-vkcs/config/networking.ExtractSpecAddress()
	// +kubebuilder:validation:Optional
	FloatingIP *string `json:"floatingIp,omitempty" tf:"floating_ip,omitempty"`

	// Reference to a Floatingip to populate floatingIp.
	// +kubebuilder:validation:Optional
	FloatingIPRef *v1.Reference `json:"floatingIpRef,omitempty" tf:"-"`

	// Selector for a Floatingip to populate floatingIp.
	// +kubebuilder:validation:Optional
	FloatingIPSelector *v1.Selector `json:"floatingIpSelector,omitempty" tf:"-"`

	// required string →  ID of an existing port with at least one IP address to associate with this floating IP.
	// ID of an existing port with at least one IP address to associate with this floating IP.
	// +crossplane:generate:reference:type=Port
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Reference to a Port to populate portId.
	// +kubebuilder:validation:Optional
	PortIDRef *v1.Reference `json:"portIdRef,omitempty" tf:"-"`

	// Selector for a Port to populate portId.
	// +kubebuilder:validation:Optional
	PortIDSelector *v1.Selector `json:"portIdSelector,omitempty" tf:"-"`

	// optional string →  The region in which to obtain the Networking client. A Networking client is needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted, the region argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).
	// The region in which to obtain the Networking client. A Networking client is needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted, the `region` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is project's default SDN.
	// SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is project's default SDN.
	// +kubebuilder:validation:Optional
	Sdn *string `json:"sdn,omitempty" tf:"sdn,omitempty"`
}

// FloatingipAssociateSpec defines the desired state of FloatingipAssociate
type FloatingipAssociateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FloatingipAssociateParameters `json:"forProvider"`
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
	InitProvider FloatingipAssociateInitParameters `json:"initProvider,omitempty"`
}

// FloatingipAssociateStatus defines the observed state of FloatingipAssociate.
type FloatingipAssociateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FloatingipAssociateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FloatingipAssociate is the Schema for the FloatingipAssociates API. Associates a Floating IP to a Port
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vkcs}
type FloatingipAssociate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FloatingipAssociateSpec   `json:"spec"`
	Status            FloatingipAssociateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FloatingipAssociateList contains a list of FloatingipAssociates
type FloatingipAssociateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FloatingipAssociate `json:"items"`
}

// Repository type metadata.
var (
	FloatingipAssociate_Kind             = "FloatingipAssociate"
	FloatingipAssociate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FloatingipAssociate_Kind}.String()
	FloatingipAssociate_KindAPIVersion   = FloatingipAssociate_Kind + "." + CRDGroupVersion.String()
	FloatingipAssociate_GroupVersionKind = CRDGroupVersion.WithKind(FloatingipAssociate_Kind)
)

func init() {
	SchemeBuilder.Register(&FloatingipAssociate{}, &FloatingipAssociateList{})
}
