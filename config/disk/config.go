package disk

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("vkcs_images_image", func(r *ujconfig.Resource) {
		r.ShortGroup = "disk"
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
	p.AddResourceConfigurator("vkcs_blockstorage_volume", func(r *ujconfig.Resource) {
		r.ShortGroup = "disk"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["image_id"] = ujconfig.Reference{
			Type: "Image",
		}
		r.References["source_vol_id"] = ujconfig.Reference{
			Type: "Volume",
		}
		r.References["snapshot_id"] = ujconfig.Reference{
			Type: "Snapshot",
		}
	})
	p.AddResourceConfigurator("vkcs_blockstorage_snapshot", func(r *ujconfig.Resource) {
		r.ShortGroup = "disk"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["volume_id"] = ujconfig.Reference{
			Type: "Volume",
		}
	})
}
