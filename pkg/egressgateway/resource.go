// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package egressgateway

import (
	"github.com/khulnasoft/shipyard/pkg/hive"
	v2 "github.com/khulnasoft/shipyard/pkg/k8s/apis/cilium.io/v2"
	"github.com/khulnasoft/shipyard/pkg/k8s/client"
	"github.com/khulnasoft/shipyard/pkg/k8s/resource"
	"github.com/khulnasoft/shipyard/pkg/k8s/utils"
)

type Policy = v2.CiliumEgressGatewayPolicy

func newPolicyResource(lc hive.Lifecycle, c client.Clientset) resource.Resource[*Policy] {
	if !c.IsEnabled() {
		return nil
	}
	lw := utils.ListerWatcherFromTyped[*v2.CiliumEgressGatewayPolicyList](c.CiliumV2().CiliumEgressGatewayPolicies())
	return resource.New[*Policy](lc, lw)
}
