// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package azure

import (
	chainingapi "github.com/khulnasoft/shipyard/plugins/cilium-cni/chaining/api"
	genericveth "github.com/khulnasoft/shipyard/plugins/cilium-cni/chaining/generic-veth"
)

func init() {
	chainingapi.Register("azure", &genericveth.GenericVethChainer{})
}
