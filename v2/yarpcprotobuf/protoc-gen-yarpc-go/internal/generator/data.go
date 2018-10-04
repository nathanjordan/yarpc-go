package generator

import (
	"fmt"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// Data holds the information required for
// the protoc-gen-yarpc-go plugin.
type Data struct {
	File    *File
	Imports Imports
}

// File represents a Protobuf file descriptor.
type File struct {
	descriptor *descriptor.FileDescriptorProto

	Name     string
	Package  *Package
	Services []*Service
}

// Package holds information with respect
// to a Proto type's package.
type Package struct {
	alias string
	name  string

	GoPackage string
}

// fqn returns the fully-qualified name for
// the given name based on this package.
//
//  Ex:
//  p := &Package{name: "foo.bar"}
//  p.fqn("Baz") -> "foo.bar.Baz"
func (p *Package) fqn(name string) string {
	return fmt.Sprintf("%s.%s", p.name, name)
}

// Service represents a Protobuf service definition.
type Service struct {
	Name       string
	FQN        string
	Client     string
	ClientImpl string
	FxClient   string
	Server     string
	ServerImpl string
	FxServer   string
	Procedures string
	Methods    []*Method
}

// Method represents a standard RPC method.
type Method struct {
	Name            string
	Request         *Message
	Response        *Message
	ClientStreaming bool
	ServerStreaming bool
	StreamClient    string
	StreamServer    string
	EmptyRequest    string
	EmptyResponse   string
	NewRequest      string
	NewResponse     string
}

// Message represents a Protobuf message definition.
type Message struct {
	Name    string
	Package *Package
}
