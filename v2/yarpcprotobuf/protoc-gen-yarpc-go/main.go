package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/gogo/protobuf/proto"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	"go.uber.org/yarpc/v2/yarpcprotobuf/protoc-gen-yarpc-go/internal/generator"
)

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("%v\n", err))
		os.Exit(1)
	}
}

func run(input io.Reader, output io.Writer) error {
	req, err := fileToGeneratorRequest(input)
	if err != nil {
		return fmt.Errorf("failed to create CodeGeneratorRequest: %v", err)
	}

	res, err := generator.Generate(req)
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
