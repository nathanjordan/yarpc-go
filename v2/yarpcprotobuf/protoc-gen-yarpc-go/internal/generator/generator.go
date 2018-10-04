package generator

import (
	"fmt"
	"go/format"
	"path/filepath"
	"strings"

	"github.com/gogo/protobuf/proto"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
)

const _plugin = "protoc-gen-yarpc-go"

// Generate uses the given *descriptor.FileDescriptorProto to generate
// YARPC client and server stubs.
func Generate(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	r, err := newRegistry(req)
	if err != nil {
		return nil, err
	}

	targets := getTargetFiles(req.GetFileToGenerate())
	var files []*plugin.CodeGeneratorResponse_File
	for _, f := range req.GetProtoFile() {
		filename := f.GetName()
		if _, ok := targets[filename]; !ok {
			continue
		}
		data, err := r.GetData(filename)
		if err != nil {
			return nil, err
		}
		raw, err := execTemplate(data)
		if err != nil {
			return nil, err
		}
		formatted, err := format.Source(raw)
		if err != nil {
			return nil, err
		}
		files = append(files, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(fmt.Sprintf("%s.pb.yarpc.go", strings.TrimSuffix(filename, filepath.Ext(filename)))),
			Content: proto.String(string(formatted)),
		})

	}
	return &plugin.CodeGeneratorResponse{
		File: files,
	}, nil
}

func getTargetFiles(ts []string) map[string]struct{} {
	m := make(map[string]struct{}, len(ts))
	for _, t := range ts {
		m[t] = struct{}{}
	}
	return m
}
