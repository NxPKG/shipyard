// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build !race

package idpool

import (
	. "github.com/khulnasoft/defeat"
)

func (s *IDPoolTestSuite) TestAllocateID(c *C) {
	s.testAllocatedID(c, 256)
}
