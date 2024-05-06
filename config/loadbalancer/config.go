package loadbalancer

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "loadbalancer"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("vkcs_lb_loadbalancer", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "LoadBalancer"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/network/v1alpha1.SecurityGroup",
		}
		r.References["vip_network_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/network/v1alpha1.Network",
		}
		r.References["vip_subnet_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/network/v1alpha1.Subnet",
		}
		r.References["vip_port_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/network/v1alpha1.Port",
		}

	})
	p.AddResourceConfigurator("vkcs_lb_listener", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["loadbalancer_id"] = ujconfig.Reference{
			Type: "LoadBalancer",
		}
	})
	p.AddResourceConfigurator("vkcs_lb_listener", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["loadbalancer_id"] = ujconfig.Reference{
			Type: "LoadBalancer",
		}
	})
	p.AddResourceConfigurator("vkcs_lb_pool", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["loadbalancer_id"] = ujconfig.Reference{
			Type: "LoadBalancer",
		}
		r.References["listener_id"] = ujconfig.Reference{
			Type: "Listener",
		}
	})
	p.AddResourceConfigurator("vkcs_lb_member", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["pool_id"] = ujconfig.Reference{
			Type: "Pool",
		}
		r.References["subnet_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-vkcs/apis/network/v1alpha1.Subnet",
		}
	})
	p.AddResourceConfigurator("vkcs_lb_monitor", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["pool_id"] = ujconfig.Reference{
			Type: "Pool",
		}
	})
	p.AddResourceConfigurator("vkcs_lb_l7policy", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "L7Policy"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["listener_id"] = ujconfig.Reference{
			Type: "Listener",
		}
		r.References["redirect_pool_id"] = ujconfig.Reference{
			Type: "Pool",
		}
	})
	p.AddResourceConfigurator("vkcs_lb_l7rule", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "L7Rule"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["l7policy_id"] = ujconfig.Reference{
			Type: "L7Policy",
		}
	})

}
