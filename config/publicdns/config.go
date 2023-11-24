package publicdns

import (
	"context"
	"fmt"
	"github.com/pkg/errors"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("vkcs_publicdns_zone", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
	p.AddResourceConfigurator("vkcs_publicdns_record", func(r *ujconfig.Resource) {
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.ExternalName.GetExternalNameFn = getNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = getFullyQualifiedIDfunc
		r.References["zone_id"] = ujconfig.Reference{
			Type: "Zone",
		}
	})

}

func getNameFromFullyQualifiedID(tfstate map[string]any) (string, error) {
	id, ok := tfstate["id"]
	if !ok {
		return "", errors.New("undefined attribute: id")
	}
	idStr, ok := id.(string)
	if !ok {
		return "", errors.New("unexpected type: id")
	}
	return idStr, nil
}

func getFullyQualifiedIDfunc(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
	zoneId, ok := parameters["zone_id"]
	if !ok {
		return "", errors.New("undefined attribute: zone_id")
	}
	zoneIdStr, ok := zoneId.(string)
	if !ok {
		return "", errors.New("unexpected type: zone_id")
	}

	recordType, ok := parameters["type"]
	if !ok {
		return "", errors.New("undefined attribute: type")
	}
	recordTypeStr, ok := recordType.(string)
	if !ok {
		return "", errors.New("unexpected type: type")
	}

	return fmt.Sprintf("%s/%s/%s", zoneIdStr, recordTypeStr, externalName), nil
}
