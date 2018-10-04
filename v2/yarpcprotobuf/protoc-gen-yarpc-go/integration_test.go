package main

import (
	"testing"

	keyvaluepb "go.uber.org/yarpc/v2/yarpcprotobuf/protoc-gen-yarpc-go/internal/testdata/gen/proto/src/keyvalue"
	keyvalue "go.uber.org/yarpc/v2/yarpcprotobuf/protoc-gen-yarpc-go/internal/testdata/src/keyvalue"
	"go.uber.org/yarpc/v2/yarpcrouter"
)

func TestIntegration(t *testing.T) {
	t.Skip("TODO(mensch): Use the gRPC transport when available")
	router := yarpcrouter.NewMapRouter("keyvalue")
	procedures := keyvaluepb.BuildStoreProcedures(keyvalue.NewServer())
	router.Register(procedures)
}
