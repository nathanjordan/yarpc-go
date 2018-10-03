package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
)

// registry is used to collect and register all
// of the Protobuf types relevant for
// protoc-gen-yarpc-go. This concept is inspired
// by the registry implemented for protoc-gen-grpc-gateway,
// but has been adapted so that it only concerns itself
// with a minimal feature set.
type registry struct {
	files    map[string]*File
	messages map[string]*Message
	imports  Imports
}

func newRegistry(req *plugin.CodeGeneratorRequest) (*registry, error) {
	r := &registry{
		files:    make(map[string]*File),
		messages: make(map[string]*Message),
		imports: newImports(
			"context",
			"io/ioutil",
			"github.com/gogo/protobuf/proto",
			"go.uber.org/fx",
			"go.uber.org/yarpc/v2/yarpc",
			"go.uber.org/yarpc/v2/yarpcprotobuf",
		),
	}
	return r, r.Load(req)
}

// Load registers all of the Proto types provided in the CodeGeneratorRequest
// with the registry. Note that all CodeGeneratorRequests SHOULD only request
// to generate a single go package. Otherwise, multiple go packages will
// be deposited in the same directory and will therefore not compile. This
// is synonymous to the protoc-gen-go plugin restriction.
func (r *registry) Load(req *plugin.CodeGeneratorRequest) error {
	for _, f := range req.GetProtoFile() {
		r.loadFile(f)
	}
	for _, name := range req.FileToGenerate {
		target, ok := r.files[name]
		if !ok {
			return fmt.Errorf("file target %q was not registered", name)
		}
		for _, s := range target.GetService() {
			if err := r.loadService(target, s); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetTemplate returns the template data the corresponds
// to the given filename.
func (r *registry) GetTemplateData(filename string) (*Data, error) {
	f, err := r.getFile(filename)
	if err != nil {
		return nil, err
	}
	return &Data{
		File:    f,
		Imports: r.imports,
	}, nil
}

// getFile returns the File that corresponds to the
// given filename.
func (r *registry) getFile(filename string) (*File, error) {
	f, ok := r.files[filename]
	if !ok {
		return nil, fmt.Errorf("file %q was not found", filename)
	}
	return f, nil
}

// getMessage returns the Message that corresponds to the
// given name. This method expects the input to be formed
// as an input or output type, such as .foo.Bar.
func (r *registry) getMessage(name string) (*Message, error) {
	// All input and output types are represented as
	// .$(Package).$(Message), so we explicitly trim
	// the leading '.' prefix.
	msg := strings.TrimPrefix(name, ".")
	m, ok := r.messages[msg]
	if !ok {
		return nil, fmt.Errorf("message %q was not found", msg)
	}
	return m, nil
}

// loadFile registers the given file's message types.
// Note that we load the messages for all files up-front so that
// all of the message types potentially referneced in the proto
// services can reference these types.
func (r *registry) loadFile(f *descriptor.FileDescriptorProto) {
	pkg := r.newPackage(f)
	file := &File{
		FileDescriptorProto: f,
		Name:                f.GetName(),
		Package:             pkg,
	}
	r.files[file.Name] = file
	for _, m := range f.GetMessageType() {
		r.loadMessage(file, m)
	}
}

func (r *registry) loadMessage(f *File, m *descriptor.DescriptorProto) {
	msg := &Message{
		Name:    m.GetName(),
		Package: f.Package,
	}
	r.messages[msg.key()] = msg

	for _, n := range m.GetNestedType() {
		r.loadMessage(f, n)
	}
}

func (r *registry) loadService(f *File, s *descriptor.ServiceDescriptorProto) error {
	svc := &Service{
		Name:    s.GetName(),
		Package: f.Package,
	}
	for _, m := range s.GetMethod() {
		method, err := r.newMethod(m, s.GetName())
		if err != nil {
			return err
		}
		svc.Methods = append(svc.Methods, method)
	}
	f.Services = append(f.Services, svc)
	return nil
}

func (r *registry) newMethod(m *descriptor.MethodDescriptorProto, svc string) (*Method, error) {
	request, err := r.getMessage(m.GetInputType())
	if err != nil {
		return nil, err
	}
	response, err := r.getMessage(m.GetOutputType())
	if err != nil {
		return nil, err
	}
	return &Method{
		Name:            m.GetName(),
		Service:         svc,
		Request:         request,
		Response:        response,
		ClientStreaming: m.GetClientStreaming(),
		ServerStreaming: m.GetServerStreaming(),
		ClientStream:    streamInterface(svc, m.GetName(), _clientStream),
		ServerStream:    streamInterface(svc, m.GetName(), _serverStream),
	}, nil
}

func (r *registry) newPackage(f *descriptor.FileDescriptorProto) *Package {
	return &Package{
		Name:      f.GetPackage(),
		GoPackage: goPackage(f),
		Alias:     r.imports.Add(importPath(f)),
	}
}

func goPackage(f *descriptor.FileDescriptorProto) string {
	if f.Options != nil && f.Options.GoPackage != nil {
		gopkg := f.Options.GetGoPackage()
		idx := strings.LastIndex(gopkg, "/")
		if idx < 0 {
			return gopkg
		}

		return gopkg[idx+1:]
	}

	pkg := f.GetPackage()
	if f.Package == nil {
		base := filepath.Base(f.GetName())
		ext := filepath.Ext(base)
		pkg = strings.TrimSuffix(base, ext)
	}
	return strings.Replace(pkg, ".", "_", -1)
}

// importPath returns the package import path that corresponds to
// the given file descriptor. If the go_package option explicitly configures
// an import path, use it. Otherwise, use the directory from which the
// Protobuf definition is defined.
//
//  Ex:
//   ./foo/bar/baz.proto
//   -> "foo/bar"
//
//   option go_package = "gen/proto:bazpb";
//   -> "gen/proto"
func importPath(f *descriptor.FileDescriptorProto) string {
	gopkg := f.Options.GetGoPackage()
	if idx := strings.LastIndex(gopkg, "/"); idx >= 0 {
		return gopkg[:idx]
	}
	return filepath.Dir(f.GetName())
}

// streamInterface constructs a stream interface name
// from the given method and service name, appending the
// stream type as a suffix.
//
//  Ex:
//  service Foo {
//    Bar(stream BarRequest) returns (BarResponse)
//
//  -> FooBarClientStream
func streamInterface(service, method, stream string) string {
	return strings.Join([]string{service, method, stream}, "")
}
