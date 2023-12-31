// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Package service contains the logic for Cilium Load Balancer Controller via envoy config
package ciliumenvoyconfig

import (
	"github.com/khulnasoft/shipyard/pkg/logging"
	"github.com/khulnasoft/shipyard/pkg/logging/logfields"
)

const Subsys = "envoy-lb-controller"

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, Subsys)
