// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package routing

import (
	"context"

	"github.com/insolar/assured-ledger/ledger-core/network"
	"github.com/insolar/assured-ledger/ledger-core/network/hostnetwork/host"
	"github.com/insolar/assured-ledger/ledger-core/network/storage"
	"github.com/insolar/assured-ledger/ledger-core/reference"

	errors "github.com/insolar/assured-ledger/ledger-core/vanilla/throw"
)

type Table struct {
	NodeKeeper    network.NodeKeeper    `inject:""`
	PulseAccessor storage.PulseAccessor `inject:""`
}

func (t *Table) isLocalNode(reference.Global) bool {
	return true
}

func (t *Table) resolveRemoteNode(reference.Global) (*host.Host, error) {
	return nil, errors.New("not implemented")
}

// Resolve NodeID -> ShortID, Address. Can initiate network requests.
func (t *Table) Resolve(ref reference.Global) (*host.Host, error) {
	if t.isLocalNode(ref) {
		p, err := t.PulseAccessor.GetLatestPulse(context.Background())
		if err != nil {
			return nil, errors.W(err, "failed to get latest pulse --==-- ")
		}

		node := t.NodeKeeper.GetAccessor(p.PulseNumber).GetActiveNode(ref)
		if node == nil {
			return nil, errors.New("no such local node with NodeID: " + ref.String())
		}
		return host.NewHostNS(node.Address(), node.ID(), node.ShortID())
	}
	return t.resolveRemoteNode(ref)
}
