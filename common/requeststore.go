/*

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

package common

import (
	"container/list"
    pb "github.com/bft/bftprotos"
	"github.com/bft/util"
)

type RequestContainer struct {
	key string
	req *pb.Request
}

type OrderedRequests struct {
	order    list.List
	presence map[string]*list.Element
}

func (a *OrderedRequests) Len() int {
	return a.order.Len()
}

func (a *OrderedRequests) RrapRequest(req *pb.Request) RequestContainer {
	return RequestContainer{
		key: util.Hash(req),
		req: req,
	}
}

func (a *OrderedRequests) Has(key string) bool {
	_, ok := a.presence[key]
	return ok
}

func (a *OrderedRequests) add(request *pb.Request) {
	rc := a.RrapRequest(request)
	if !a.Has(rc.key) {
		e := a.order.PushBack(rc)
		a.presence[rc.key] = e
	}
}

func (a *OrderedRequests) Adds(requests []*pb.Request) {
	for _, req := range requests {
		a.add(req)
	}
}

func (a *OrderedRequests) Remove(request *pb.Request) bool {
	rc := a.RrapRequest(request)
	e, ok := a.presence[rc.key]
	if !ok {
		return false
	}
	a.order.Remove(e)
	delete(a.presence, rc.key)
	return true
}

func (a *OrderedRequests) Removes(requests []*pb.Request) bool {
	allSuccess := true
	for _, req := range requests {
		if !a.Remove(req) {
			allSuccess = false
		}
	}

	return allSuccess
}

func (a *OrderedRequests) Empty() {
	a.order.Init()
	a.presence = make(map[string]*list.Element)
}

type RequestStore struct {
	OutstandingRequests *OrderedRequests
	PendingRequests     *OrderedRequests
}

// NewRequestStore creates a new RequestStore.
func NewRequestStore() *RequestStore {
	rs := &RequestStore{
		OutstandingRequests: &OrderedRequests{},
		PendingRequests:     &OrderedRequests{},
	}
	// initialize data structures
	rs.OutstandingRequests.Empty()
	rs.PendingRequests.Empty()

	return rs
}

// StoreOutstanding Adds a request to the outstanding request list
func (rs *RequestStore) StoreOutstanding(request *pb.Request) {
	rs.OutstandingRequests.add(request)
}

// StorePending Adds a request to the pending request list
func (rs *RequestStore) StorePending(request *pb.Request) {
	rs.PendingRequests.add(request)
}

// StorePending Adds a slice of requests to the pending request list
func (rs *RequestStore) StorePendings(requests []*pb.Request) {
	rs.PendingRequests.Adds(requests)
}

// Remove deletes the request from both the outstanding and pending lists, it returns whether it was found in each list respectively
func (rs *RequestStore) Remove(request *pb.Request) (outstanding, pending bool) {
	outstanding = rs.OutstandingRequests.Remove(request)
	pending = rs.PendingRequests.Remove(request)
	return
}

// GetNextNonPending returns up to the next n outstanding, but not pending requests
func (rs *RequestStore) HasNonPending() bool {
	return rs.OutstandingRequests.Len() > rs.PendingRequests.Len()
}

// GetNextNonPending returns up to the next n outstanding, but not pending requests
func (rs *RequestStore) GetNextNonPending(n int) (result []*pb.Request) {
	for oreqc := rs.OutstandingRequests.order.Front(); oreqc != nil; oreqc = oreqc.Next() {
		oreq := oreqc.Value.(RequestContainer)
		if rs.PendingRequests.Has(oreq.key) {
			continue
		}
		result = append(result, oreq.req)
		if len(result) == n {
			break
		}
	}

	return result
}
