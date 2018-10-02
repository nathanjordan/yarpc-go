package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
)

const _plugin = "protoc-gen-yarpc-go"

// generator orchestrates code generation for protoc-gen-yarpc-go.
type generator struct {
	tmpl     *template.Template
	registry *registry
}

func newGenerator(r *registry) *generator {
	return &generator{
		tmpl: template.Must(
			parseTemplates(
				_baseTemplate,
				_clientTemplate,
				_serverTemplate,
			),
		),
		registry: r,
	}
}

func parseTemplates(templates ...string) (*template.Template, error) {
	t := template.New(_plugin)
	for _, tmpl := range templates {
		_, err := t.Parse(tmpl)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

// Generate uses the given *descriptor.FileDescriptorProto to generate
// YARPC client and server stubs.
func Generate(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	r, err := newRegistry(req)
	if err != nil {
		return nil, err
	}
	g := newGenerator(r)

	targets := getTargetFiles(req.GetFileToGenerate())
	var files []*plugin.CodeGeneratorResponse_File
	for _, f := range req.GetProtoFile() {
		filename := f.GetName()
		if _, ok := targets[filename]; !ok {
			continue
		}
		data, err := descriptorToFileData(f)
		if err != nil {
			return nil, err
		}
		raw, err := g.execTemplate(data)
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

func (g *generator) execTemplate(data interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := g.tmpl.Execute(buffer, data); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// TODO(mensch): Dummy data used for testing.
func descriptorToFileData(f *descriptor.FileDescriptorProto) (File, error) {
	return File{}, nil
}

func getTargetFiles(ts []string) map[string]struct{} {
	m := make(map[string]struct{}, len(ts))
	for _, t := range ts {
		m[t] = struct{}{}
	}
	return m
}
