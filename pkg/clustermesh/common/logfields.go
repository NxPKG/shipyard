// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package common

import (
	"github.com/khulnasoft/shipyard/pkg/logging"
	"github.com/khulnasoft/shipyard/pkg/logging/logfields"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "clustermesh")

const (
	fieldClusterName   = "clusterName"
	fieldConfig        = "config"
	fieldConfigDir     = "configDir"
	fieldEvent         = "event"
	fieldKVStoreStatus = "kvstoreStatus"
	fieldKVStoreErr    = "kvstoreErr"
)
