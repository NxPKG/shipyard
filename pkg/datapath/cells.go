// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package datapath

import (
	"log"
	"path/filepath"

	"github.com/khulnasoft/shipyard/pkg/bpf"
	"github.com/khulnasoft/shipyard/pkg/datapath/agentliveness"
	"github.com/khulnasoft/shipyard/pkg/datapath/garp"
	"github.com/khulnasoft/shipyard/pkg/datapath/iptables"
	"github.com/khulnasoft/shipyard/pkg/datapath/l2responder"
	"github.com/khulnasoft/shipyard/pkg/datapath/link"
	linuxdatapath "github.com/khulnasoft/shipyard/pkg/datapath/linux"
	dpcfg "github.com/khulnasoft/shipyard/pkg/datapath/linux/config"
	"github.com/khulnasoft/shipyard/pkg/datapath/linux/utime"
	"github.com/khulnasoft/shipyard/pkg/datapath/tables"
	"github.com/khulnasoft/shipyard/pkg/datapath/types"
	"github.com/khulnasoft/shipyard/pkg/defaults"
	"github.com/khulnasoft/shipyard/pkg/hive"
	"github.com/khulnasoft/shipyard/pkg/hive/cell"
	"github.com/khulnasoft/shipyard/pkg/maps"
	"github.com/khulnasoft/shipyard/pkg/maps/eventsmap"
	"github.com/khulnasoft/shipyard/pkg/maps/nodemap"
	monitorAgent "github.com/khulnasoft/shipyard/pkg/monitor/agent"
	"github.com/khulnasoft/shipyard/pkg/node"
	"github.com/khulnasoft/shipyard/pkg/option"
	wg "github.com/khulnasoft/shipyard/pkg/wireguard/agent"
	wgTypes "github.com/khulnasoft/shipyard/pkg/wireguard/types"
)

// Datapath provides the privileged operations to apply control-plane
// decision to the kernel.
//
// For integration testing a fake counterpart of this module is defined
// in pkg/datapath/fake/cells.go.
var Cell = cell.Module(
	"datapath",
	"Datapath",

	// Provides all BPF Map which are already provided by via hive cell.
	maps.Cell,

	// Utime synchronizes utime from userspace to datapath via configmap.Map.
	utime.Cell,

	// The cilium events map, used by the monitor agent.
	eventsmap.Cell,

	// The monitor agent, which multicasts cilium and agent events to its subscribers.
	monitorAgent.Cell,

	cell.Provide(
		newWireguardAgent,
		newDatapath,
	),

	// This cell periodically updates the agent liveness value in configmap.Map to inform
	// the datapath of the liveness of the agent.
	agentliveness.Cell,

	// The responder reconciler takes desired state about L3->L2 address translation responses and reconciles
	// it to the BPF L2 responder map.
	l2responder.Cell,

	// This cell defines StateDB tables and their schemas for tables which are used to transfer information
	// between datapath components and more high-level components.
	tables.Cell,

	// Gratuitous ARP event processor emits GARP packets on k8s pod creation events.
	garp.Cell,

	// This cell provides the object used to write the headers for datapath program types.
	dpcfg.Cell,

	cell.Provide(func(dp types.Datapath) types.NodeIDHandler {
		return dp.NodeIDs()
	}),

	// DevicesController manages the devices and routes tables
	linuxdatapath.DevicesControllerCell,
	cell.Provide(func(cfg *option.DaemonConfig) linuxdatapath.DevicesConfig {
		// Provide the configured devices to the devices controller.
		// This is temporary until DevicesController takes ownership of the
		// device-related configuration options.
		return linuxdatapath.DevicesConfig{Devices: cfg.GetDevices()}
	}),
)

func newWireguardAgent(lc hive.Lifecycle, localNodeStore *node.LocalNodeStore) *wg.Agent {
	var wgAgent *wg.Agent
	if option.Config.EnableWireguard {
		if option.Config.EnableIPSec {
			log.Fatalf("WireGuard (--%s) cannot be used with IPsec (--%s)",
				option.EnableWireguard, option.EnableIPSecName)
		}

		var err error
		privateKeyPath := filepath.Join(option.Config.StateDir, wgTypes.PrivKeyFilename)
		wgAgent, err = wg.NewAgent(privateKeyPath, localNodeStore)
		if err != nil {
			log.Fatalf("failed to initialize WireGuard: %s", err)
		}

		lc.Append(hive.Hook{
			OnStop: func(hive.HookContext) error {
				wgAgent.Close()
				return nil
			},
		})
	} else {
		// Delete WireGuard device from previous run (if such exists)
		link.DeleteByName(wgTypes.IfaceName)
	}
	return wgAgent
}

func newDatapath(params datapathParams) types.Datapath {
	datapathConfig := linuxdatapath.DatapathConfiguration{
		HostDevice: defaults.HostDevice,
		ProcFs:     option.Config.ProcFs,
	}

	iptablesManager := &iptables.IptablesManager{}

	params.LC.Append(hive.Hook{
		OnStart: func(hive.HookContext) error {
			// FIXME enableIPForwarding should not live here
			if err := enableIPForwarding(); err != nil {
				log.Fatalf("enabling IP forwarding via sysctl failed: %s", err)
			}

			iptablesManager.Init()
			return nil
		}})

	datapath := linuxdatapath.NewDatapath(datapathConfig, iptablesManager, params.WgAgent, params.NodeMap, params.ConfigWriter)

	params.LC.Append(hive.Hook{
		OnStart: func(hive.HookContext) error {
			datapath.NodeIDs().RestoreNodeIDs()
			return nil
		},
	})

	return datapath
}

type datapathParams struct {
	cell.In

	LC      hive.Lifecycle
	WgAgent *wg.Agent

	// Force map initialisation before loader. You should not use these otherwise.
	// Some of the entries in this slice may be nil.
	BpfMaps []bpf.BpfMap `group:"bpf-maps"`

	NodeMap nodemap.Map

	// Depend on DeviceManager to ensure devices have been resolved.
	// This is required until option.Config.GetDevices() has been removed and
	// uses of it converted to Table[Device].
	DeviceManager *linuxdatapath.DeviceManager

	ConfigWriter types.ConfigWriter
}
