// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/khulnasoft/shipyard/pkg/command"
	"github.com/khulnasoft/shipyard/pkg/common"
	"github.com/khulnasoft/shipyard/pkg/maps/tunnel"
)

const (
	tunnelTitle      = "TUNNEL"
	destinationTitle = "VALUE"
)

var bpfTunnelListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List tunnel endpoint entries",
	Run: func(cmd *cobra.Command, args []string) {
		common.RequireRootPrivilege("cilium bpf tunnel list")

		tunnelList := make(map[string][]string)
		if err := tunnel.TunnelMap().Dump(tunnelList); err != nil {
			os.Exit(1)
		}

		if command.OutputOption() {
			if err := command.PrintOutput(tunnelList); err != nil {
				os.Exit(1)
			}
			return
		}

		TablePrinter(tunnelTitle, destinationTitle, tunnelList)
	},
}

func init() {
	BPFTunnelCmd.AddCommand(bpfTunnelListCmd)
	command.AddOutputOption(bpfTunnelListCmd)
}
