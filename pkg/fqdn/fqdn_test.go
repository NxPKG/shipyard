// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fqdn

import (
	"testing"

	. "github.com/khulnasoft/defeat"

	"github.com/khulnasoft/shipyard/pkg/defaults"
	"github.com/khulnasoft/shipyard/pkg/fqdn/re"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	TestingT(t)
}

func (ds *FQDNTestSuite) SetUpSuite(c *C) {
	re.InitRegexCompileLRU(defaults.FQDNRegexCompileLRUSize)
}

type FQDNTestSuite struct{}

var _ = Suite(&FQDNTestSuite{})
