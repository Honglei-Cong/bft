package rbft

import (
	pb "github.com/bft/bftprotos"
	"github.com/bft/util/events"
	"time"
	"github.com/bft/comm"
	"github.com/bft/common"
	"github.com/bft/protos"
)
type RbftBatch struct {

	broadcaster *comm.Broadcaster

	batchSize int
	batchStore []*pb.Request
	batchTimer events.Timer
	batchTimerActive bool
	batchTimerDuration time.Duration

	incomingChan chan *pb.BatchMessage
	idleChan chan struct{}

	reqStore *common.RequestStore
}

func newRbftBatch (id uint64) *RbftBatch {
	return nil
}

func (rb *RbftBatch) submitToLeader(req *pb.Request) events.Event {
	return nil
}

func (rb *RbftBatch) broadcastMsg(msg *pb.BatchMessage) error {
	return nil
}

func (rb *RbftBatch) unicastMsg(msg *pb.BatchMessage, receiverID uint64) error {
	return nil
}

func (rb *RbftBatch) execute(seqNo uint64, reqBatch *pb.RequestBatch) error {
	return nil
}

func (rb *RbftBatch) leaderProcessReq(req *pb.Request) events.Event {
	return nil
}

func (rb *RbftBatch) ProcessMessage(msg *protos.Message, senderHandle *protos.PeerID) events.Event {
	return nil
}

func (rb *RbftBatch) ProcessEvent(event events.Event) events.Event {
	return nil
}