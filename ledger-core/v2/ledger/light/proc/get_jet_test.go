// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package proc_test

import (
	"context"
	"testing"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"

	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/bus"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/gen"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/jet"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/payload"
	"github.com/insolar/assured-ledger/ledger-core/v2/instrumentation/inslogger"
	"github.com/insolar/assured-ledger/ledger-core/v2/ledger/light/proc"
	"github.com/insolar/assured-ledger/ledger-core/v2/pulse"
)

func TestGetJet_Proceed(t *testing.T) {
	ctx := inslogger.TestContext(t)
	mc := minimock.NewController(t)

	var (
		jetAccessor *jet.AccessorMock
		sender      *bus.SenderMock
	)

	setup := func() {
		jetAccessor = jet.NewAccessorMock(mc)
		sender = bus.NewSenderMock(mc)
	}

	t.Run("basic ok", func(t *testing.T) {
		setup()
		defer mc.Finish()

		jetID := gen.JetID()
		jetAccessor.ForIDMock.Return(jetID, true)

		expectedMsg, _ := payload.NewMessage(&payload.Jet{
			JetID:  jetID,
			Actual: true,
		})

		sender.ReplyMock.Inspect(func(ctx context.Context, origin payload.Meta, reply *message.Message) {
			assert.Equal(t, expectedMsg.Payload, reply.Payload)
		}).Return()

		p := proc.NewGetJet(payload.Meta{}, gen.ID(), pulse.MinTimePulse)
		p.Dep(jetAccessor, sender)
		err := p.Proceed(ctx)
		assert.NoError(t, err)
	})

}
