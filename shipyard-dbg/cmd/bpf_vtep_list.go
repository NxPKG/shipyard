// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/khulnasoft/shipyard/pkg/command"
	"github.com/khulnasoft/shipyard/pkg/common"
	"github.com/khulnasoft/shipyard/pkg/maps/vtep"
)

const (
	vtepCidrTitle = "IP PREFIX/ADDRESS"
	vtepTitle     = "VTEP"
)

var (
	vtepListUsage = "List VTEP CIDR and their corresponding VTEP MAC/IP.\n"
)

var bpfVtepListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List VTEP CIDR and their corresponding VTEP MAC/IP",
	Long:    vtepListUsage,
	Run: func(cmd *cobra.Command, args []string) {
		common.RequireRootPrivilege("cilium bpf vtep list")

		bpfVtepList := make(map[string][]string)
		if err := vtep.VtepMap().Dump(bpfVtepList); err != nil {
			fmt.Fprintf(os.Stderr, "error dumping contents of map: %s\n", err)
			os.Exit(1)
		}

		if command.OutputOption() {
			if err := command.PrintOutput(bpfVtepList); err != nil {
				fmt.Fprintf(os.Stderr, "error getting output of map in %s: %s\n", command.OutputOptionString(), err)
				os.Exit(1)
			}
			return
		}

		if len(bpfVtepList) == 0 {
			fmt.Fprintf(os.Stderr, "No entries found.\n")
		} else {
			TablePrinter(vtepCidrTitle, vtepTitle, bpfVtepList)
		}
	},
}

func init() {
	BPFVtepCmd.AddCommand(bpfVtepListCmd)
	command.AddOutputOption(bpfVtepListCmd)
}
