// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package execute

import (
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/insolar/assured-ledger/ledger-core/v2/conveyor/smachine"
	"github.com/insolar/assured-ledger/ledger-core/v2/conveyor/smachine/smsync"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/contract"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/payload"
	"github.com/insolar/assured-ledger/ledger-core/v2/instrumentation/inslogger"
	"github.com/insolar/assured-ledger/ledger-core/v2/reference"
	"github.com/insolar/assured-ledger/ledger-core/v2/testutils/gen"
	"github.com/insolar/assured-ledger/ledger-core/v2/virtual/object"
	"github.com/insolar/assured-ledger/ledger-core/v2/virtual/testutils/slotdebugger"
)

func TestSMExecute_Semi_IncrementPendingCounters(t *testing.T) {
	var (
		mc  = minimock.NewController(t)
		ctx = inslogger.TestContext(t)

		class       = gen.UniqueReference()
		caller      = gen.UniqueReference()
		callee      = gen.UniqueReference()
		outgoing    = gen.UniqueID()
		objectRef   = reference.NewSelf(outgoing)
		sharedState = &object.SharedState{
			Info: object.Info{
				PendingTable:   object.NewRequestTable(),
				KnownRequests:  object.NewRequestTable(),
				ReadyToWork:    smsync.NewConditional(1, "ReadyToWork").SyncLink(),
				OrderedExecute: smsync.NewConditional(1, "MutableExecution").SyncLink(),
			},
		}
	)

	slotMachine := slotdebugger.New(ctx, t, true)
	slotMachine.InitEmptyMessageSender(mc)
	slotMachine.PrepareRunner(ctx, mc)

	smExecute := SMExecute{
		Payload: &payload.VCallRequest{
			CallType:     payload.CTConstructor,
			CallFlags:    payload.BuildCallFlags(contract.CallTolerable, contract.CallDirty),
			CallOutgoing: outgoing,

			Caller:              caller,
			Callee:              callee,
			CallSiteDeclaration: class,
			CallSiteMethod:      "New",
		},
		Meta: &payload.Meta{
			Sender: caller,
		},
	}
	catalogWrapper := object.NewCatalogMockWrapper(mc)

	{
		var catalog object.Catalog = catalogWrapper.Mock()
		slotMachine.AddInterfaceDependency(&catalog)

		sharedStateData := smachine.NewUnboundSharedData(sharedState)
		smObjectAccessor := object.SharedStateAccessor{SharedDataLink: sharedStateData}

		catalogWrapper.AddObject(objectRef, smObjectAccessor)
		catalogWrapper.AllowAccessMode(object.CatalogMockAccessGetOrCreate)
	}

	slotMachine.Start()
	defer slotMachine.Stop()

	smWrapper := slotMachine.AddStateMachine(ctx, &smExecute)

	require.Equal(t, uint8(0), sharedState.PotentialOrderedPendingCount)
	require.Equal(t, uint8(0), sharedState.PotentialUnorderedPendingCount)

	slotMachine.RunTil(smWrapper.BeforeStep(smExecute.stepExecuteStart))

	require.Equal(t, uint8(1), sharedState.PotentialOrderedPendingCount)
	require.Equal(t, uint8(0), sharedState.PotentialUnorderedPendingCount)

	require.NoError(t, catalogWrapper.CheckDone())
	mc.Finish()
}

func TestSMExecute_MigrateBeforeLock(t *testing.T) {
	var (
		mc  = minimock.NewController(t)
		ctx = inslogger.TestContext(t)

		class       = gen.UniqueReference()
		caller      = gen.UniqueReference()
		callee      = gen.UniqueReference()
		outgoing    = gen.UniqueID()
		objectRef   = reference.NewSelf(outgoing)
		sharedState = &object.SharedState{
			Info: object.Info{
				PendingTable:   object.NewRequestTable(),
				KnownRequests:  object.NewRequestTable(),
				ReadyToWork:    smsync.NewConditional(1, "ReadyToWork").SyncLink(),
				OrderedExecute: smsync.NewConditional(1, "MutableExecution").SyncLink(),
			},
		}
	)

	slotMachine := slotdebugger.New(ctx, t, true)
	slotMachine.InitEmptyMessageSender(mc)
	slotMachine.PrepareRunner(ctx, mc)

	smExecute := SMExecute{
		Payload: &payload.VCallRequest{
			CallType:     payload.CTConstructor,
			CallFlags:    payload.BuildCallFlags(contract.CallTolerable, contract.CallDirty),
			CallOutgoing: outgoing,

			Caller:              caller,
			Callee:              callee,
			CallSiteDeclaration: class,
			CallSiteMethod:      "New",
		},
		Meta: &payload.Meta{
			Sender: caller,
		},
	}
	catalogWrapper := object.NewCatalogMockWrapper(mc)

	{
		var catalog object.Catalog = catalogWrapper.Mock()
		slotMachine.AddInterfaceDependency(&catalog)

		sharedStateData := smachine.NewUnboundSharedData(sharedState)
		smObjectAccessor := object.SharedStateAccessor{SharedDataLink: sharedStateData}

		catalogWrapper.AddObject(objectRef, smObjectAccessor)
		catalogWrapper.AllowAccessMode(object.CatalogMockAccessGetOrCreate)
	}

	slotMachine.Start()
	defer slotMachine.Stop()

	smWrapper := slotMachine.AddStateMachine(ctx, &smExecute)

	require.False(t, smExecute.migrationHappened)

	slotMachine.RunTil(smWrapper.BeforeStep(smExecute.stepTakeLock))

	slotMachine.Migrate()

	slotMachine.RunTil(smWrapper.AfterStop())

	require.False(t, smExecute.migrationHappened)

	require.NoError(t, catalogWrapper.CheckDone())
	mc.Finish()
}

func TestSMExecute_MigrateAfterLock(t *testing.T) {
	var (
		mc  = minimock.NewController(t)
		ctx = inslogger.TestContext(t)

		class       = gen.UniqueReference()
		caller      = gen.UniqueReference()
		callee      = gen.UniqueReference()
		outgoing    = gen.UniqueID()
		objectRef   = reference.NewSelf(outgoing)
		sharedState = &object.SharedState{
			Info: object.Info{
				PendingTable:   object.NewRequestTable(),
				KnownRequests:  object.NewRequestTable(),
				ReadyToWork:    smsync.NewConditional(1, "ReadyToWork").SyncLink(),
				OrderedExecute: smsync.NewConditional(1, "MutableExecution").SyncLink(),
			},
		}
	)

	slotMachine := slotdebugger.New(ctx, t, true)
	slotMachine.InitEmptyMessageSender(mc)
	slotMachine.PrepareRunner(ctx, mc)

	smExecute := SMExecute{
		Payload: &payload.VCallRequest{
			CallType:     payload.CTConstructor,
			CallFlags:    payload.BuildCallFlags(contract.CallTolerable, contract.CallDirty),
			CallOutgoing: outgoing,

			Caller:              caller,
			Callee:              callee,
			CallSiteDeclaration: class,
			CallSiteMethod:      "New",
		},
		Meta: &payload.Meta{
			Sender: caller,
		},
	}
	catalogWrapper := object.NewCatalogMockWrapper(mc)

	{
		var catalog object.Catalog = catalogWrapper.Mock()
		slotMachine.AddInterfaceDependency(&catalog)

		sharedStateData := smachine.NewUnboundSharedData(sharedState)
		smObjectAccessor := object.SharedStateAccessor{SharedDataLink: sharedStateData}

		catalogWrapper.AddObject(objectRef, smObjectAccessor)
		catalogWrapper.AllowAccessMode(object.CatalogMockAccessGetOrCreate)
	}

	slotMachine.Start()
	defer slotMachine.Stop()

	smWrapper := slotMachine.AddStateMachine(ctx, &smExecute)

	require.False(t, smExecute.migrationHappened)

	slotMachine.RunTil(smWrapper.BeforeStep(smExecute.stepExecuteStart))

	slotMachine.Migrate()

	slotMachine.RunTil(smWrapper.AfterAnyMigrate())

	assert.True(t, smExecute.migrationHappened)

	require.NoError(t, catalogWrapper.CheckDone())
	mc.Finish()
}

func TestSMExecute_Semi_ConstructorOnMissingObject(t *testing.T) {
	var (
		mc  = minimock.NewController(t)
		ctx = inslogger.TestContext(t)

		class       = gen.UniqueReference()
		caller      = gen.UniqueReference()
		callee      = gen.UniqueReference()
		outgoing    = gen.UniqueID()
		objectRef   = reference.NewSelf(outgoing)
		sharedState = &object.SharedState{
			Info: object.Info{
				PendingTable:   object.NewRequestTable(),
				KnownRequests:  object.NewRequestTable(),
				ReadyToWork:    smsync.NewConditional(1, "ReadyToWork").SyncLink(),
				OrderedExecute: smsync.NewConditional(1, "MutableExecution").SyncLink(),
			},
		}
	)

	sharedState.SetState(object.Missing)

	slotMachine := slotdebugger.New(ctx, t, true)
	slotMachine.InitEmptyMessageSender(mc)
	slotMachine.PrepareRunner(ctx, mc)

	smExecute := SMExecute{
		Payload: &payload.VCallRequest{
			CallType:     payload.CTConstructor,
			CallFlags:    payload.BuildCallFlags(contract.CallTolerable, contract.CallDirty),
			CallOutgoing: outgoing,

			Caller:              caller,
			Callee:              callee,
			CallSiteDeclaration: class,
			CallSiteMethod:      "New",
		},
		Meta: &payload.Meta{
			Sender: caller,
		},
	}
	catalogWrapper := object.NewCatalogMockWrapper(mc)

	{
		var catalog object.Catalog = catalogWrapper.Mock()
		slotMachine.AddInterfaceDependency(&catalog)

		sharedStateData := smachine.NewUnboundSharedData(sharedState)
		smObjectAccessor := object.SharedStateAccessor{SharedDataLink: sharedStateData}

		catalogWrapper.AddObject(objectRef, smObjectAccessor)
		catalogWrapper.AllowAccessMode(object.CatalogMockAccessGetOrCreate)
	}

	slotMachine.Start()
	defer slotMachine.Stop()

	smWrapper := slotMachine.AddStateMachine(ctx, &smExecute)

	require.Equal(t, uint8(0), sharedState.PotentialOrderedPendingCount)
	require.Equal(t, uint8(0), sharedState.PotentialUnorderedPendingCount)

	slotMachine.RunTil(smWrapper.BeforeStep(smExecute.stepExecuteStart))

	require.Equal(t, uint8(1), sharedState.PotentialOrderedPendingCount)
	require.Equal(t, uint8(0), sharedState.PotentialUnorderedPendingCount)

	require.NoError(t, catalogWrapper.CheckDone())
	mc.Finish()
}

func TestSMExecute_Semi_ConstructorOnBadObject(t *testing.T) {
	var (
		mc  = minimock.NewController(t)
		ctx = inslogger.TestContext(t)

		class       = gen.UniqueReference()
		caller      = gen.UniqueReference()
		callee      = gen.UniqueReference()
		outgoing    = gen.UniqueID()
		objectRef   = reference.NewSelf(outgoing)
		sharedState = &object.SharedState{
			Info: object.Info{
				PendingTable:   object.NewRequestTable(),
				KnownRequests:  object.NewRequestTable(),
				ReadyToWork:    smsync.NewConditional(1, "ReadyToWork").SyncLink(),
				OrderedExecute: smsync.NewConditional(1, "MutableExecution").SyncLink(),
			},
		}
	)

	sharedState.SetState(object.Inactive)

	slotMachine := slotdebugger.New(ctx, t, true)
	slotMachine.InitEmptyMessageSender(mc)
	slotMachine.PrepareRunner(ctx, mc)

	smExecute := SMExecute{
		Payload: &payload.VCallRequest{
			CallType:     payload.CTConstructor,
			CallFlags:    payload.BuildCallFlags(contract.CallTolerable, contract.CallDirty),
			CallOutgoing: outgoing,

			Caller:              caller,
			Callee:              callee,
			CallSiteDeclaration: class,
			CallSiteMethod:      "New",
		},
		Meta: &payload.Meta{
			Sender: caller,
		},
	}
	catalogWrapper := object.NewCatalogMockWrapper(mc)

	{
		var catalog object.Catalog = catalogWrapper.Mock()
		slotMachine.AddInterfaceDependency(&catalog)

		sharedStateData := smachine.NewUnboundSharedData(sharedState)
		smObjectAccessor := object.SharedStateAccessor{SharedDataLink: sharedStateData}

		catalogWrapper.AddObject(objectRef, smObjectAccessor)
		catalogWrapper.AllowAccessMode(object.CatalogMockAccessGetOrCreate)
	}

	slotMachine.Start()
	defer slotMachine.Stop()

	smWrapper := slotMachine.AddStateMachine(ctx, &smExecute)

	require.Equal(t, uint8(0), sharedState.PotentialOrderedPendingCount)
	require.Equal(t, uint8(0), sharedState.PotentialUnorderedPendingCount)

	slotMachine.RunTil(smWrapper.AfterStop())

	require.Equal(t, uint8(0), sharedState.PotentialOrderedPendingCount)
	require.Equal(t, uint8(0), sharedState.PotentialUnorderedPendingCount)

	require.NoError(t, catalogWrapper.CheckDone())
	mc.Finish()
}

func TestSMExecute_Semi_MethodOnEmptyObject(t *testing.T) {
	var (
		mc  = minimock.NewController(t)
		ctx = inslogger.TestContext(t)

		class       = gen.UniqueReference()
		caller      = gen.UniqueReference()
		outgoing    = gen.UniqueID()
		objectRef   = reference.NewSelf(outgoing)
		sharedState = &object.SharedState{
			Info: object.Info{
				PendingTable:   object.NewRequestTable(),
				KnownRequests:  object.NewRequestTable(),
				ReadyToWork:    smsync.NewConditional(1, "ReadyToWork").SyncLink(),
				OrderedExecute: smsync.NewConditional(1, "MutableExecution").SyncLink(),
			},
		}
	)

	sharedState.SetState(object.Empty)

	slotMachine := slotdebugger.New(ctx, t, true)
	slotMachine.InitEmptyMessageSender(mc)
	slotMachine.PrepareRunner(ctx, mc)

	smExecute := SMExecute{
		Payload: &payload.VCallRequest{
			CallType:     payload.CTMethod,
			CallFlags:    payload.BuildCallFlags(contract.CallTolerable, contract.CallDirty),
			CallOutgoing: outgoing,

			Caller:              caller,
			Callee:              objectRef,
			CallSiteDeclaration: class,
			CallSiteMethod:      "New",
		},
		Meta: &payload.Meta{
			Sender: caller,
		},
	}
	catalogWrapper := object.NewCatalogMockWrapper(mc)

	{
		var catalog object.Catalog = catalogWrapper.Mock()
		slotMachine.AddInterfaceDependency(&catalog)

		sharedStateData := smachine.NewUnboundSharedData(sharedState)
		smObjectAccessor := object.SharedStateAccessor{SharedDataLink: sharedStateData}

		catalogWrapper.AddObject(objectRef, smObjectAccessor)
		catalogWrapper.AllowAccessMode(object.CatalogMockAccessGetOrCreate)
	}

	slotMachine.Start()
	defer slotMachine.Stop()

	smWrapper := slotMachine.AddStateMachine(ctx, &smExecute)

	require.Equal(t, uint8(0), sharedState.PotentialOrderedPendingCount)
	require.Equal(t, uint8(0), sharedState.PotentialUnorderedPendingCount)

	slotMachine.RunTil(smWrapper.AfterStop())

	require.Equal(t, uint8(0), sharedState.PotentialOrderedPendingCount)
	require.Equal(t, uint8(0), sharedState.PotentialUnorderedPendingCount)

	require.NoError(t, catalogWrapper.CheckDone())
	mc.Finish()
}
