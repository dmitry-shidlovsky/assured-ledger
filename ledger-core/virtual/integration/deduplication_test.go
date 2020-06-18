// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package integration

import (
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/insolar/assured-ledger/ledger-core/insolar/contract"
	"github.com/insolar/assured-ledger/ledger-core/insolar/payload"
	"github.com/insolar/assured-ledger/ledger-core/runner/execution"
	"github.com/insolar/assured-ledger/ledger-core/runner/executionupdate"
	"github.com/insolar/assured-ledger/ledger-core/runner/requestresult"
	"github.com/insolar/assured-ledger/ledger-core/testutils/gen"
	"github.com/insolar/assured-ledger/ledger-core/testutils/runner/logicless"
	"github.com/insolar/assured-ledger/ledger-core/virtual/execute"
	"github.com/insolar/assured-ledger/ledger-core/virtual/integration/utils"
)

type SynchronizationPoint struct {
	count int

	input  chan struct{}
	output chan struct{}
}

func (p *SynchronizationPoint) Synchronize() {
	p.input <- struct{}{}

	<-p.output
}

func (p *SynchronizationPoint) Wait(t *testing.T) {
	for i := 0; i < p.count; i++ {
		select {
		case <-p.input:
		case <-time.After(10 * time.Second):
			t.Fatal("timeout: failed to wait until all goroutines are synced")
		}
	}
}

func (p *SynchronizationPoint) WakeUp() {
	for i := 0; i < p.count; i++ {
		p.output <- struct{}{}
	}

}

func NewSynchronizationPoint(count int) *SynchronizationPoint {
	return &SynchronizationPoint{
		count: count,

		input:  make(chan struct{}, count),
		output: make(chan struct{}, 0),
	}
}

func TestDeduplication_Constructor_DuringExecution_DifferentPulses(t *testing.T) {
	t.Log("C4998")

	var (
		mc = minimock.NewController(t)
	)

	server, ctx := utils.NewUninitializedServer(nil, t)
	defer server.Stop()

	executeDone := server.Journal.WaitStopOf(&execute.SMExecute{}, 2)

	runnerMock := logicless.NewServiceMock(ctx, t, nil)
	server.ReplaceRunner(runnerMock)
	server.Init(ctx)

	var (
		isolation = contract.ConstructorIsolation()
		outgoing  = server.RandomLocalWithPulse()
		class     = gen.UniqueReference()
	)

	pl := payload.VCallRequest{
		CallType:       payload.CTConstructor,
		CallFlags:      payload.BuildCallFlags(isolation.Interference, isolation.State),
		Callee:         class,
		CallSiteMethod: "New",
		CallOutgoing:   outgoing,
	}

	synchronizeExecution := NewSynchronizationPoint(1)

	{
		requestResult := requestresult.New([]byte("123"), gen.UniqueReference())
		requestResult.SetActivate(gen.UniqueReference(), class, []byte("234"))

		executionMock := runnerMock.AddExecutionMock(calculateOutgoing(pl).String())
		executionMock.AddStart(func(ctx execution.Context) {
			synchronizeExecution.Synchronize()
		}, &executionupdate.ContractExecutionStateUpdate{
			Type:   executionupdate.Done,
			Result: requestResult,
		})
	}

	typedChecker := server.PublisherMock.SetTypedChecker(ctx, mc, server)
	typedChecker.VCallResult.SetResend(false)
	typedChecker.VDelegatedCallRequest.SetResend(true)
	typedChecker.VDelegatedCallResponse.SetResend(true)
	typedChecker.VDelegatedRequestFinished.SetResend(true)
	typedChecker.VStateReport.SetResend(true)

	{
		msg := server.WrapPayload(&pl).Finalize()
		server.SendMessage(ctx, msg)
	}

	synchronizeExecution.Wait(t)
	server.IncrementPulse(ctx)
	{
		server.SuspendConveyorAndWaitThenResetActive()
		server.SendMessage(ctx, server.WrapPayload(&pl).Finalize())
		server.WaitActiveThenIdleConveyor()
	}
	synchronizeExecution.WakeUp()

	{
		select {
		case <-executeDone:
		case <-time.After(10 * time.Second):
			require.FailNow(t, "timeout")
		}

		select {
		case <-server.Journal.WaitAllAsyncCallsDone():
		case <-time.After(10 * time.Second):
			require.FailNow(t, "timeout")
		}
	}

	{
		assert.Equal(t, 1, typedChecker.VCallResult.Count())
		assert.Equal(t, 1, typedChecker.VDelegatedCallRequest.Count())
		assert.Equal(t, 1, typedChecker.VDelegatedCallResponse.Count())
		assert.Equal(t, 1, typedChecker.VDelegatedCallRequest.Count())
		assert.Equal(t, 1, typedChecker.VStateReport.Count())
	}

	mc.Finish()
}

func TestDeduplication_Constructor_AfterExecution_DifferentPulses(t *testing.T) {
	t.Log("C5005")

	var (
		mc = minimock.NewController(t)
	)

	server, ctx := utils.NewUninitializedServer(nil, t)
	defer server.Stop()

	executeDone := server.Journal.WaitStopOf(&execute.SMExecute{}, 2)

	runnerMock := logicless.NewServiceMock(ctx, t, nil)
	server.ReplaceRunner(runnerMock)
	server.Init(ctx)

	var (
		isolation = contract.ConstructorIsolation()
		outgoing  = server.RandomLocalWithPulse()
		class     = gen.UniqueReference()
	)

	pl := payload.VCallRequest{
		CallType:       payload.CTConstructor,
		CallFlags:      payload.BuildCallFlags(isolation.Interference, isolation.State),
		Callee:         class,
		CallSiteMethod: "New",
		CallOutgoing:   outgoing,
	}

	synchronizeExecution := NewSynchronizationPoint(1)

	{
		requestResult := requestresult.New([]byte("123"), gen.UniqueReference())
		requestResult.SetActivate(gen.UniqueReference(), class, []byte("234"))

		executionMock := runnerMock.AddExecutionMock(calculateOutgoing(pl).String())
		executionMock.AddStart(nil, &executionupdate.ContractExecutionStateUpdate{
			Type:   executionupdate.Done,
			Result: requestResult,
		})
	}

	typedChecker := server.PublisherMock.SetTypedChecker(ctx, t, server)
	typedChecker.VCallResult.Set(func(result *payload.VCallResult) bool {
		synchronizeExecution.Synchronize()
		return false
	})
	typedChecker.VStateReport.SetResend(true)

	{
		msg := server.WrapPayload(&pl).Finalize()
		server.SendMessage(ctx, msg)
	}

	synchronizeExecution.Wait(t)
	server.IncrementPulse(ctx)
	{
		server.SuspendConveyorAndWaitThenResetActive()
		server.SendMessage(ctx, server.WrapPayload(&pl).Finalize())
		server.WaitActiveThenIdleConveyor()
	}
	synchronizeExecution.WakeUp()

	{
		select {
		case <-executeDone:
		case <-time.After(10 * time.Second):
			require.FailNow(t, "timeout")
		}

		select {
		case <-server.Journal.WaitAllAsyncCallsDone():
		case <-time.After(10 * time.Second):
			require.FailNow(t, "timeout")
		}
	}

	{
		assert.Equal(t, 1, typedChecker.VCallResult.Count())
		assert.Equal(t, 1, typedChecker.VStateReport.Count())
	}

	mc.Finish()
}

func TestDeduplication_Constructor_DuringExecution_SamePulse(t *testing.T) {
	t.Log("C4998")

	server, ctx := utils.NewUninitializedServer(nil, t)
	defer server.Stop()

	executeDone := server.Journal.WaitStopOf(&execute.SMExecute{}, 2)

	runnerMock := logicless.NewServiceMock(ctx, t, nil)
	server.ReplaceRunner(runnerMock)
	server.Init(ctx)

	var (
		isolation = contract.ConstructorIsolation()
		outgoing  = server.RandomLocalWithPulse()
		class     = gen.UniqueReference()
	)

	pl := payload.VCallRequest{
		CallType:       payload.CTConstructor,
		CallFlags:      payload.BuildCallFlags(isolation.Interference, isolation.State),
		Callee:         class,
		CallSiteMethod: "New",
		CallOutgoing:   outgoing,
	}

	synchronizeExecution := NewSynchronizationPoint(1)

	{
		requestResult := requestresult.New([]byte("123"), gen.UniqueReference())
		requestResult.SetActivate(gen.UniqueReference(), class, []byte("234"))

		executionMock := runnerMock.AddExecutionMock(calculateOutgoing(pl).String())
		executionMock.AddStart(func(ctx execution.Context) {
			synchronizeExecution.Synchronize()
		}, &executionupdate.ContractExecutionStateUpdate{
			Type:   executionupdate.Done,
			Result: requestResult,
		})
	}

	typedChecker := server.PublisherMock.SetTypedChecker(ctx, t, server)
	typedChecker.VCallResult.SetResend(false)

	{
		msg := server.WrapPayload(&pl).Finalize()
		server.SendMessage(ctx, msg)
	}

	synchronizeExecution.Wait(t)
	{
		server.SuspendConveyorAndWaitThenResetActive()
		server.SendMessage(ctx, server.WrapPayload(&pl).Finalize())
		server.WaitActiveThenIdleConveyor()
	}
	synchronizeExecution.WakeUp()

	{
		select {
		case <-executeDone:
		case <-time.After(10 * time.Second):
			require.FailNow(t, "timeout")
		}

		select {
		case <-server.Journal.WaitAllAsyncCallsDone():
		case <-time.After(10 * time.Second):
			require.FailNow(t, "timeout")
		}
	}

	{
		assert.Equal(t, 1, typedChecker.VCallResult.Count())
	}
}

func TestDeduplication_Constructor_AfterExecution_SamePulse(t *testing.T) {
	t.Log("C5005")

	var (
		mc = minimock.NewController(t)
	)

	server, ctx := utils.NewUninitializedServer(nil, t)
	defer server.Stop()

	executeDone := server.Journal.WaitStopOf(&execute.SMExecute{}, 2)

	runnerMock := logicless.NewServiceMock(ctx, t, nil)
	server.ReplaceRunner(runnerMock)
	server.Init(ctx)

	var (
		isolation = contract.ConstructorIsolation()
		outgoing  = server.RandomLocalWithPulse()
		class     = gen.UniqueReference()
	)

	pl := payload.VCallRequest{
		CallType:       payload.CTConstructor,
		CallFlags:      payload.BuildCallFlags(isolation.Interference, isolation.State),
		Callee:         class,
		CallSiteMethod: "New",
		CallOutgoing:   outgoing,
	}

	synchronizeExecution := NewSynchronizationPoint(1)

	{
		requestResult := requestresult.New([]byte("123"), gen.UniqueReference())
		requestResult.SetActivate(gen.UniqueReference(), class, []byte("234"))

		executionMock := runnerMock.AddExecutionMock(calculateOutgoing(pl).String())
		executionMock.AddStart(nil, &executionupdate.ContractExecutionStateUpdate{
			Type:   executionupdate.Done,
			Result: requestResult,
		})
	}

	typedChecker := server.PublisherMock.SetTypedChecker(ctx, t, server)
	typedChecker.VCallResult.Set(func(result *payload.VCallResult) bool {
		synchronizeExecution.Synchronize()
		return false
	})

	{
		msg := server.WrapPayload(&pl).Finalize()
		server.SendMessage(ctx, msg)
	}

	synchronizeExecution.Wait(t)
	{
		server.SuspendConveyorAndWaitThenResetActive()
		server.SendMessage(ctx, server.WrapPayload(&pl).Finalize())
		server.WaitActiveThenIdleConveyor()
	}
	synchronizeExecution.WakeUp()

	{
		select {
		case <-executeDone:
		case <-time.After(10 * time.Second):
			require.FailNow(t, "timeout")
		}

		select {
		case <-server.Journal.WaitAllAsyncCallsDone():
		case <-time.After(10 * time.Second):
			require.FailNow(t, "timeout")
		}
	}

	{
		assert.Equal(t, 1, typedChecker.VCallResult.Count())
	}

	mc.Finish()
}
