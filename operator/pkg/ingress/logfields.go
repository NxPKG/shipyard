// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package ingress

import (
	"github.com/khulnasoft/shipyard/pkg/logging"
	"github.com/khulnasoft/shipyard/pkg/logging/logfields"
)

const Subsys = "ingress-controller"

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, Subsys)
