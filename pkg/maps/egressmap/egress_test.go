// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package egressmap

import (
	"testing"

	"github.com/khulnasoft/shipyard/pkg/hive"
)

func TestCell(t *testing.T) {
	err := hive.New(Cell).Populate()
	if err != nil {
		t.Fatal(err)
	}
}
