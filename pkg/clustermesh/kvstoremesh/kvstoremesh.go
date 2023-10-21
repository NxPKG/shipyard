// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package kvstoremesh

import (
	"context"

	"github.com/khulnasoft/shipyard/pkg/clustermesh/common"
	"github.com/khulnasoft/shipyard/pkg/clustermesh/types"
	"github.com/khulnasoft/shipyard/pkg/hive"
	"github.com/khulnasoft/shipyard/pkg/hive/cell"
	identityCache "github.com/khulnasoft/shipyard/pkg/identity/cache"
	"github.com/khulnasoft/shipyard/pkg/ipcache"
	"github.com/khulnasoft/shipyard/pkg/kvstore"
	"github.com/khulnasoft/shipyard/pkg/kvstore/store"
	nodeStore "github.com/khulnasoft/shipyard/pkg/node/store"
	"github.com/khulnasoft/shipyard/pkg/promise"
	serviceStore "github.com/khulnasoft/shipyard/pkg/service/store"
)

// KVStoreMesh is a cache of multiple remote clusters
type KVStoreMesh struct {
	common common.ClusterMesh

	// backend is the interface to operate the local kvstore
	backend        kvstore.BackendOperations
	backendPromise promise.Promise[kvstore.BackendOperations]

	storeFactory store.Factory
}

type params struct {
	cell.In

	ClusterInfo types.ClusterInfo
	common.Config

	BackendPromise promise.Promise[kvstore.BackendOperations]

	Metrics      common.Metrics
	StoreFactory store.Factory
}

func newKVStoreMesh(lc hive.Lifecycle, params params) *KVStoreMesh {
	km := KVStoreMesh{
		backendPromise: params.BackendPromise,
		storeFactory:   params.StoreFactory,
	}
	km.common = common.NewClusterMesh(common.Configuration{
		Config:           params.Config,
		ClusterInfo:      params.ClusterInfo,
		NewRemoteCluster: km.newRemoteCluster,
		Metrics:          params.Metrics,
	})

	lc.Append(&km)

	// The "common" Start hook needs to be executed after that the kvstoremesh one
	// terminated, to ensure that the backend promise has already been resolved.
	lc.Append(&km.common)

	return &km
}

func (km *KVStoreMesh) Start(ctx hive.HookContext) error {
	backend, err := km.backendPromise.Await(ctx)
	if err != nil {
		return err
	}

	km.backend = backend
	return nil
}

func (km *KVStoreMesh) Stop(hive.HookContext) error {
	return nil
}

func (km *KVStoreMesh) newRemoteCluster(name string, _ common.StatusFunc) common.RemoteCluster {
	ctx, cancel := context.WithCancel(context.Background())

	rc := &remoteCluster{
		name:         name,
		localBackend: km.backend,

		cancel: cancel,

		nodes:        newReflector(km.backend, name, nodeStore.NodeStorePrefix, km.storeFactory),
		services:     newReflector(km.backend, name, serviceStore.ServiceStorePrefix, km.storeFactory),
		identities:   newReflector(km.backend, name, identityCache.IdentitiesPath, km.storeFactory),
		ipcache:      newReflector(km.backend, name, ipcache.IPIdentitiesPath, km.storeFactory),
		storeFactory: km.storeFactory,
	}

	run := func(fn func(context.Context)) {
		rc.wg.Add(1)
		go func() {
			fn(ctx)
			rc.wg.Done()
		}()
	}

	run(rc.nodes.syncer.Run)
	run(rc.services.syncer.Run)
	run(rc.identities.syncer.Run)
	run(rc.ipcache.syncer.Run)

	return rc
}
