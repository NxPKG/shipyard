// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package xds

import (
	"github.com/khulnasoft/shipyard/pkg/logging"
	"github.com/khulnasoft/shipyard/pkg/logging/logfields"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "xds")
