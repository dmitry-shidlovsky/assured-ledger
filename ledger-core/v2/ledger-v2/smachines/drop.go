package smachines

import (
	"github.com/insolar/assured-ledger/ledger-core/v2/conveyor/smachine"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/record"
	"github.com/insolar/assured-ledger/ledger-core/v2/vanilla/injector"
)

type sharedDropBatch struct {
	records      []record.Record
	recordNumber int
	syncFinished smachine.SyncLink
}

func (b *sharedDropBatch) appendRecords(recs []record.Record) {
	b.records = append(b.records, recs...)
	b.recordNumber += len(recs)
}

/* -------- Declaration ------------- */

var declDropBatch smachine.StateMachineDeclaration = &declarationDropBatch{}

type declarationDropBatch struct {
	smachine.StateMachineDeclTemplate
}

func (*declarationDropBatch) InjectDependencies(_ smachine.StateMachine, _ smachine.SlotLink, _ *injector.DependencyInjector) {
}

func (*declarationDropBatch) GetInitStateFor(sm smachine.StateMachine) smachine.InitFunc {
	return sm.(*Object).Init
}

/* -------- Instance ------------- */

type DropBatch struct {
	jetID insolar.JetID

	ownedDropBatch sharedDropBatch
	syncFinished   smachine.BoolConditionalLink
}

func NewDropBatch(jetID insolar.JetID) *DropBatch {
	return &DropBatch{jetID: jetID}
}

func (s *DropBatch) GetStateMachineDeclaration() smachine.StateMachineDeclaration {
	return declDropBatch
}

func (s *DropBatch) Init(ctx smachine.InitializationContext) smachine.StateUpdate {
	s.syncFinished = smachine.NewConditionalBool(false, "syncFinished")
	s.ownedDropBatch.syncFinished = s.syncFinished.SyncLink()

	link := ctx.Share(&s.ownedDropBatch, smachine.ShareDataWakesUpAfterUse)
	if !ctx.Publish(s.jetID, link) {
		return ctx.Stop()
	}

	return ctx.Jump(s.waitForBatch)
}

func (s *DropBatch) waitForBatch(ctx smachine.ExecutionContext) smachine.StateUpdate {
	// TODO: calculate hash and replicate records.
	ctx.ApplyAdjustment(s.syncFinished.NewValue(true))
	return ctx.Stop()
}

func sharedDropBatchLink(ctx smachine.ExecutionContext, jetID insolar.JetID) smachine.SharedDataLink {
	link := ctx.GetPublishedLink(jetID)
	if link.IsZero() {
		ctx.InitChild(func(ctx smachine.ConstructionContext) smachine.StateMachine {
			return NewDropBatch(jetID)
		})
		link = ctx.GetPublishedLink(jetID)
		if link.IsZero() {
			panic("failed to acquire shared drop")
		}
	}
	return link
}
