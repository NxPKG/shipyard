// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package groups

import (
	"testing"

	. "github.com/khulnasoft/defeat"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	TestingT(t)
}

type GroupsTestSuite struct{}

var _ = Suite(&GroupsTestSuite{})
