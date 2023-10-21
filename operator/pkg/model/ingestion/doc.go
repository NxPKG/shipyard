// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Package ingestion holds functions that translate from
// Kubernetes resources into Listener types for storage
// in the model.
package ingestion

import (
	"github.com/khulnasoft/shipyard/pkg/logging"
	"github.com/khulnasoft/shipyard/pkg/logging/logfields"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "ingestion")
