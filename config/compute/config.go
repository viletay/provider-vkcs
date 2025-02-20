package compute

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("vkcs_compute_keypair", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
	p.AddResourceConfigurator("vkcs_compute_instance", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["block_device.uuid"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/disk/v1alpha1.Volume",
		}
		r.References["image_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/disk/v1alpha1.Image",
		}
		r.References["network.port"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/network/v1alpha1.Port",
		}
		r.References["network.uuid"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/network/v1alpha1.Network",
		}
		r.References["flavor_id"] = ujconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-openstack/apis/compute/v1alpha1.FlavorV2",
		}
	})
	p.AddResourceConfigurator("vkcs_compute_volume_attach", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["instance_id"] = ujconfig.Reference{
			Type: "Instance",
		}
		r.References["volume_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/disk/v1alpha1.Volume",
		}

	})
	p.AddResourceConfigurator("vkcs_compute_interface_attach", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["instance_id"] = ujconfig.Reference{
			Type: "Instance",
		}
		r.References["port_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/network/v1alpha1.Port",
		}
		r.References["network_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/network/v1alpha1.Network",
		}
		r.LateInitializer = ujconfig.LateInitializer{
			IgnoredFields: []string{
				"network_id",
				"port_id",
				"fixed_ip",
			},
		}
	})
	p.AddResourceConfigurator("vkcs_compute_floatingip_associate", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.Kind = "FloatingIPAssociate"
		r.References["instance_id"] = ujconfig.Reference{
			Type: "Instance",
		}
		r.References["floating_ip"] = ujconfig.Reference{
			Type:      "github.com/viletay/provider-vkcs/apis/network/v1alpha1.FloatingIP",
			Extractor: "github.com/viletay/provider-vkcs/config/network.ExtractSpecAddress()",
		}
	})
	p.AddResourceConfigurator("vkcs_compute_servergroup", func(r *ujconfig.Resource) {
		r.Kind = "ServerGroup"
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
}
