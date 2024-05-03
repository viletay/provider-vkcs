package network

import (
	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/resource"
)

const (
	// Group is the short group for this provider.
	Group = "network"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("vkcs_networking_network", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
	p.AddResourceConfigurator("vkcs_networking_subnet", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["network_id"] = ujconfig.Reference{
			Type: "Network",
		}
	})
	p.AddResourceConfigurator("vkcs_networking_subnet_route", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["subnet_id"] = ujconfig.Reference{
			Type: "Subnet",
		}
	})
	p.AddResourceConfigurator("vkcs_networking_secgroup", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "SecurityGroup"
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
	p.AddResourceConfigurator("vkcs_networking_secgroup_rule", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "SecurityGroupRule"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["security_group_id"] = ujconfig.Reference{
			Type: "SecurityGroup",
		}
		r.References["remote_group_id"] = ujconfig.Reference{
			Type: "SecurityGroup",
		}
	})
	p.AddResourceConfigurator("vkcs_networking_port", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["network_id"] = ujconfig.Reference{
			Type: "Network",
		}
		r.References["fixed_ip.subnet_id"] = ujconfig.Reference{
			Type: "Subnet",
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: "SecurityGroup",
		}
	})
	p.AddResourceConfigurator("vkcs_networking_port_secgroup_associate", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "SecurityGroupAssociate"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["port_id"] = ujconfig.Reference{
			Type: "Port",
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: "SecurityGroup",
		}
	})
	p.AddResourceConfigurator("vkcs_networking_router", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["external_network_id"] = ujconfig.Reference{
			Type: "Network",
		}
	})
	p.AddResourceConfigurator("vkcs_networking_router_interface", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "RouterInterface"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["router_id"] = ujconfig.Reference{
			Type: "Router",
		}
		r.References["port_id"] = ujconfig.Reference{
			Type: "Port",
		}
		r.References["subnet_id"] = ujconfig.Reference{
			Type: "Subnet",
		}

	})
	p.AddResourceConfigurator("vkcs_networking_router_route", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "RouterRoute"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["router_id"] = ujconfig.Reference{
			Type: "Router",
		}

	})
	p.AddResourceConfigurator("vkcs_networking_floatingip", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "FloatingIP"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["pool"] = ujconfig.Reference{
			Type:      "Network",
			Extractor: "github.com/viletay/provider-vkcs/config/network.ExtractSpecName()",
		}
		r.References["port_id"] = ujconfig.Reference{
			Type: "Port",
		}
		r.References["subnet_id"] = ujconfig.Reference{
			Type: "Subnet",
		}
		r.References["subnet_ids"] = ujconfig.Reference{
			Type: "Subnet",
		}
	})
	p.AddResourceConfigurator("vkcs_networking_floatingip_associate", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "FloatingIPAssociate"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["floating_ip"] = ujconfig.Reference{
			Type:      "FloatingIP",
			Extractor: "github.com/viletay/provider-vkcs/config/network.ExtractSpecAddress()",
		}
		r.References["port_id"] = ujconfig.Reference{
			Type: "Port",
		}
	})
}

// ExtractSpecName extracts the value of `spec.forProvider.name`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractSpecName() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetObservation()
		if err != nil {
			return ""
		}
		if k := o["name"]; k != nil {
			return k.(string)
		}
		return ""
	}
}

// ExtractSpecAddress extracts the value of `spec.forProvider.address`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractSpecAddress() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetObservation()
		if err != nil {
			return ""
		}
		if k := o["address"]; k != nil {
			return k.(string)
		}
		return ""
	}
}
