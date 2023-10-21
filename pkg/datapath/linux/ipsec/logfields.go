// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package ipsec

import (
	"github.com/khulnasoft/shipyard/pkg/logging"
	"github.com/khulnasoft/shipyard/pkg/logging/logfields"
)

const subsystem = "ipsec"

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, subsystem)
