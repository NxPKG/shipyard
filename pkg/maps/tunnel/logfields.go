// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package tunnel

import (
	"github.com/khulnasoft/shipyard/pkg/logging"
	"github.com/khulnasoft/shipyard/pkg/logging/logfields"
)

const (
	fieldEndpoint = "endpoint"
	fieldPrefix   = "prefix"
	fieldKey      = "key"
	fieldNodeID   = "nodeID"
)

var (
	log = logging.DefaultLogger.WithField(logfields.LogSubsys, "map-tunnel")
)
