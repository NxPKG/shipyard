// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build linux

package datapath

import (
	"github.com/khulnasoft/shipyard/pkg/testutils"

	. "github.com/khulnasoft/defeat"
)

type DaemonPrivilegedSuite struct{}

var _ = Suite(&DaemonPrivilegedSuite{})

func (s *DaemonPrivilegedSuite) SetUpSuite(c *C) {
	testutils.PrivilegedTest(c)
}

func (s *DaemonPrivilegedSuite) TestEnableIPForwarding(c *C) {
	err := enableIPForwarding()
	c.Assert(err, IsNil)
}
