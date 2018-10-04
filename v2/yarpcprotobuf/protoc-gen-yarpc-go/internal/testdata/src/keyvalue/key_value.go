package keyvalue

import (
	"context"
	"fmt"
	"sync"

	commonpb "go.uber.org/yarpc/v2/yarpcprotobuf/protoc-gen-yarpc-go/internal/testdata/gen/proto/src/common"
	keyvaluepb "go.uber.org/yarpc/v2/yarpcprotobuf/protoc-gen-yarpc-go/internal/testdata/gen/proto/src/keyvalue"
)

type kvServer struct {
	rw sync.RWMutex

	store map[string]string
}

// NewServer returns a new keyvaluepb.StoreServer.
func NewServer() keyvaluepb.StoreServer {
	return &kvServer{
		store: make(map[string]string),
	}
}

func (s *kvServer) Get(ctx context.Context, req *commonpb.GetRequest) (*commonpb.GetResponse, error) {
	s.rw.RLock()
	defer s.rw.RUnlock()

	val, ok := s.store[req.GetKey()]
	if !ok {
		return nil, fmt.Errorf("failed to find value for key: %q", req.Key)
	}
	return &commonpb.GetResponse{Value: val}, nil
}

func (s *kvServer) Set(ctx context.Context, req *commonpb.SetRequest) (*commonpb.SetResponse, error) {
	s.rw.Lock()
	defer s.rw.Unlock()

	s.store[req.GetKey()] = req.GetValue()
	return &commonpb.SetResponse{}, nil
}
