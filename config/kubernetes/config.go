package kubernetes

import ujconfig "github.com/crossplane/upjet/pkg/config"

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("vkcs_kubernetes_cluster", func(r *ujconfig.Resource) {
		r.ShortGroup = "kubernetes"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["network_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/network/v1alpha1.Network",
		}
		r.References["subnet_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/network/v1alpha1.Subnet",
		}
		r.References["loadbalancer_subnet_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/network/v1alpha1.Subnet",
		}

	})
	p.AddResourceConfigurator("vkcs_kubernetes_node_group", func(r *ujconfig.Resource) {
		r.ShortGroup = "kubernetes"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["cluster_id"] = ujconfig.Reference{
			Type: "Cluster",
		}
		r.References["flavor_id"] = ujconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-openstack/apis/compute/v1alpha1.FlavorV2",
		}
	})
}
