/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package peer

import (
	"github.com/op/go-logging"
	pb "github.com/bft/protos"
)

// Peer provides interface for a peer
type Peer interface {
	GetPeerEndpoint() (*pb.PeerEndpoint, error)
	NewOpenchainDiscoveryHello() (*pb.Message, error)
}

// MessageHandler standard interface for handling Openchain messages.
type MessageHandler interface {
	HandleMessage(msg *pb.Message) error
	SendMessage(msg *pb.Message) error
	To() (pb.PeerEndpoint, error)
	Stop() error
}

// MessageHandlerCoordinator responsible for coordinating between the registered MessageHandler's
type MessageHandlerCoordinator interface {
	Peer
	RegisterHandler(messageHandler MessageHandler) error
	DeregisterHandler(messageHandler MessageHandler) error
	Broadcast(*pb.Message, pb.PeerEndpoint_Type) []error
	Unicast(*pb.Message, *pb.PeerID) error
	GetPeers() (*pb.PeersMessage, error)
	PeersDiscovered(*pb.PeersMessage) error
	ExecuteTransaction(transaction *pb.Transaction) *pb.Response
}

// Coordinator is used to initiate state transfer.  Start must be called before use, and Stop should be called to free allocated resources
type StateTransferCoordinator interface {
	Start() // Start the block transfer go routine
	Stop()  // Stop up the block transfer go routine

	// SyncToTarget attempts to move the state to the given target, returning an error, and whether this target might succeed if attempted at a later time
	SyncToTarget(blockNumber uint64, blockHash []byte, peerIDs []*pb.PeerID) (error, bool)
}

// ChatStream interface supported by stream between Peers
type ChatStream interface {
	Send(*pb.Message) error
	Recv() (*pb.Message, error)
}

// TransactionProccesor responsible for processing of Transactions
type TransactionProccesor interface {
	ProcessTransactionMsg(*pb.Message, *pb.Transaction) *pb.Response
}

// Engine Responsible for managing Peer network communications (Handlers) and processing of Transactions
type Engine interface {
	TransactionProccesor
	// GetHandlerFactory return a handler for an accepted Chat stream
	GetHandlerFactory() HandlerFactory
	//GetInputChannel() (chan<- *pb.Transaction, error)
}

var peerLogger = logging.MustGetLogger("peer")

// HandlerFactory for creating new MessageHandlers
type HandlerFactory func(MessageHandlerCoordinator, ChatStream, bool) (MessageHandler, error)

// EngineFactory for creating new engines
type EngineFactory func(MessageHandlerCoordinator) (Engine, error)
