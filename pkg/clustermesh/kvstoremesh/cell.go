// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package kvstoremesh

import (
	"github.com/khulnasoft/shipyard/pkg/clustermesh/common"
	"github.com/khulnasoft/shipyard/pkg/hive/cell"
	"github.com/khulnasoft/shipyard/pkg/kvstore/store"
)

var Cell = cell.Module(
	"kvstoremesh",
	"KVStoreMesh caches remote cluster information in a local kvstore",

	cell.Provide(newKVStoreMesh),

	cell.Config(common.Config{}),
	store.Cell,

	cell.Metric(common.MetricsProvider("")),
)
