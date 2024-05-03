// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	floatingipassociate "github.com/viletay/provider-vkcs/internal/controller/compute/floatingipassociate"
	instance "github.com/viletay/provider-vkcs/internal/controller/compute/instance"
	interfaceattach "github.com/viletay/provider-vkcs/internal/controller/compute/interfaceattach"
	keypair "github.com/viletay/provider-vkcs/internal/controller/compute/keypair"
	servergroup "github.com/viletay/provider-vkcs/internal/controller/compute/servergroup"
	volumeattach "github.com/viletay/provider-vkcs/internal/controller/compute/volumeattach"
	image "github.com/viletay/provider-vkcs/internal/controller/disk/image"
	snapshot "github.com/viletay/provider-vkcs/internal/controller/disk/snapshot"
	volume "github.com/viletay/provider-vkcs/internal/controller/disk/volume"
	record "github.com/viletay/provider-vkcs/internal/controller/dns/record"
	zone "github.com/viletay/provider-vkcs/internal/controller/dns/zone"
	floatingip "github.com/viletay/provider-vkcs/internal/controller/network/floatingip"
	floatingipassociatenetwork "github.com/viletay/provider-vkcs/internal/controller/network/floatingipassociate"
	network "github.com/viletay/provider-vkcs/internal/controller/network/network"
	port "github.com/viletay/provider-vkcs/internal/controller/network/port"
	router "github.com/viletay/provider-vkcs/internal/controller/network/router"
	routerinterface "github.com/viletay/provider-vkcs/internal/controller/network/routerinterface"
	routerroute "github.com/viletay/provider-vkcs/internal/controller/network/routerroute"
	securitygroup "github.com/viletay/provider-vkcs/internal/controller/network/securitygroup"
	securitygroupassociate "github.com/viletay/provider-vkcs/internal/controller/network/securitygroupassociate"
	securitygrouprule "github.com/viletay/provider-vkcs/internal/controller/network/securitygrouprule"
	subnet "github.com/viletay/provider-vkcs/internal/controller/network/subnet"
	subnetroute "github.com/viletay/provider-vkcs/internal/controller/network/subnetroute"
	providerconfig "github.com/viletay/provider-vkcs/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		floatingipassociate.Setup,
		instance.Setup,
		interfaceattach.Setup,
		keypair.Setup,
		servergroup.Setup,
		volumeattach.Setup,
		image.Setup,
		snapshot.Setup,
		volume.Setup,
		record.Setup,
		zone.Setup,
		floatingip.Setup,
		floatingipassociatenetwork.Setup,
		network.Setup,
		port.Setup,
		router.Setup,
		routerinterface.Setup,
		routerroute.Setup,
		securitygroup.Setup,
		securitygroupassociate.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		subnetroute.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
