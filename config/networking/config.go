package networking

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("vkcs_networking_network", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
	p.AddResourceConfigurator("vkcs_networking_subnet", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["network_id"] = ujconfig.Reference{
			Type: "Network",
		}
	})
	p.AddResourceConfigurator("vkcs_networking_secgroup", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
	p.AddResourceConfigurator("vkcs_networking_secgroup_rule", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["security_group_id"] = ujconfig.Reference{
			Type: "Secgroup",
		}
		r.References["remote_group_id"] = ujconfig.Reference{
			Type: "Secgroup",
		}
	})

}
