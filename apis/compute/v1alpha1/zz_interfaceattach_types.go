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

type InterfaceAttachInitParameters struct {

	// An IP address to assosciate with the port. _note_ This option cannot be used with port_id. You must specify a network_id. The IP address must lie in a range on the supplied network.
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// The region in which to create the interface attachment. If omitted, the `region` argument of the provider is used. Changing this creates a new attachment.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type InterfaceAttachObservation struct {

	// An IP address to assosciate with the port. _note_ This option cannot be used with port_id. You must specify a network_id. The IP address must lie in a range on the supplied network.
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Instance to attach the Port or Network to.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The ID of the Network to attach to an Instance. A port will be created automatically. _note_ This option and `port_id` are mutually exclusive.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// The ID of the Port to attach to an Instance. _note_ This option and `network_id` are mutually exclusive.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The region in which to create the interface attachment. If omitted, the `region` argument of the provider is used. Changing this creates a new attachment.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type InterfaceAttachParameters struct {

	// An IP address to assosciate with the port. _note_ This option cannot be used with port_id. You must specify a network_id. The IP address must lie in a range on the supplied network.
	// +kubebuilder:validation:Optional
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// The ID of the Instance to attach the Port or Network to.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// The ID of the Network to attach to an Instance. A port will be created automatically. _note_ This option and `port_id` are mutually exclusive.
	// +crossplane:generate:reference:type=github.com/viletay/provider-vkcs/apis/network/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// The ID of the Port to attach to an Instance. _note_ This option and `network_id` are mutually exclusive.
	// +crossplane:generate:reference:type=github.com/viletay/provider-vkcs/apis/network/v1alpha1.Port
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Reference to a Port in network to populate portId.
	// +kubebuilder:validation:Optional
	PortIDRef *v1.Reference `json:"portIdRef,omitempty" tf:"-"`

	// Selector for a Port in network to populate portId.
	// +kubebuilder:validation:Optional
	PortIDSelector *v1.Selector `json:"portIdSelector,omitempty" tf:"-"`

	// The region in which to create the interface attachment. If omitted, the `region` argument of the provider is used. Changing this creates a new attachment.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// InterfaceAttachSpec defines the desired state of InterfaceAttach
type InterfaceAttachSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InterfaceAttachParameters `json:"forProvider"`
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
	InitProvider InterfaceAttachInitParameters `json:"initProvider,omitempty"`
}

// InterfaceAttachStatus defines the observed state of InterfaceAttach.
type InterfaceAttachStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InterfaceAttachObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InterfaceAttach is the Schema for the InterfaceAttachs API. Attaches a Network Interface to an Instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vkcs}
type InterfaceAttach struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InterfaceAttachSpec   `json:"spec"`
	Status            InterfaceAttachStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InterfaceAttachList contains a list of InterfaceAttachs
type InterfaceAttachList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InterfaceAttach `json:"items"`
}

// Repository type metadata.
var (
	InterfaceAttach_Kind             = "InterfaceAttach"
	InterfaceAttach_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InterfaceAttach_Kind}.String()
	InterfaceAttach_KindAPIVersion   = InterfaceAttach_Kind + "." + CRDGroupVersion.String()
	InterfaceAttach_GroupVersionKind = CRDGroupVersion.WithKind(InterfaceAttach_Kind)
)

func init() {
	SchemeBuilder.Register(&InterfaceAttach{}, &InterfaceAttachList{})
}
