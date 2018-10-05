package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	keyvaluepb "go.uber.org/yarpc/v2/yarpcprotobuf/protoc-gen-yarpc-go/internal/testdata/gen/proto/src/keyvalue"
	keyvalue "go.uber.org/yarpc/v2/yarpcprotobuf/protoc-gen-yarpc-go/internal/testdata/src/keyvalue"
)

func TestIntegration(t *testing.T) {
	// TODO(mensch): Use the gRPC transport when available
	procedures := keyvaluepb.BuildStoreProcedures(keyvalue.NewServer())
	assert.Equal(t, 4, len(procedures))
}
