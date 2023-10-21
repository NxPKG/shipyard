// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package cmd

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/khulnasoft/shipyard/api/v1/models"
	. "github.com/khulnasoft/shipyard/api/v1/server/restapi/service"
	"github.com/khulnasoft/shipyard/pkg/logging/logfields"
	"github.com/khulnasoft/shipyard/pkg/redirectpolicy"
)

func getLRPHandler(d *Daemon, params GetLrpParams) middleware.Responder {
	log.WithField(logfields.Params, logfields.Repr(params)).Debug("GET /lrp request")
	return NewGetLrpOK().WithPayload(getLRPs(d.redirectPolicyManager))
}

func getLRPs(rpm *redirectpolicy.Manager) []*models.LRPSpec {
	lrps := rpm.GetLRPs()
	list := make([]*models.LRPSpec, 0, len(lrps))
	for _, v := range lrps {
		list = append(list, v.GetModel())
	}
	return list
}
