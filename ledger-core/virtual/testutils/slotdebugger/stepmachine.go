// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package slotdebugger

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"

	"github.com/insolar/assured-ledger/ledger-core/runner"
	"github.com/insolar/assured-ledger/ledger-core/runner/machine"
	"github.com/insolar/assured-ledger/ledger-core/testutils/slotdebugger"
	"github.com/insolar/assured-ledger/ledger-core/virtual/testutils"
)

type VirtualStepController struct {
	*slotdebugger.StepController

	RunnerDescriptorCache *testutils.DescriptorCacheMockWrapper
	MachineManager        machine.Manager
}

func New(ctx context.Context, t *testing.T, suppressLogError bool) *VirtualStepController {
	stepController := slotdebugger.New(ctx, t, suppressLogError)

	w := &VirtualStepController{
		StepController: stepController,
	}

	return w
}

func (c *VirtualStepController) PrepareRunner(ctx context.Context, mc *minimock.Controller) {
	c.RunnerDescriptorCache = testutils.NewDescriptorsCacheMockWrapper(mc)
	c.MachineManager = machine.NewManager()

	runnerService := runner.NewService()
	runnerService.Manager = c.MachineManager
	runnerService.Cache = c.RunnerDescriptorCache.Mock()

	runnerAdapter := runnerService.CreateAdapter(ctx)
	c.SlotMachine.AddInterfaceDependency(&runnerAdapter)
}