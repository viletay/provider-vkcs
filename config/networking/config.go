package networking

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("vkcs_networking_network", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
}
