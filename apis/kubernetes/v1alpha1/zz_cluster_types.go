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

type ClusterInitParameters struct {

	// optional string →  API LoadBalancer fip. IP address field.
	// API LoadBalancer fip. IP address field.
	APILBFip *string `json:"apiLbFip,omitempty" tf:"api_lb_fip,omitempty"`

	// optional string →  API LoadBalancer vip. IP address field.
	// API LoadBalancer vip. IP address field.
	APILBVip *string `json:"apiLbVip,omitempty" tf:"api_lb_vip,omitempty"`

	// required string →  Availability zone of the cluster.
	// Availability zone of the cluster.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// required string →  The UUID of the Kubernetes cluster template. It can be obtained using the cluster_template data source.
	// The UUID of the Kubernetes cluster template. It can be obtained using the cluster_template data source.
	ClusterTemplateID *string `json:"clusterTemplateId,omitempty" tf:"cluster_template_id,omitempty"`

	// optional string →  Custom DNS cluster domain. Changing this creates a new cluster.
	// Custom DNS cluster domain. Changing this creates a new cluster.
	DNSDomain *string `json:"dnsDomain,omitempty" tf:"dns_domain,omitempty"`

	// required boolean →  Floating ip is enabled.
	// Floating ip is enabled.
	FloatingIPEnabled *bool `json:"floatingIpEnabled,omitempty" tf:"floating_ip_enabled,omitempty"`

	// optional string →  Floating IP created for ingress service.
	// Floating IP created for ingress service.
	IngressFloatingIP *string `json:"ingressFloatingIp,omitempty" tf:"ingress_floating_ip,omitempty"`

	// optional string →  Addresses of registries from which you can download images without checking certificates. Changing this creates a new cluster.
	// Addresses of registries from which you can download images without checking certificates. Changing this creates a new cluster.
	InsecureRegistries []*string `json:"insecureRegistries,omitempty" tf:"insecure_registries,omitempty"`

	// optional string →  The name of the Compute service SSH keypair. Changing this creates a new cluster.
	// The name of the Compute service SSH keypair. Changing this creates a new cluster.
	Keypair *string `json:"keypair,omitempty" tf:"keypair,omitempty"`

	// optional map of string →  The list of optional key value pairs representing additional properties of the cluster. Note: Updating this attribute will not immediately apply the changes; these options will be used when recreating or deleting cluster nodes, for example, during an upgrade operation.
	// The list of optional key value pairs representing additional properties of the cluster. _note_ Updating this attribute will not immediately apply the changes; these options will be used when recreating or deleting cluster nodes, for example, during an upgrade operation.
	//
	// * `calico_ipv4pool` to set subnet where pods will be created. Default 10.100.0.0/16. _note_ Updating this value while the cluster is running is dangerous because it can lead to loss of connectivity of the cluster nodes.
	// * `clean_volumes` to remove pvc volumes when deleting a cluster. Default False. _note_ Changes to this value will be applied immediately.
	// * `cloud_monitoring` to enable cloud monitoring feature. Default False.
	// * `etcd_volume_size` to set etcd volume size in GB. Default 10.
	// * `kube_log_level` to set log level for kubelet in range 0 to 8. Default 0.
	// * `master_volume_size` to set master vm volume size in GB. Default 50.
	// * `cluster_node_volume_type` to set master vm volume type. Default ceph-hdd.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// optional number →  The number of master nodes for the cluster. Changing this creates a new cluster.
	// The number of master nodes for the cluster. Changing this creates a new cluster.
	MasterCount *float64 `json:"masterCount,omitempty" tf:"master_count,omitempty"`

	// optional string →  The UUID of a flavor for the master nodes. If master_flavor is not present, value from cluster_template will be used.
	// The UUID of a flavor for the master nodes. If master_flavor is not present, value from cluster_template will be used.
	MasterFlavor *string `json:"masterFlavor,omitempty" tf:"master_flavor,omitempty"`

	// zA-Z][a-zA-Z0-9_.-]*$.
	// The name of the cluster. Changing this creates a new cluster. Should match the pattern `^[a-zA-Z][a-zA-Z0-9_.-]*$`.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// optional string →  Network cidr of k8s virtual network
	// Network cidr of k8s virtual network
	PodsNetworkCidr *string `json:"podsNetworkCidr,omitempty" tf:"pods_network_cidr,omitempty"`

	// optional string →  Region to use for the cluster. Default is a region configured for provider.
	// Region to use for the cluster. Default is a region configured for provider.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  Docker registry access password.
	// Docker registry access password.
	RegistryAuthPassword *string `json:"registryAuthPassword,omitempty" tf:"registry_auth_password,omitempty"`

	// optional string →  Current state of a cluster. Changing this to RUNNING or SHUTOFF will turn cluster on/off.
	// Current state of a cluster. Changing this to `RUNNING` or `SHUTOFF` will turn cluster on/off.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// optional boolean →  Enables syncing of security policies of cluster. Default value is false.New since v0.7.0.
	// Enables syncing of security policies of cluster. Default value is false._new_since_v0.7.0_.
	SyncSecurityPolicy *bool `json:"syncSecurityPolicy,omitempty" tf:"sync_security_policy,omitempty"`
}

type ClusterObservation struct {

	// string →  COE API address.
	// COE API address.
	APIAddress *string `json:"apiAddress,omitempty" tf:"api_address,omitempty"`

	// optional string →  API LoadBalancer fip. IP address field.
	// API LoadBalancer fip. IP address field.
	APILBFip *string `json:"apiLbFip,omitempty" tf:"api_lb_fip,omitempty"`

	// optional string →  API LoadBalancer vip. IP address field.
	// API LoadBalancer vip. IP address field.
	APILBVip *string `json:"apiLbVip,omitempty" tf:"api_lb_vip,omitempty"`

	// only map of all cluster labels.New since v0.5.1.
	// The read-only map of all cluster labels._new_since_v0.5.1_.
	AllLabels map[string]*string `json:"allLabels,omitempty" tf:"all_labels,omitempty"`

	// required string →  Availability zone of the cluster.
	// Availability zone of the cluster.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// required string →  The UUID of the Kubernetes cluster template. It can be obtained using the cluster_template data source.
	// The UUID of the Kubernetes cluster template. It can be obtained using the cluster_template data source.
	ClusterTemplateID *string `json:"clusterTemplateId,omitempty" tf:"cluster_template_id,omitempty"`

	// string →  The time at which cluster was created.
	// The time at which cluster was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// optional string →  Custom DNS cluster domain. Changing this creates a new cluster.
	// Custom DNS cluster domain. Changing this creates a new cluster.
	DNSDomain *string `json:"dnsDomain,omitempty" tf:"dns_domain,omitempty"`

	// required boolean →  Floating ip is enabled.
	// Floating ip is enabled.
	FloatingIPEnabled *bool `json:"floatingIpEnabled,omitempty" tf:"floating_ip_enabled,omitempty"`

	// string →  ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// optional string →  Floating IP created for ingress service.
	// Floating IP created for ingress service.
	IngressFloatingIP *string `json:"ingressFloatingIp,omitempty" tf:"ingress_floating_ip,omitempty"`

	// optional string →  Addresses of registries from which you can download images without checking certificates. Changing this creates a new cluster.
	// Addresses of registries from which you can download images without checking certificates. Changing this creates a new cluster.
	InsecureRegistries []*string `json:"insecureRegistries,omitempty" tf:"insecure_registries,omitempty"`

	// optional string →  The name of the Compute service SSH keypair. Changing this creates a new cluster.
	// The name of the Compute service SSH keypair. Changing this creates a new cluster.
	Keypair *string `json:"keypair,omitempty" tf:"keypair,omitempty"`

	// optional map of string →  The list of optional key value pairs representing additional properties of the cluster. Note: Updating this attribute will not immediately apply the changes; these options will be used when recreating or deleting cluster nodes, for example, during an upgrade operation.
	// The list of optional key value pairs representing additional properties of the cluster. _note_ Updating this attribute will not immediately apply the changes; these options will be used when recreating or deleting cluster nodes, for example, during an upgrade operation.
	//
	// * `calico_ipv4pool` to set subnet where pods will be created. Default 10.100.0.0/16. _note_ Updating this value while the cluster is running is dangerous because it can lead to loss of connectivity of the cluster nodes.
	// * `clean_volumes` to remove pvc volumes when deleting a cluster. Default False. _note_ Changes to this value will be applied immediately.
	// * `cloud_monitoring` to enable cloud monitoring feature. Default False.
	// * `etcd_volume_size` to set etcd volume size in GB. Default 10.
	// * `kube_log_level` to set log level for kubelet in range 0 to 8. Default 0.
	// * `master_volume_size` to set master vm volume size in GB. Default 50.
	// * `cluster_node_volume_type` to set master vm volume type. Default ceph-hdd.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// optional string →  The UUID of the load balancer's subnet. Changing this creates new cluster.
	// The UUID of the load balancer's subnet. Changing this creates new cluster.
	LoadbalancerSubnetID *string `json:"loadbalancerSubnetId,omitempty" tf:"loadbalancer_subnet_id,omitempty"`

	// string →  IP addresses of the master node of the cluster.
	// IP addresses of the master node of the cluster.
	MasterAddresses []*string `json:"masterAddresses,omitempty" tf:"master_addresses,omitempty"`

	// optional number →  The number of master nodes for the cluster. Changing this creates a new cluster.
	// The number of master nodes for the cluster. Changing this creates a new cluster.
	MasterCount *float64 `json:"masterCount,omitempty" tf:"master_count,omitempty"`

	// optional string →  The UUID of a flavor for the master nodes. If master_flavor is not present, value from cluster_template will be used.
	// The UUID of a flavor for the master nodes. If master_flavor is not present, value from cluster_template will be used.
	MasterFlavor *string `json:"masterFlavor,omitempty" tf:"master_flavor,omitempty"`

	// zA-Z][a-zA-Z0-9_.-]*$.
	// The name of the cluster. Changing this creates a new cluster. Should match the pattern `^[a-zA-Z][a-zA-Z0-9_.-]*$`.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// required string →  UUID of the cluster's network.
	// UUID of the cluster's network.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// optional string →  Network cidr of k8s virtual network
	// Network cidr of k8s virtual network
	PodsNetworkCidr *string `json:"podsNetworkCidr,omitempty" tf:"pods_network_cidr,omitempty"`

	// string →  The project of the cluster.
	// The project of the cluster.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// optional string →  Region to use for the cluster. Default is a region configured for provider.
	// Region to use for the cluster. Default is a region configured for provider.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  Docker registry access password.
	// Docker registry access password.
	RegistryAuthPassword *string `json:"registryAuthPassword,omitempty" tf:"registry_auth_password,omitempty"`

	// string →  UUID of the Orchestration service stack.
	// UUID of the Orchestration service stack.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// optional string →  Current state of a cluster. Changing this to RUNNING or SHUTOFF will turn cluster on/off.
	// Current state of a cluster. Changing this to `RUNNING` or `SHUTOFF` will turn cluster on/off.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// required string →  UUID of the cluster's subnet.
	// UUID of the cluster's subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// optional boolean →  Enables syncing of security policies of cluster. Default value is false.New since v0.7.0.
	// Enables syncing of security policies of cluster. Default value is false._new_since_v0.7.0_.
	SyncSecurityPolicy *bool `json:"syncSecurityPolicy,omitempty" tf:"sync_security_policy,omitempty"`

	// string →  The time at which cluster was created.
	// The time at which cluster was created.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// string →  The user of the cluster.
	// The user of the cluster.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type ClusterParameters struct {

	// optional string →  API LoadBalancer fip. IP address field.
	// API LoadBalancer fip. IP address field.
	// +kubebuilder:validation:Optional
	APILBFip *string `json:"apiLbFip,omitempty" tf:"api_lb_fip,omitempty"`

	// optional string →  API LoadBalancer vip. IP address field.
	// API LoadBalancer vip. IP address field.
	// +kubebuilder:validation:Optional
	APILBVip *string `json:"apiLbVip,omitempty" tf:"api_lb_vip,omitempty"`

	// required string →  Availability zone of the cluster.
	// Availability zone of the cluster.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// required string →  The UUID of the Kubernetes cluster template. It can be obtained using the cluster_template data source.
	// The UUID of the Kubernetes cluster template. It can be obtained using the cluster_template data source.
	// +kubebuilder:validation:Optional
	ClusterTemplateID *string `json:"clusterTemplateId,omitempty" tf:"cluster_template_id,omitempty"`

	// optional string →  Custom DNS cluster domain. Changing this creates a new cluster.
	// Custom DNS cluster domain. Changing this creates a new cluster.
	// +kubebuilder:validation:Optional
	DNSDomain *string `json:"dnsDomain,omitempty" tf:"dns_domain,omitempty"`

	// required boolean →  Floating ip is enabled.
	// Floating ip is enabled.
	// +kubebuilder:validation:Optional
	FloatingIPEnabled *bool `json:"floatingIpEnabled,omitempty" tf:"floating_ip_enabled,omitempty"`

	// optional string →  Floating IP created for ingress service.
	// Floating IP created for ingress service.
	// +kubebuilder:validation:Optional
	IngressFloatingIP *string `json:"ingressFloatingIp,omitempty" tf:"ingress_floating_ip,omitempty"`

	// optional string →  Addresses of registries from which you can download images without checking certificates. Changing this creates a new cluster.
	// Addresses of registries from which you can download images without checking certificates. Changing this creates a new cluster.
	// +kubebuilder:validation:Optional
	InsecureRegistries []*string `json:"insecureRegistries,omitempty" tf:"insecure_registries,omitempty"`

	// optional string →  The name of the Compute service SSH keypair. Changing this creates a new cluster.
	// The name of the Compute service SSH keypair. Changing this creates a new cluster.
	// +kubebuilder:validation:Optional
	Keypair *string `json:"keypair,omitempty" tf:"keypair,omitempty"`

	// optional map of string →  The list of optional key value pairs representing additional properties of the cluster. Note: Updating this attribute will not immediately apply the changes; these options will be used when recreating or deleting cluster nodes, for example, during an upgrade operation.
	// The list of optional key value pairs representing additional properties of the cluster. _note_ Updating this attribute will not immediately apply the changes; these options will be used when recreating or deleting cluster nodes, for example, during an upgrade operation.
	//
	// * `calico_ipv4pool` to set subnet where pods will be created. Default 10.100.0.0/16. _note_ Updating this value while the cluster is running is dangerous because it can lead to loss of connectivity of the cluster nodes.
	// * `clean_volumes` to remove pvc volumes when deleting a cluster. Default False. _note_ Changes to this value will be applied immediately.
	// * `cloud_monitoring` to enable cloud monitoring feature. Default False.
	// * `etcd_volume_size` to set etcd volume size in GB. Default 10.
	// * `kube_log_level` to set log level for kubelet in range 0 to 8. Default 0.
	// * `master_volume_size` to set master vm volume size in GB. Default 50.
	// * `cluster_node_volume_type` to set master vm volume type. Default ceph-hdd.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// optional string →  The UUID of the load balancer's subnet. Changing this creates new cluster.
	// The UUID of the load balancer's subnet. Changing this creates new cluster.
	// +crossplane:generate:reference:type=github.com/viletay/provider-vkcs/apis/network/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	LoadbalancerSubnetID *string `json:"loadbalancerSubnetId,omitempty" tf:"loadbalancer_subnet_id,omitempty"`

	// Reference to a Subnet in network to populate loadbalancerSubnetId.
	// +kubebuilder:validation:Optional
	LoadbalancerSubnetIDRef *v1.Reference `json:"loadbalancerSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate loadbalancerSubnetId.
	// +kubebuilder:validation:Optional
	LoadbalancerSubnetIDSelector *v1.Selector `json:"loadbalancerSubnetIdSelector,omitempty" tf:"-"`

	// optional number →  The number of master nodes for the cluster. Changing this creates a new cluster.
	// The number of master nodes for the cluster. Changing this creates a new cluster.
	// +kubebuilder:validation:Optional
	MasterCount *float64 `json:"masterCount,omitempty" tf:"master_count,omitempty"`

	// optional string →  The UUID of a flavor for the master nodes. If master_flavor is not present, value from cluster_template will be used.
	// The UUID of a flavor for the master nodes. If master_flavor is not present, value from cluster_template will be used.
	// +kubebuilder:validation:Optional
	MasterFlavor *string `json:"masterFlavor,omitempty" tf:"master_flavor,omitempty"`

	// zA-Z][a-zA-Z0-9_.-]*$.
	// The name of the cluster. Changing this creates a new cluster. Should match the pattern `^[a-zA-Z][a-zA-Z0-9_.-]*$`.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// required string →  UUID of the cluster's network.
	// UUID of the cluster's network.
	// +crossplane:generate:reference:type=github.com/viletay/provider-vkcs/apis/network/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// optional string →  Network cidr of k8s virtual network
	// Network cidr of k8s virtual network
	// +kubebuilder:validation:Optional
	PodsNetworkCidr *string `json:"podsNetworkCidr,omitempty" tf:"pods_network_cidr,omitempty"`

	// optional string →  Region to use for the cluster. Default is a region configured for provider.
	// Region to use for the cluster. Default is a region configured for provider.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// optional string →  Docker registry access password.
	// Docker registry access password.
	// +kubebuilder:validation:Optional
	RegistryAuthPassword *string `json:"registryAuthPassword,omitempty" tf:"registry_auth_password,omitempty"`

	// optional string →  Current state of a cluster. Changing this to RUNNING or SHUTOFF will turn cluster on/off.
	// Current state of a cluster. Changing this to `RUNNING` or `SHUTOFF` will turn cluster on/off.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// required string →  UUID of the cluster's subnet.
	// UUID of the cluster's subnet.
	// +crossplane:generate:reference:type=github.com/viletay/provider-vkcs/apis/network/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// optional boolean →  Enables syncing of security policies of cluster. Default value is false.New since v0.7.0.
	// Enables syncing of security policies of cluster. Default value is false._new_since_v0.7.0_.
	// +kubebuilder:validation:Optional
	SyncSecurityPolicy *bool `json:"syncSecurityPolicy,omitempty" tf:"sync_security_policy,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
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
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. Manages a kubernetes cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vkcs}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clusterTemplateId) || (has(self.initProvider) && has(self.initProvider.clusterTemplateId))",message="spec.forProvider.clusterTemplateId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.floatingIpEnabled) || (has(self.initProvider) && has(self.initProvider.floatingIpEnabled))",message="spec.forProvider.floatingIpEnabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ClusterSpec   `json:"spec"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}