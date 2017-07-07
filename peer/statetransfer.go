package peer

import (
	pb "github.com/bft/protos"
)

type stateTransferCoordinatorImpl struct {
}

func NewCoordiatorImpl(peer MessageHandlerCoordinator) StateTransferCoordinator {
	stc := &stateTransferCoordinatorImpl{}

	return stc
}

func (stc *stateTransferCoordinatorImpl) Start() {
}

func (stc *stateTransferCoordinatorImpl) Stop() {
}

func (stc *stateTransferCoordinatorImpl) SyncToTarget(blockNumber uint64, blockHash []byte, peerIDs []*pb.PeerID) (error, bool) {
	return nil, false
}
