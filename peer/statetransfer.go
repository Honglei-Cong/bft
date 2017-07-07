package peer

import (
	pb "github.com/bft/protos"
)

// Coordinator is used to initiate state transfer.  Start must be called before use, and Stop should be called to free allocated resources
type StateTransferCoordinator interface {
	Start() // Start the block transfer go routine
	Stop()  // Stop up the block transfer go routine

	// SyncToTarget attempts to move the state to the given target, returning an error, and whether this target might succeed if attempted at a later time
	SyncToTarget(blockNumber uint64, blockHash []byte, peerIDs []*pb.PeerID) (error, bool)
}

func NewCoordiatorImpl(peer MessageHandlerCoordinator) StateTransferCoordinator {
	return nil
}
