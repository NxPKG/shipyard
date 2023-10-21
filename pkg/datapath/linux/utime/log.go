// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package utime

import (
	"github.com/khulnasoft/shipyard/pkg/logging"
	"github.com/khulnasoft/shipyard/pkg/logging/logfields"
)

var (
	subsystem = "utime"
	log       = logging.DefaultLogger.WithField(logfields.LogSubsys, subsystem)
)
