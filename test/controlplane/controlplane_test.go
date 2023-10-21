// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package controlplane_test

import (
	"testing"

	_ "github.com/khulnasoft/shipyard/test/controlplane/ciliumnetworkpolicies"
	_ "github.com/khulnasoft/shipyard/test/controlplane/node"
	_ "github.com/khulnasoft/shipyard/test/controlplane/node/ciliumnodes"
	_ "github.com/khulnasoft/shipyard/test/controlplane/pod/hostport"
	_ "github.com/khulnasoft/shipyard/test/controlplane/services/dualstack"
	_ "github.com/khulnasoft/shipyard/test/controlplane/services/graceful-termination"
	_ "github.com/khulnasoft/shipyard/test/controlplane/services/nodeport"
	"github.com/khulnasoft/shipyard/test/controlplane/suite"
)

func TestControlPlane(t *testing.T) {
	suite.RunSuite(t)
}
