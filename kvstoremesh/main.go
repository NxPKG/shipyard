// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	cmmetrics "github.com/khulnasoft/shipyard/clustermesh-apiserver/metrics"
	kmopt "github.com/khulnasoft/shipyard/kvstoremesh/option"
	"github.com/khulnasoft/shipyard/pkg/clustermesh/kvstoremesh"
	"github.com/khulnasoft/shipyard/pkg/clustermesh/types"
	cmtypes "github.com/khulnasoft/shipyard/pkg/clustermesh/types"
	"github.com/khulnasoft/shipyard/pkg/controller"
	"github.com/khulnasoft/shipyard/pkg/defaults"
	"github.com/khulnasoft/shipyard/pkg/gops"
	"github.com/khulnasoft/shipyard/pkg/hive"
	"github.com/khulnasoft/shipyard/pkg/hive/cell"
	"github.com/khulnasoft/shipyard/pkg/kvstore"
	"github.com/khulnasoft/shipyard/pkg/logging"
	"github.com/khulnasoft/shipyard/pkg/logging/logfields"
	"github.com/khulnasoft/shipyard/pkg/metrics"
	"github.com/khulnasoft/shipyard/pkg/option"
	"github.com/khulnasoft/shipyard/pkg/pprof"
)

var (
	log = logging.DefaultLogger.WithField(logfields.LogSubsys, "kvstoremesh")

	rootHive *hive.Hive

	rootCmd = &cobra.Command{
		Use:   "kvstoremesh",
		Short: "Run KVStoreMesh",
		Run: func(cmd *cobra.Command, args []string) {
			if err := rootHive.Run(); err != nil {
				log.Fatal(err)
			}
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			// Overwrite the metrics namespace with the one specific for KVStoreMesh
			metrics.Namespace = metrics.CiliumKVStoreMeshNamespace
			option.Config.Populate(rootHive.Viper())
			if option.Config.Debug {
				log.Logger.SetLevel(logrus.DebugLevel)
			}
			option.LogRegisteredOptions(rootHive.Viper(), log)
		},
	}
)

func init() {
	rootHive = hive.New(
		pprof.Cell,
		cell.Config(pprof.Config{
			PprofAddress: kmopt.PprofAddress,
			PprofPort:    kmopt.PprofPort,
		}),
		controller.Cell,

		gops.Cell(defaults.GopsPortKVStoreMesh),
		cmmetrics.Cell,

		cell.Config(kmopt.KVStoreMeshConfig{}),
		cell.Config(cmtypes.DefaultClusterInfo),
		cell.Invoke(registerClusterInfoValidator),

		kvstore.Cell(kvstore.EtcdBackendName),
		cell.Provide(func() *kvstore.ExtraOptions { return nil }),
		kvstoremesh.Cell,

		cell.Invoke(func(*kvstoremesh.KVStoreMesh) {}),
	)

	rootHive.RegisterFlags(rootCmd.Flags())
	rootCmd.AddCommand(rootHive.Command())
}

func registerClusterInfoValidator(lc hive.Lifecycle, cinfo types.ClusterInfo) {
	lc.Append(hive.Hook{
		OnStart: func(hive.HookContext) error { return cinfo.ValidateStrict() },
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
