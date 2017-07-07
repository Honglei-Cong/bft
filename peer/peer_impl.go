package peer

import (
	pb "github.com/bft/protos"
	"context"
	"github.com/golang/protobuf/proto"
	"fmt"
	"github.com/bft/util"
)

type peerImpl struct {
	engine Engine
}

func NewPeer(engine Engine) (*peerImpl, error){
	return &peerImpl{
		engine: engine,
	}, nil
}

func (p *peerImpl) Unicast(msg *pb.Message, receiverHandle *pb.PeerID) error {
	return nil
}

func (p *peerImpl) Broadcast(msg *pb.Message, typ pb.PeerEndpoint_Type) []error {
	return nil
}

func (p *peerImpl) GetPeerEndpoint() (*pb.PeerEndpoint, error) {
	return nil, nil
}

func (p *peerImpl) GetPeers() (*pb.PeersMessage, error) {
	return nil, nil
}

func (p *peerImpl) ProcessTransaction(ctx context.Context, tx *pb.Transaction)(*pb.Response, error) {

	data, err := proto.Marshal(tx)
	if err != nil {
		return &pb.Response{
			Status: pb.Response_FAILURE,
			Msg: []byte(fmt.Sprintf("Error sending txn to consenter: %s", err)),
		}, nil
	}

	msg := &pb.Message{
		Type: pb.Message_CHAIN_TRANSACTION,
		Payload: data,
		Timestamp: util.CreateUtcTimestamp(),
	}

	return p.engine.ProcessTransactionMsg(msg, tx), nil
}
