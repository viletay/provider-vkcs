package dns

import (
	"context"
	"fmt"
	"strings"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("vkcs_publicdns_zone", func(r *ujconfig.Resource) {
		r.ShortGroup = "dns"
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
	p.AddResourceConfigurator("vkcs_publicdns_record", func(r *ujconfig.Resource) {
		r.ShortGroup = "dns"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.ExternalName.GetExternalNameFn = getRecordExternalNameFn
		r.ExternalName.GetIDFn = getRecordIDFn
		r.References["zone_id"] = ujconfig.Reference{
			Type: "Zone",
		}
	})
}

func getRecordExternalNameFn(tfstate map[string]any) (string, error) {
	id, ok := tfstate["id"]
	if !ok {
		return "", errors.New("undefined attribute: id")
	}
	idStr, ok := id.(string)
	if !ok {
		return "", errors.New("unexpected type: id")
	}

	zoneID, ok := tfstate["zone_id"]
	if !ok {
		return "", errors.New("undefined attribute: zone_id")
	}
	zoneIDStr, ok := zoneID.(string)
	if !ok {
		return "", errors.New("unexpected type: zone_id")
	}

	recordType, ok := tfstate["type"]
	if !ok {
		return "", errors.New("undefined attribute: type")
	}
	recordTypeStr, ok := recordType.(string)
	if !ok {
		return "", errors.New("unexpected type: type")
	}
	return fmt.Sprintf("%s/%s/%s", zoneIDStr, recordTypeStr, idStr), nil
}

func getRecordIDFn(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
	idStr := strings.Split(externalName, "/")
	return idStr[len(idStr)-1], nil
}
