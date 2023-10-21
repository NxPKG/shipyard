// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package proxy

import (
	"github.com/khulnasoft/shipyard/pkg/completion"
	"github.com/khulnasoft/shipyard/pkg/revert"
)

// Redirect type for custom Listeners, which are managed externally.
type CRDRedirect struct{}

func (r *CRDRedirect) UpdateRules(wg *completion.WaitGroup) (revert.RevertFunc, error) {
	return func() error { return nil }, nil
}

func (r *CRDRedirect) Close(wg *completion.WaitGroup) (revert.FinalizeFunc, revert.RevertFunc) {
	return nil, func() error { return nil }
}
