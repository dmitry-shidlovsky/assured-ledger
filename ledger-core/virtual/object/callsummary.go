// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

//go:generate sm-uml-gen -f $GOFILE

package object

import (
	"github.com/insolar/assured-ledger/ledger-core/conveyor/smachine"
	"github.com/insolar/assured-ledger/ledger-core/pulse"
	"github.com/insolar/assured-ledger/ledger-core/reference"
	"github.com/insolar/assured-ledger/ledger-core/vanilla/injector"
)

type PulseNumber = pulse.Number

type SummarySyncKey struct {
	objectRef   reference.Global
	pulseNumber pulse.Number
}

type SummarySharedKey struct {
	PulseNumber
}

type SMCallSummary struct {
	smachine.StateMachineDeclTemplate

	pulse  pulse.Number
	shared SharedCallSummary
}

type SharedCallSummary struct {
	Requests ObjectsRequestsTable
}

func (sm *SMCallSummary) InjectDependencies(_ smachine.StateMachine, _ smachine.SlotLink, _ *injector.DependencyInjector) {
}

func (sm *SMCallSummary) GetInitStateFor(_ smachine.StateMachine) smachine.InitFunc {
	return sm.Init
}

func (sm *SMCallSummary) GetStateMachineDeclaration() smachine.StateMachineDeclaration {
	return sm
}

func (sm *SMCallSummary) Init(ctx smachine.InitializationContext) smachine.StateUpdate {
	sm.shared = SharedCallSummary{Requests: NewObjectRequestTable()}

	sdl := ctx.Share(&sm.shared, 0)
	if !ctx.Publish(SummarySharedKey{PulseNumber: sm.pulse}, sdl) {
		return ctx.Stop()
	}

	ctx.SetDefaultMigration(sm.stepMigrate)

	return ctx.Jump(sm.stepLoop)
}

func (sm *SMCallSummary) stepMigrate(ctx smachine.MigrationContext) smachine.StateUpdate {
	return ctx.Stop()
}

func (sm *SMCallSummary) stepLoop(ctx smachine.ExecutionContext) smachine.StateUpdate {
	return ctx.Sleep().ThenRepeat()
}

type CallSummarySharedStateAccessor struct {
	smachine.SharedDataLink
}

func (v CallSummarySharedStateAccessor) Prepare(fn func(*SharedCallSummary)) smachine.SharedDataAccessor {
	return v.PrepareAccess(func(data interface{}) bool {
		fn(data.(*SharedCallSummary))
		return false
	})
}

func (v CallSummarySharedStateAccessor) PrepareAndWakeUp(fn func(*SharedCallSummary)) smachine.SharedDataAccessor {
	return v.PrepareAccess(func(data interface{}) bool {
		fn(data.(*SharedCallSummary))
		return true
	})
}

func GetSummarySMSharedAccessor(
	ctx smachine.ExecutionContext,
	summaryKey SummarySharedKey,
) (CallSummarySharedStateAccessor, bool) {
	if v := ctx.GetPublishedLink(summaryKey); v.IsAssignableTo((*SharedCallSummary)(nil)) {
		return CallSummarySharedStateAccessor{v}, true
	}
	return CallSummarySharedStateAccessor{}, false
}
