/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"vkcs_publicdns_zone":                     config.IdentifierFromProvider,
	"vkcs_publicdns_record":                   config.IdentifierFromProvider,
	"vkcs_images_image":                       config.IdentifierFromProvider,
	"vkcs_networking_network":                 config.IdentifierFromProvider,
	"vkcs_networking_subnet":                  config.IdentifierFromProvider,
	"vkcs_networking_secgroup":                config.IdentifierFromProvider,
	"vkcs_networking_secgroup_rule":           config.IdentifierFromProvider,
	"vkcs_networking_subnet_route":            config.IdentifierFromProvider,
	"vkcs_networking_port":                    config.IdentifierFromProvider,
	"vkcs_networking_port_secgroup_associate": config.IdentifierFromProvider,
	"vkcs_networking_router":                  config.IdentifierFromProvider,
	"vkcs_networking_router_interface":        config.IdentifierFromProvider,
	"vkcs_networking_router_route":            config.IdentifierFromProvider,
	"vkcs_networking_floatingip":              config.IdentifierFromProvider,
	"vkcs_networking_floatingip_associate":    config.IdentifierFromProvider,
	"vkcs_blockstorage_volume":                config.IdentifierFromProvider,
	"vkcs_blockstorage_snapshot":              config.IdentifierFromProvider,
	"vkcs_compute_keypair":                    config.IdentifierFromProvider,
	"vkcs_compute_instance":                   config.IdentifierFromProvider,
	"vkcs_compute_volume_attach":              config.IdentifierFromProvider,
	"vkcs_compute_interface_attach":           config.IdentifierFromProvider,
	"vkcs_compute_floatingip_associate":       config.IdentifierFromProvider,
	"vkcs_compute_servergroup":                config.IdentifierFromProvider,
	"vkcs_kubernetes_cluster":                 config.IdentifierFromProvider,
	"vkcs_kubernetes_node_group":              config.IdentifierFromProvider,
	"vkcs_kubernetes_clustertemplate":         config.IdentifierFromProvider,
	"vkcs_lb_loadbalancer":                    config.IdentifierFromProvider,
	"vkcs_lb_listener":                        config.IdentifierFromProvider,
	"vkcs_lb_pool":                            config.IdentifierFromProvider,
	"vkcs_lb_member":                          config.IdentifierFromProvider,
	"vkcs_lb_monitor":                         config.IdentifierFromProvider,
	"vkcs_lb_l7policy":                        config.IdentifierFromProvider,
	"vkcs_lb_l7rule":                          config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
