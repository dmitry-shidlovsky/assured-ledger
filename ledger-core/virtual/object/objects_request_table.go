// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package object

import (
	"github.com/insolar/assured-ledger/ledger-core/reference"
)

type ObjectsRequestsTable struct {
	knownRequests map[reference.Global]*WorkingTable
	pendingTable  map[reference.Global]*PendingTable
}

func NewObjectRequestTable() ObjectsRequestsTable {
	return ObjectsRequestsTable{
		knownRequests: make(map[reference.Global]*WorkingTable),
		pendingTable:  make(map[reference.Global]*PendingTable),
	}
}

func (ort *ObjectsRequestsTable) GetObjectsKnownRequests(ref reference.Global) (*WorkingTable, bool) {
	workingTable, ok := ort.knownRequests[ref]
	return workingTable, ok
}

func (ort *ObjectsRequestsTable) GetObjectsPendingRequests(ref reference.Global) (*PendingTable, bool) {
	pendingTable, ok := ort.pendingTable[ref]
	return pendingTable, ok
}

func (ort *ObjectsRequestsTable) AddObjectRequests(
	ref reference.Global,
	knownRequests WorkingTable,
	pendingRequests PendingTable,
) {
	ort.knownRequests[ref] = &knownRequests
	ort.pendingTable[ref] = &pendingRequests
}
