// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package predicate

import (
	"github.com/insolar/assured-ledger/ledger-core/testutils/debuglogger"
)

type Predicate = func(debuglogger.UpdateEvent) bool

func Never() Predicate {
	return func(debuglogger.UpdateEvent) bool {
		return false
	}
}

func Ever() Predicate {
	return func(debuglogger.UpdateEvent) bool {
		return true
	}
}

func Not(predicate Predicate) Predicate {
	return func(event debuglogger.UpdateEvent) bool {
		return !predicate(event)
	}
}

func And(predicates ...Predicate) Predicate {
	if len(predicates) == 0 {
		return Never()
	}
	return func(event debuglogger.UpdateEvent) bool {
		for _, fn := range predicates {
			if !fn(event) {
				return false
			}
		}
		return true
	}
}

func Or(predicates ...Predicate) Predicate {
	return func(event debuglogger.UpdateEvent) bool {
		for _, fn := range predicates {
			if fn(event) {
				return true
			}
		}
		return false
	}
}
