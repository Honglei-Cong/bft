package rbft

import (
	pb "github.com/bft/bftprotos"
	"github.com/bft/util/events"
)

type qidx struct {
	d string
	n uint64
}

type rbftCore struct {
	activeView bool
	bynatine bool

	id uint64
	f int // max number of tolerated faults
	N int // max peer count
	h uint64 // lower watermark

	K uint64 // checkpoint period
	replicaCount int // num of replicas

	seqNo uint64
	view uint64
	chkpts map[uint64]string

	pset map[uint64]*pb.ViewChange_PQ
	qset map[qidx]*pb.ViewChange_PQ
}

func newRbftCore(id uint64) *rbftCore {
	return nil
}

func (rc *rbftCore) ProcessEvent(e events.Event) events.Event {
	return nil
}