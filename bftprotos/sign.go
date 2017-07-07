package bftprotos

import pb "github.com/golang/protobuf/proto"

func (vc *ViewChange) getSignature() []byte {
	return vc.Signature
}

func (vc *ViewChange) setSignature(sig []byte) {
	vc.Signature = sig
}

func (vc *ViewChange) getID() uint64 {
	return vc.ReplicaId
}

func (vc *ViewChange) setID(id uint64) {
	vc.ReplicaId = id
}

func (vc *ViewChange) serialize() ([]byte, error) {
	return pb.Marshal(vc)
}

