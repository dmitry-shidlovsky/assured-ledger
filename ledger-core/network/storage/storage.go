// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package storage

import (
	errors "github.com/insolar/assured-ledger/ledger-core/vanilla/throw"
)

var (
	// ErrNotFound is returned when value was not found.
	ErrNotFound = errors.New("value not found")
)
