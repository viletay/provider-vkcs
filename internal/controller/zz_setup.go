// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	image "github.com/viletay/provider-vkcs/internal/controller/images/image"
	floatingip "github.com/viletay/provider-vkcs/internal/controller/networking/floatingip"
	floatingipassociate "github.com/viletay/provider-vkcs/internal/controller/networking/floatingipassociate"
	network "github.com/viletay/provider-vkcs/internal/controller/networking/network"
	port "github.com/viletay/provider-vkcs/internal/controller/networking/port"
	portsecgroupassociate "github.com/viletay/provider-vkcs/internal/controller/networking/portsecgroupassociate"
	router "github.com/viletay/provider-vkcs/internal/controller/networking/router"
	routerinterface "github.com/viletay/provider-vkcs/internal/controller/networking/routerinterface"
	routerroute "github.com/viletay/provider-vkcs/internal/controller/networking/routerroute"
	secgroup "github.com/viletay/provider-vkcs/internal/controller/networking/secgroup"
	secgrouprule "github.com/viletay/provider-vkcs/internal/controller/networking/secgrouprule"
	subnet "github.com/viletay/provider-vkcs/internal/controller/networking/subnet"
	subnetroute "github.com/viletay/provider-vkcs/internal/controller/networking/subnetroute"
	providerconfig "github.com/viletay/provider-vkcs/internal/controller/providerconfig"
	record "github.com/viletay/provider-vkcs/internal/controller/publicdns/record"
	zone "github.com/viletay/provider-vkcs/internal/controller/publicdns/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		image.Setup,
		floatingip.Setup,
		floatingipassociate.Setup,
		network.Setup,
		port.Setup,
		portsecgroupassociate.Setup,
		router.Setup,
		routerinterface.Setup,
		routerroute.Setup,
		secgroup.Setup,
		secgrouprule.Setup,
		subnet.Setup,
		subnetroute.Setup,
		providerconfig.Setup,
		record.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
