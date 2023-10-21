// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package main

import (
	"github.com/khulnasoft/shipyard/operator/cmd"
	"github.com/khulnasoft/shipyard/pkg/hive"
)

func main() {
	operatorHive := hive.New(cmd.Operator)

	cmd.Execute(cmd.NewOperatorCmd(operatorHive))
}
