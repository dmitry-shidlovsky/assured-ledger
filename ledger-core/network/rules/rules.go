// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package rules

import (
	"fmt"

	errors "github.com/insolar/assured-ledger/ledger-core/vanilla/throw"

	"github.com/insolar/assured-ledger/ledger-core/insolar/node"
	"github.com/insolar/assured-ledger/ledger-core/log/global"
	"github.com/insolar/assured-ledger/ledger-core/network"
)

// CheckMajorityRule returns error if MajorityRule check not passed, also returns active discovery nodes count
func CheckMajorityRule(cert node.Certificate, nodes []node.NetworkNode) (int, error) {
	majorityRule := cert.GetMajorityRule()
	discoveriesInList := network.FindDiscoveriesInNodeList(nodes, cert)
	activeDiscoveryNodesLen := len(discoveriesInList)
	if activeDiscoveryNodesLen >= majorityRule {
		return activeDiscoveryNodesLen, nil
	}
	strErr := fmt.Sprintf("Active discovery nodes len actual %d, expected %d. Not active ",
		activeDiscoveryNodesLen, majorityRule)
	discoveries := cert.GetDiscoveryNodes()
	for _, d := range discoveries {
		var found bool
		for _, n := range nodes {
			if d.GetNodeRef().Equal(n.ID()) {
				found = true
				break
			}
		}
		if !found {
			strErr += d.GetHost() + " " + d.GetRole().String() + " "
		}
	}
	return activeDiscoveryNodesLen, errors.W(errors.New(strErr), "MajorityRule failed")
}

// CheckMinRole returns true if MinRole check passed
func CheckMinRole(cert node.Certificate, nodes []node.NetworkNode) error {
	var virtualCount, heavyCount, lightCount uint
	for _, n := range nodes {
		switch n.Role() {
		case node.StaticRoleVirtual:
			virtualCount++
		case node.StaticRoleHeavyMaterial:
			heavyCount++
		case node.StaticRoleLightMaterial:
			lightCount++
		default:
			global.Warn("unknown node role")
		}
	}

	v, h, l := cert.GetMinRoles()
	if virtualCount >= v &&
		heavyCount >= h &&
		lightCount >= l {
		return nil
	}

	err := errors.New(fmt.Sprintf("%s actual %d expected %d, %s actual %d expected %d, %s actual %d expected %d",
		node.StaticRoleVirtual.String(), virtualCount, v,
		node.StaticRoleHeavyMaterial.String(), heavyCount, h,
		node.StaticRoleLightMaterial.String(), lightCount, l))
	return errors.W(err, "MinRoles failed")
}
