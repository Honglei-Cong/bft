/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package util

import (
	"encoding/base64"
	"crypto/sha1"
	"github.com/bft/protos"
	pb "github.com/bft/bftprotos"
	"github.com/golang/protobuf/proto"
	"strconv"
)

// ComputeCryptoHash should be used in openchain code so that we can change the actual algo used for crypto-Hash at one place
func ComputeCryptoHash(data []byte) []byte {
	return []byte(sha1.Sum(data))
}

func Hash(msg interface{}) string {
	var raw []byte
	switch converted := msg.(type) {
	case *pb.Request:
		raw, _ = proto.Marshal(converted)
	case *pb.RequestBatch:
		raw, _ = proto.Marshal(converted)
	default:
		logger.Error("Asked to Hash non-supported message type, ignoring")
		return ""
	}
	return base64.StdEncoding.EncodeToString(ComputeCryptoHash(raw))

}

// Returns the peer handle that corresponds to a validator ID (uint64 assigned to it for PBFT)
func GetValidatorHandle(id uint64) (handle *protos.PeerID, err error) {
	// as requested here: https://github.com/hyperledger/fabric/issues/462#issuecomment-170785410
	name := "vp" + strconv.FormatUint(id, 10)
	return &protos.PeerID{Name: name}, nil
}

// Returns the peer handles corresponding to a list of replica ids
func GetValidatorHandles(ids []uint64) (handles []*protos.PeerID) {
	handles = make([]*protos.PeerID, len(ids))
	for i, id := range ids {
		handles[i], _ = GetValidatorHandle(id)
	}
	return
}

