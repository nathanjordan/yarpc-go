package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	internaltemplate "go.uber.org/yarpc/v2/yarpcprotobuf/protoc-gen-yarpc-go/internal/template"
)

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("%v\n", err))
		os.Exit(1)
	}
}

func run(input io.Reader, output io.Writer) error {
	tmpl := template.Must(parseTemplates(internaltemplate.Client, internaltemplate.Server))
	req, err := fileToGeneratorRequest(input)
	if err != nil {
		return fmt.Errorf("failed to create CodeGeneratorRequest: %v", err)
	}

	res, err := generate(req, tmpl)
	if err != nil {
		return fmt.Errorf("failed to create CodeGeneratorResponse: %v", err)
	}

	out, err := proto.Marshal(res)
	if err != nil {
		return fmt.Errorf("failed to marshal CodeGeneratorResponse: %v", err)
	}

	_, err = output.Write(out)
	if err != nil {
		return fmt.Errorf("failed to write protoc-gen-yarpc-go output: %v", err)
	}
	return nil
}

func fileToGeneratorRequest(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	in, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	req := &plugin.CodeGeneratorRequest{}
	return req, proto.Unmarshal(in, req)
}

func generate(req *plugin.CodeGeneratorRequest, t *template.Template) (*plugin.CodeGeneratorResponse, error) {
	var files []*plugin.CodeGeneratorResponse_File
	targets := getTargetFiles(req.GetFileToGenerate())
	for _, f := range req.GetProtoFile() {
		filename := f.GetName()
		if _, ok := targets[filename]; !ok {
			continue
		}
		data, err := getTemplateData(f)
		if err != nil {
			return nil, err
		}
		out, err := execTemplate(t, data)
		if err != nil {
			return nil, err
		}
		files = append(files, &plugin.CodeGeneratorResponse_File{
			Name:    &filename,
			Content: &out,
		})

	}
	return &plugin.CodeGeneratorResponse{
		File: files,
	}, nil
}

func getTemplateData(f *descriptor.FileDescriptorProto) (*internaltemplate.Data, error) {
	return &internaltemplate.Data{
		Filename: f.GetName(),
		Imports: []string{
			"context",
			"io/ioutil",
			"reflect",
			"github.com/gogo/protobuf/proto",
			"go.uber.org/fx",
			"go.uber.org/v2/yarpc",
			"go.uber.org/yarpc/v2/yarpcprotobuf",
		},
		Services: []internaltemplate.Service{
			{
				Name: "Foo",
				UnaryMethods: []internaltemplate.Method{
					{
						Name:     "FooMethod",
						Request:  "FooRequest",
						Response: "FooResponse",
					},
				},
				StreamMethods: []internaltemplate.StreamMethod{
					{
						Method: internaltemplate.Method{
							Name:     "BarMethod",
							Request:  "BarRequest",
							Response: "BarResponse",
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

func parseTemplates(templates ...string) (*template.Template, error) {
	t := template.New(path.Base(internaltemplate.Base))
	for _, tmpl := range templates {
		_, err := t.Parse(tmpl)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func execTemplate(tmpl *template.Template, data interface{}) (string, error) {
	buffer := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
