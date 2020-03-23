// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package adapters

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"time"

	"github.com/insolar/assured-ledger/ledger-core/v2/insolar"
	"github.com/insolar/assured-ledger/ledger-core/v2/network/consensus/gcpv2/api"
	"github.com/insolar/assured-ledger/ledger-core/v2/vanilla/cryptkit"
)

var defaultRoundTimings = api.RoundTimings{
	StartPhase0At: 100 * time.Millisecond, // Not scaled

	StartPhase1RetryAt: 0 * time.Millisecond, // 0 for no retries
	EndOfPhase1:        180 * time.Millisecond,
	EndOfPhase2:        250 * time.Millisecond,
	EndOfPhase3:        400 * time.Millisecond,
	EndOfConsensus:     600 * time.Millisecond,

	BeforeInPhase2ChasingDelay: 0 * time.Millisecond,
	BeforeInPhase3ChasingDelay: 0 * time.Millisecond,
}

var defaultEphemeralTimings = api.RoundTimings{
	StartPhase0At: 100 * time.Millisecond, // Not scaled

	StartPhase1RetryAt: 0 * time.Millisecond, // 0 for no retries
	EndOfPhase1:        200 * time.Millisecond,
	EndOfPhase2:        600 * time.Millisecond,
	EndOfPhase3:        800 * time.Millisecond,
	EndOfConsensus:     900 * time.Millisecond,

	BeforeInPhase2ChasingDelay: 0 * time.Millisecond,
	BeforeInPhase3ChasingDelay: 0 * time.Millisecond,
}

// var _ api.LocalNodeConfiguration = &LocalNodeConfiguration{}

type LocalNodeConfiguration struct {
	ctx              context.Context
	timings          api.RoundTimings
	ephemeralTimings api.RoundTimings
	secretKeyStore   cryptkit.SecretKeyStore
	pulsarKeys       []crypto.PublicKey
}

func (c *LocalNodeConfiguration) GetNodeCountHint() int {
	return 10 // should provide some rough estimate of a size of a network to be joined
}

func NewLocalNodeConfiguration(ctx context.Context, keyStore insolar.KeyStore, pulsarKeys []crypto.PublicKey) *LocalNodeConfiguration {
	privateKey, err := keyStore.GetPrivateKey("")
	if err != nil {
		panic(err)
	}

	ecdsaPrivateKey := privateKey.(*ecdsa.PrivateKey)

	return &LocalNodeConfiguration{
		ctx:              ctx,
		timings:          defaultRoundTimings,
		ephemeralTimings: defaultEphemeralTimings,
		secretKeyStore:   NewECDSASecretKeyStore(ecdsaPrivateKey),
		pulsarKeys:       pulsarKeys,
	}
}

func (c *LocalNodeConfiguration) GetParentContext() context.Context {
	return c.ctx
}

func (c *LocalNodeConfiguration) getConsensusTimings(t api.RoundTimings, nextPulseDelta uint16) api.RoundTimings {
	if nextPulseDelta == 1 {
		return t
	}
	m := time.Duration(nextPulseDelta) // this is NOT a duration, but a multiplier

	t.StartPhase0At *= 1 // don't scale!
	t.StartPhase1RetryAt *= m
	t.EndOfPhase1 *= m
	t.EndOfPhase2 *= m
	t.EndOfPhase3 *= m
	t.EndOfConsensus *= m
	t.BeforeInPhase2ChasingDelay *= m
	t.BeforeInPhase3ChasingDelay *= m

	return t
}

func (c *LocalNodeConfiguration) GetPulsarPublicKeys() []crypto.PublicKey {
	return append(c.pulsarKeys[:0:0], c.pulsarKeys...)
}

func (c *LocalNodeConfiguration) GetConsensusTimings(nextPulseDelta uint16) api.RoundTimings {
	return c.getConsensusTimings(c.timings, nextPulseDelta)
}

func (c *LocalNodeConfiguration) GetEphemeralTimings(nextPulseDelta uint16) api.RoundTimings {
	return c.getConsensusTimings(c.ephemeralTimings, nextPulseDelta)
}

func (c *LocalNodeConfiguration) GetSecretKeyStore() cryptkit.SecretKeyStore {
	return c.secretKeyStore
}

type ConsensusConfiguration struct{}

func NewConsensusConfiguration() *ConsensusConfiguration {
	return &ConsensusConfiguration{}
}
