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

// generator orchestrates code generation for protoc-gen-yarpc-go.
type generator struct {
	tmpl *template.Template
}

func newGenerator() *generator {
	return &generator{
		tmpl: template.Must(
			parseTemplates(
				_baseTemplate,
				_clientTemplate,
				_serverTemplate,
			),
		),
	}
}

func parseTemplates(templates ...string) (*template.Template, error) {
	t := template.New("protoc-gen-yarpc-go")
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
	g := newGenerator()
	targets := getTargetFiles(req.GetFileToGenerate())
	var files []*plugin.CodeGeneratorResponse_File

	for _, f := range req.GetProtoFile() {
		filename := f.GetName()
		if _, ok := targets[filename]; !ok {
			continue
		}
		data, err := getData(f)
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
func getData(f *descriptor.FileDescriptorProto) (*Data, error) {
	return &Data{
		Filename: f.GetName(),
		Package:  "keyvaluepb",
		Imports: []string{
			"context",
			"io/ioutil",
			"github.com/gogo/protobuf/proto",
			"go.uber.org/fx",
			"go.uber.org/yarpc/v2/yarpc",
			"go.uber.org/yarpc/v2/yarpcprotobuf",
		},
		Services: []Service{
			{
				Name:         "KeyValue",
				UnaryMethods: []Method{},
				StreamMethods: []StreamMethod{
					{
						Method: Method{
							Name:     "Get",
							Request:  "GetRequest",
							Response: "GetResponse",
						},
						ClientSide: true,
						ServerSide: true,
					},
				},
			},
		},
	}, nil
}

func getTargetFiles(ts []string) map[string]struct{} {
	m := make(map[string]struct{}, len(ts))
	for _, t := range ts {
		m[t] = struct{}{}
	}
	return m
}
