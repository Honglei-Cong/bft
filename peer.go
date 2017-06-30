package bft

import (
	"crypto"
	pb "github.com/bft/protos"
)

// MessageHandler standard interface for handling Openchain messages.
type MessageHandler interface {
	HandleMessage(msg *pb.Message) error
	SendMessage(msg *pb.Message) error
	To() (pb.PeerEndpoint, error)
	Stop() error
}

// ChatStream interface supported by stream between Peers
type ChatStream interface {
	Send(*pb.Message) error
	Recv() (*pb.Message, error)
}

// MessageHandlerCoordinator responsible for coordinating between the registered MessageHandler's
type MessageHandlerCoordinator interface {
	RegisterHandler(messageHandler MessageHandler) error
	DeregisterHandler(messageHandler MessageHandler) error
	Broadcast(*pb.Message, pb.PeerEndpoint_Type) []error
	Unicast(*pb.Message, *pb.PeerID) error
	GetPeers() (*pb.PeersMessage, error)
	//GetRemoteLedger(receiver *pb.PeerID) (RemoteLedger, error)
	PeersDiscovered(*pb.PeersMessage) error
	ExecuteTransaction(transaction *pb.Transaction) *pb.Response
}

// Coordinator is used to initiate state transfer.  Start must be called before use, and Stop should be called to free allocated resources
type Coordinator interface {
	Start() // Start the block transfer go routine
	Stop()  // Stop up the block transfer go routine

	// SyncToTarget attempts to move the state to the given target, returning an error, and whether this target might succeed if attempted at a later time
	SyncToTarget(blockNumber uint64, blockHash []byte, peerIDs []*pb.PeerID) (error, bool)
}




