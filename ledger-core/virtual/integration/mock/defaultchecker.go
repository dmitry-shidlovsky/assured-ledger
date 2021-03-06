// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package mock

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

var _ Checker = &DefaultPublishChecker{}

type DefaultPublishChecker struct {
	checkerFn CheckerFn
}

func NewDefaultPublishChecker(fn CheckerFn) *DefaultPublishChecker {
	return &DefaultPublishChecker{
		checkerFn: fn,
	}
}

func (p *DefaultPublishChecker) CheckMessages(topic string, messages ...*message.Message) error {
	return p.checkerFn(topic, messages...)
}

func NewResenderPublishChecker(ctx context.Context, sender Sender) *DefaultPublishChecker {
	return &DefaultPublishChecker{
		checkerFn: func(topic string, messages ...*message.Message) error {
			for _, msg := range messages {
				sender.SendMessage(ctx, msg)
			}
			return nil
		},
	}
}
