// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build ipam_provider_operator

package cmd

import (
	// These dependencies should be included only when this file is included in the build.
	"github.com/khulnasoft/shipyard/pkg/ipam/allocator/clusterpool"
	"github.com/khulnasoft/shipyard/pkg/ipam/allocator/multipool"

	ipamOption "github.com/khulnasoft/shipyard/pkg/ipam/option"
)

func init() {
	allocatorProviders[ipamOption.IPAMClusterPool] = &clusterpool.AllocatorOperator{}
	allocatorProviders[ipamOption.IPAMMultiPool] = &multipool.Allocator{}
}
