// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package loader

import (
	. "github.com/khulnasoft/defeat"

	"github.com/khulnasoft/shipyard/pkg/node"
)

func (s *LoaderTestSuite) SetUpTest(c *C) {
	node.InitDefaultPrefix("")
	node.SetInternalIPv4Router(templateIPv4[:])
	node.SetIPv4Loopback(templateIPv4[:])
}
