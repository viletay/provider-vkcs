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

type ListenerInitParameters struct {

	// optional boolean →  The administrative state of the Listener. A valid value is true (UP) or false (DOWN).
	// The administrative state of the Listener. A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// optional string →  A list of CIDR blocks that are permitted to connect to this listener, denying all other source addresses. If not present, defaults to allow all.
	// A list of CIDR blocks that are permitted to connect to this listener, denying all other source addresses. If not present, defaults to allow all.
	AllowedCidrs []*string `json:"allowedCidrs,omitempty" tf:"allowed_cidrs,omitempty"`

	// optional number →  The maximum number of connections allowed for the Listener.
	// The maximum number of connections allowed for the Listener.
	ConnectionLimit *float64 `json:"connectionLimit,omitempty" tf:"connection_limit,omitempty"`

	// optional string →  The ID of the default pool with which the Listener is associated.
	// The ID of the default pool with which the Listener is associated.
	DefaultPoolID *string `json:"defaultPoolId,omitempty" tf:"default_pool_id,omitempty"`

	// optional string →  A reference to a Keymanager Secrets container which stores TLS information. This is required if the protocol is TERMINATED_HTTPS.
	// A reference to a Keymanager Secrets container which stores TLS information. This is required if the protocol is `TERMINATED_HTTPS`.
	DefaultTLSContainerRef *string `json:"defaultTlsContainerRef,omitempty" tf:"default_tls_container_ref,omitempty"`

	// readable description for the Listener.
	// Human-readable description for the Listener.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// optional map of string →  The list of key value pairs representing headers to insert into the request before it is sent to the backend members. Changing this updates the headers of the existing listener.
	// The list of key value pairs representing headers to insert into the request before it is sent to the backend members. Changing this updates the headers of the existing listener.
	InsertHeaders map[string]*string `json:"insertHeaders,omitempty" tf:"insert_headers,omitempty"`

	// readable name for the Listener. Does not have to be unique.
	// Human-readable name for the Listener. Does not have to be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// can either be TCP, HTTP, HTTPS, TERMINATED_HTTPS, UDP. Changing this creates a new Listener.
	// The protocol - can either be TCP, HTTP, HTTPS, TERMINATED_HTTPS, UDP. Changing this creates a new Listener.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// required number →  The port on which to listen for client traffic. Changing this creates a new Listener.
	// The port on which to listen for client traffic. Changing this creates a new Listener.
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// optional string →  The region in which to obtain the Loadbalancer client. If omitted, the region argument of the provider is used. Changing this creates a new Listener.
	// The region in which to obtain the Loadbalancer client. If omitted, the `region` argument of the provider is used. Changing this creates a new Listener.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  A list of references to Keymanager Secrets containers which store SNI information.
	// A list of references to Keymanager Secrets containers which store SNI information.
	SniContainerRefs []*string `json:"sniContainerRefs,omitempty" tf:"sni_container_refs,omitempty"`

	// optional number →  The client inactivity timeout in milliseconds.
	// The client inactivity timeout in milliseconds.
	TimeoutClientData *float64 `json:"timeoutClientData,omitempty" tf:"timeout_client_data,omitempty"`

	// optional number →  The member connection timeout in milliseconds.
	// The member connection timeout in milliseconds.
	TimeoutMemberConnect *float64 `json:"timeoutMemberConnect,omitempty" tf:"timeout_member_connect,omitempty"`

	// optional number →  The member inactivity timeout in milliseconds.
	// The member inactivity timeout in milliseconds.
	TimeoutMemberData *float64 `json:"timeoutMemberData,omitempty" tf:"timeout_member_data,omitempty"`

	// optional number →  The time in milliseconds, to wait for additional TCP packets for content inspection.
	// The time in milliseconds, to wait for additional TCP packets for content inspection.
	TimeoutTCPInspect *float64 `json:"timeoutTcpInspect,omitempty" tf:"timeout_tcp_inspect,omitempty"`
}

type ListenerObservation struct {

	// optional boolean →  The administrative state of the Listener. A valid value is true (UP) or false (DOWN).
	// The administrative state of the Listener. A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// optional string →  A list of CIDR blocks that are permitted to connect to this listener, denying all other source addresses. If not present, defaults to allow all.
	// A list of CIDR blocks that are permitted to connect to this listener, denying all other source addresses. If not present, defaults to allow all.
	AllowedCidrs []*string `json:"allowedCidrs,omitempty" tf:"allowed_cidrs,omitempty"`

	// optional number →  The maximum number of connections allowed for the Listener.
	// The maximum number of connections allowed for the Listener.
	ConnectionLimit *float64 `json:"connectionLimit,omitempty" tf:"connection_limit,omitempty"`

	// optional string →  The ID of the default pool with which the Listener is associated.
	// The ID of the default pool with which the Listener is associated.
	DefaultPoolID *string `json:"defaultPoolId,omitempty" tf:"default_pool_id,omitempty"`

	// optional string →  A reference to a Keymanager Secrets container which stores TLS information. This is required if the protocol is TERMINATED_HTTPS.
	// A reference to a Keymanager Secrets container which stores TLS information. This is required if the protocol is `TERMINATED_HTTPS`.
	DefaultTLSContainerRef *string `json:"defaultTlsContainerRef,omitempty" tf:"default_tls_container_ref,omitempty"`

	// readable description for the Listener.
	// Human-readable description for the Listener.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// string →  ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// optional map of string →  The list of key value pairs representing headers to insert into the request before it is sent to the backend members. Changing this updates the headers of the existing listener.
	// The list of key value pairs representing headers to insert into the request before it is sent to the backend members. Changing this updates the headers of the existing listener.
	InsertHeaders map[string]*string `json:"insertHeaders,omitempty" tf:"insert_headers,omitempty"`

	// required string →  The load balancer on which to provision this Listener. Changing this creates a new Listener.
	// The load balancer on which to provision this Listener. Changing this creates a new Listener.
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// readable name for the Listener. Does not have to be unique.
	// Human-readable name for the Listener. Does not have to be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// can either be TCP, HTTP, HTTPS, TERMINATED_HTTPS, UDP. Changing this creates a new Listener.
	// The protocol - can either be TCP, HTTP, HTTPS, TERMINATED_HTTPS, UDP. Changing this creates a new Listener.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// required number →  The port on which to listen for client traffic. Changing this creates a new Listener.
	// The port on which to listen for client traffic. Changing this creates a new Listener.
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// optional string →  The region in which to obtain the Loadbalancer client. If omitted, the region argument of the provider is used. Changing this creates a new Listener.
	// The region in which to obtain the Loadbalancer client. If omitted, the `region` argument of the provider is used. Changing this creates a new Listener.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  A list of references to Keymanager Secrets containers which store SNI information.
	// A list of references to Keymanager Secrets containers which store SNI information.
	SniContainerRefs []*string `json:"sniContainerRefs,omitempty" tf:"sni_container_refs,omitempty"`

	// optional number →  The client inactivity timeout in milliseconds.
	// The client inactivity timeout in milliseconds.
	TimeoutClientData *float64 `json:"timeoutClientData,omitempty" tf:"timeout_client_data,omitempty"`

	// optional number →  The member connection timeout in milliseconds.
	// The member connection timeout in milliseconds.
	TimeoutMemberConnect *float64 `json:"timeoutMemberConnect,omitempty" tf:"timeout_member_connect,omitempty"`

	// optional number →  The member inactivity timeout in milliseconds.
	// The member inactivity timeout in milliseconds.
	TimeoutMemberData *float64 `json:"timeoutMemberData,omitempty" tf:"timeout_member_data,omitempty"`

	// optional number →  The time in milliseconds, to wait for additional TCP packets for content inspection.
	// The time in milliseconds, to wait for additional TCP packets for content inspection.
	TimeoutTCPInspect *float64 `json:"timeoutTcpInspect,omitempty" tf:"timeout_tcp_inspect,omitempty"`
}

type ListenerParameters struct {

	// optional boolean →  The administrative state of the Listener. A valid value is true (UP) or false (DOWN).
	// The administrative state of the Listener. A valid value is true (UP) or false (DOWN).
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// optional string →  A list of CIDR blocks that are permitted to connect to this listener, denying all other source addresses. If not present, defaults to allow all.
	// A list of CIDR blocks that are permitted to connect to this listener, denying all other source addresses. If not present, defaults to allow all.
	// +kubebuilder:validation:Optional
	AllowedCidrs []*string `json:"allowedCidrs,omitempty" tf:"allowed_cidrs,omitempty"`

	// optional number →  The maximum number of connections allowed for the Listener.
	// The maximum number of connections allowed for the Listener.
	// +kubebuilder:validation:Optional
	ConnectionLimit *float64 `json:"connectionLimit,omitempty" tf:"connection_limit,omitempty"`

	// optional string →  The ID of the default pool with which the Listener is associated.
	// The ID of the default pool with which the Listener is associated.
	// +kubebuilder:validation:Optional
	DefaultPoolID *string `json:"defaultPoolId,omitempty" tf:"default_pool_id,omitempty"`

	// optional string →  A reference to a Keymanager Secrets container which stores TLS information. This is required if the protocol is TERMINATED_HTTPS.
	// A reference to a Keymanager Secrets container which stores TLS information. This is required if the protocol is `TERMINATED_HTTPS`.
	// +kubebuilder:validation:Optional
	DefaultTLSContainerRef *string `json:"defaultTlsContainerRef,omitempty" tf:"default_tls_container_ref,omitempty"`

	// readable description for the Listener.
	// Human-readable description for the Listener.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// optional map of string →  The list of key value pairs representing headers to insert into the request before it is sent to the backend members. Changing this updates the headers of the existing listener.
	// The list of key value pairs representing headers to insert into the request before it is sent to the backend members. Changing this updates the headers of the existing listener.
	// +kubebuilder:validation:Optional
	InsertHeaders map[string]*string `json:"insertHeaders,omitempty" tf:"insert_headers,omitempty"`

	// required string →  The load balancer on which to provision this Listener. Changing this creates a new Listener.
	// The load balancer on which to provision this Listener. Changing this creates a new Listener.
	// +crossplane:generate:reference:type=LoadBalancer
	// +kubebuilder:validation:Optional
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Reference to a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDRef *v1.Reference `json:"loadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDSelector *v1.Selector `json:"loadbalancerIdSelector,omitempty" tf:"-"`

	// readable name for the Listener. Does not have to be unique.
	// Human-readable name for the Listener. Does not have to be unique.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// can either be TCP, HTTP, HTTPS, TERMINATED_HTTPS, UDP. Changing this creates a new Listener.
	// The protocol - can either be TCP, HTTP, HTTPS, TERMINATED_HTTPS, UDP. Changing this creates a new Listener.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// required number →  The port on which to listen for client traffic. Changing this creates a new Listener.
	// The port on which to listen for client traffic. Changing this creates a new Listener.
	// +kubebuilder:validation:Optional
	ProtocolPort *float64 `json:"protocolPort,omitempty" tf:"protocol_port,omitempty"`

	// optional string →  The region in which to obtain the Loadbalancer client. If omitted, the region argument of the provider is used. Changing this creates a new Listener.
	// The region in which to obtain the Loadbalancer client. If omitted, the `region` argument of the provider is used. Changing this creates a new Listener.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  A list of references to Keymanager Secrets containers which store SNI information.
	// A list of references to Keymanager Secrets containers which store SNI information.
	// +kubebuilder:validation:Optional
	SniContainerRefs []*string `json:"sniContainerRefs,omitempty" tf:"sni_container_refs,omitempty"`

	// optional number →  The client inactivity timeout in milliseconds.
	// The client inactivity timeout in milliseconds.
	// +kubebuilder:validation:Optional
	TimeoutClientData *float64 `json:"timeoutClientData,omitempty" tf:"timeout_client_data,omitempty"`

	// optional number →  The member connection timeout in milliseconds.
	// The member connection timeout in milliseconds.
	// +kubebuilder:validation:Optional
	TimeoutMemberConnect *float64 `json:"timeoutMemberConnect,omitempty" tf:"timeout_member_connect,omitempty"`

	// optional number →  The member inactivity timeout in milliseconds.
	// The member inactivity timeout in milliseconds.
	// +kubebuilder:validation:Optional
	TimeoutMemberData *float64 `json:"timeoutMemberData,omitempty" tf:"timeout_member_data,omitempty"`

	// optional number →  The time in milliseconds, to wait for additional TCP packets for content inspection.
	// The time in milliseconds, to wait for additional TCP packets for content inspection.
	// +kubebuilder:validation:Optional
	TimeoutTCPInspect *float64 `json:"timeoutTcpInspect,omitempty" tf:"timeout_tcp_inspect,omitempty"`
}

// ListenerSpec defines the desired state of Listener
type ListenerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ListenerParameters `json:"forProvider"`
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
	InitProvider ListenerInitParameters `json:"initProvider,omitempty"`
}

// ListenerStatus defines the observed state of Listener.
type ListenerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Listener is the Schema for the Listeners API. Manages a listener resource within VKCS.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vkcs}
type Listener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocolPort) || (has(self.initProvider) && has(self.initProvider.protocolPort))",message="spec.forProvider.protocolPort is a required parameter"
	Spec   ListenerSpec   `json:"spec"`
	Status ListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ListenerList contains a list of Listeners
type ListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Listener `json:"items"`
}

// Repository type metadata.
var (
	Listener_Kind             = "Listener"
	Listener_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Listener_Kind}.String()
	Listener_KindAPIVersion   = Listener_Kind + "." + CRDGroupVersion.String()
	Listener_GroupVersionKind = CRDGroupVersion.WithKind(Listener_Kind)
)

func init() {
	SchemeBuilder.Register(&Listener{}, &ListenerList{})
}
