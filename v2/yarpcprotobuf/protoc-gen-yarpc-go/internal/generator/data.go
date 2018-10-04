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
	*descriptor.FileDescriptorProto

	Name     string
	Package  *Package
	Services []*Service
}

// Package holds information with respect
// to a Proto type's package.
type Package struct {
	Name      string
	GoPackage string
	Alias     string
}

// Service represents a Protobuf service definition.
type Service struct {
	Name    string
	Package *Package
	Methods []*Method

	FQN    string
	Client string
	Server string
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

// key returns the unique key for the given message type.
// This is used to uniquely represent the message type so
// that it can be referenced throughout the code generation
// process.
func (m *Message) key() string {
	return fmt.Sprintf("%s.%s", m.Package.Name, m.Name)
}
