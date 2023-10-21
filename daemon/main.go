// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package main

import (
	"github.com/khulnasoft/shipyard/daemon/cmd"
	"github.com/khulnasoft/shipyard/pkg/hive"
)

func main() {
	agentHive := hive.New(cmd.Agent)

	cmd.Execute(cmd.NewAgentCmd(agentHive))
}
