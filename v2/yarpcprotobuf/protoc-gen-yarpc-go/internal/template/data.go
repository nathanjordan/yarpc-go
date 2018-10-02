package template

// Data holds the information referenced in the
// protoc-gen-yarpc-go templates.
type Data struct {
	Filename string
	Imports  []string
	Services []Service
}

// Service represents a Protobuf service definition.
type Service struct {
	Name          string
	UnaryMethods  []Method
	StreamMethods []StreamMethod
}

// Method represents a standard RPC method.
type Method struct {
	Name     string
	Request  string
	Response string
}

// StreamMethod represents an RPC method with
// either client-side or server-side streaming.
type StreamMethod struct {
	Method

	ClientSide bool
	ServerSide bool
}
