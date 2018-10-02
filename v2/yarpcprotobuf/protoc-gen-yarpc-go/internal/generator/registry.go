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
	Files    map[string]*File
	Messages map[string]*Message
	Imports  Imports
}

func newRegistry(req *plugin.CodeGeneratorRequest) (*registry, error) {
	r := &registry{
		Files:    make(map[string]*File),
		Messages: make(map[string]*Message),
		Imports:  newImports(),
	}
	return r, r.Load(req)
}

// Load registers all of the Proto types provided in the CodeGeneratorRequest
// with the registry. Note that all CodeGeneratorRequests MUST only request
// to generate a single go package. Otherwise, multiple go packages will
// be deposited in the same directory and will therefore not compile. This
// is synonymous to the protoc-gen-go plugin, so we implement the same
// restriction here.
func (r *registry) Load(req *plugin.CodeGeneratorRequest) error {
	for _, f := range req.GetProtoFile() {
		r.loadFile(f)
	}

	var targetPkg string
	for _, name := range req.FileToGenerate {
		target := r.Files[name]
		if target == nil {
			return fmt.Errorf("file target %q was not registered", name)
		}
		pkg := target.Package.Name
		if targetPkg == "" {
			targetPkg = pkg
		}
		if targetPkg != pkg {
			return fmt.Errorf("cannot generate multiple go packages: found %q and %q", targetPkg, pkg)
		}
		for _, s := range target.GetService() {
			if err := r.loadService(target, s); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetFile returns the File that corresponds to the
// given name.
func (r *registry) GetFile(name string) (*Message, error) {
	m, ok := r.File[name]
	if !ok {
		return nil, fmt.Errorf("file %q was not found", name)
	}
	return m, nil
}


// GetMessage returns the Message that corresponds to the
// given name. This method expects the input to be formed
// as an input or output type, such as .foo.Bar.
func (r *registry) GetMessage(name string) (*Message, error) {
	// All input and output types are represented as
	// .$(Package).$(Message), so we explicitly trim
	// the leading '.' prefix.
	msg := strings.TrimPrefix(name, ".")
	m, ok := r.Messages[msg]
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
	pkg := newPackage(f)
	file := &File{
		Package: pkg,
	}
	r.Files[file.GetName()] = file
	for _, m := range f.GetMessageType() {
		r.loadMessage(m, pkg)
	}
}

func (r *registry) loadMessage(f *File, m *descriptor.DescriptorProto) {
	msg := &Message{
		Name:    m.GetName(),
		Package: f.Package,
	}
	f.Messages = append(file.Messages, msg)
	r.Messages[msg.FQN()] = msg

	for _, n := range m.GetNestedType() {
		r.loadMessage(f, m.GetNestedType())
	}
}

func (r *registry) loadService(f *File, s *descriptor.ServiceDescriptorProto) error {
	svc := &Service{
		Name:    s.GetName(),
		Package: f.Package,
	}
	for _, m := range s.GetMethod() {
		method, err := r.newMethod(m)
		if err != nil {
			return err
		}
		svc.Methods = append(svc.Methods, method)
	}
	f.Services = append(f.Services, svc)
	return nil
}

func (r *registry) newMethod(m *descriptor.MethodDescriptorProto) (*Method, error) {
	request, err := r.GetMessage(m.GetInputType())
	if err != nil {
		return nil, err
	}
	response, err := r.GetMessage(m.GetOutputType())
	if err != nil {
		return nil, err
	}
	return &Method{
		Name:     m.GetName(),
		Request:  request,
		Response: response,
	}, nil
}

func (r *registry) newPackage(f *descriptor.FileDescriptorProto) *Package {
	importPath := importPath(f)
	return &Package{
		Name:       f.GetPackage(),
		GoPackage:  goPackage(f),
		ImportPath: importPath,
		Alias:      r.Imports.Add(importPath),
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
func (r *registry) importPath(f *descriptor.FileDescriptorProto) string {
	gopkg := f.Options.GetGoPackage()
	if idx := strings.LastIndex(gopkg, "/"); idx >= 0 {
		return gopkg[:idx]
	}
	return filepath.Dir(f.Name())
}
