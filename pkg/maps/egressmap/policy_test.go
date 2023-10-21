// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package egressmap

import (
	"errors"
	"net/netip"
	"testing"

	"github.com/khulnasoft/gbpf/rlimit"
	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/shipyard/pkg/bpf"
	"github.com/khulnasoft/shipyard/pkg/ebpf"
	"github.com/khulnasoft/shipyard/pkg/hive/hivetest"
	"github.com/khulnasoft/shipyard/pkg/testutils"
)

func TestPolicyMap(t *testing.T) {
	testutils.PrivilegedTest(t)

	bpf.CheckOrMountFS("")
	assert.Nil(t, rlimit.RemoveMemlock())

	egressPolicyMap := createPolicyMap(hivetest.Lifecycle(t), DefaultPolicyConfig, ebpf.PinNone)

	sourceIP1 := netip.MustParseAddr("1.1.1.1")
	sourceIP2 := netip.MustParseAddr("1.1.1.2")

	destCIDR1 := netip.MustParsePrefix("2.2.1.0/24")
	destCIDR2 := netip.MustParsePrefix("2.2.2.0/24")

	egressIP1 := netip.MustParseAddr("3.3.3.1")
	egressIP2 := netip.MustParseAddr("3.3.3.2")

	err := egressPolicyMap.Update(sourceIP1, destCIDR1, egressIP1, egressIP1)
	assert.Nil(t, err)

	err = egressPolicyMap.Update(sourceIP2, destCIDR2, egressIP2, egressIP2)
	assert.Nil(t, err)

	val, err := egressPolicyMap.Lookup(sourceIP1, destCIDR1)
	assert.Nil(t, err)

	assert.Equal(t, val.EgressIP.Addr(), egressIP1)
	assert.Equal(t, val.GatewayIP.Addr(), egressIP1)

	val, err = egressPolicyMap.Lookup(sourceIP2, destCIDR2)
	assert.Nil(t, err)

	assert.Equal(t, val.EgressIP.Addr(), egressIP2)
	assert.Equal(t, val.GatewayIP.Addr(), egressIP2)

	err = egressPolicyMap.Delete(sourceIP2, destCIDR2)
	assert.Nil(t, err)

	val, err = egressPolicyMap.Lookup(sourceIP1, destCIDR1)
	assert.Nil(t, err)

	assert.Equal(t, val.EgressIP.Addr(), egressIP1)
	assert.Equal(t, val.GatewayIP.Addr(), egressIP1)

	_, err = egressPolicyMap.Lookup(sourceIP2, destCIDR2)
	assert.True(t, errors.Is(err, ebpf.ErrKeyNotExist))
}
