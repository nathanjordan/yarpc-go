package generator

const _handlerTemplate = `
{{define "handler" -}}
{{with .File}}
{{$gopkg := .Package.GoPackage}}

{{range .Services -}}
  {{$svc := .Name}}
  {{$handler := printf "_%s%s" $svc "Handler"}}

  {{/* Handler implementation */}}

  type {{$handler}} struct {
    server {{$svc}}Server
  }

  {{range .Methods}}
    {{if and .ClientStreaming .ServerStreaming -}}
      func (h *{{$handler}}) {{.Name}}(s *yarpcprotobuf.ServerStream) error {
        h.server.{{.Name}}(&_{{.ServerStream}}{stream: s})
      }
    {{end -}}
  {{end -}}
{{end -}}

{{end -}}{{end -}}
`

//type _KeyValueYARPCHandler struct {
//server KeyValueYARPCServer
//}

//func (h *_KeyValueYARPCHandler) Foo(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
//var request *GetRequest
//var ok bool
//if requestMessage != nil {
//request, ok = requestMessage.(*GetRequest)
//if !ok {
//return nil, protobuf.CastError(emptyKeyValueServiceFooYARPCRequest, requestMessage)
//}
//}
//response, err := h.server.Foo(ctx, request)
//if response == nil {
//return nil, err
//}
//return response, err
//}

//func (h *_KeyValueYARPCHandler) Bar(serverStream *protobuf.ServerStream) error {
//response, err := h.server.Bar(&_KeyValueServiceBarYARPCServer{serverStream: serverStream})
//if err != nil {
//return err
//}
//return serverStream.Send(response)
//}

//func (h *_KeyValueYARPCHandler) Baz(serverStream *protobuf.ServerStream) error {
//requestMessage, err := serverStream.Receive(newKeyValueServiceBazYARPCRequest)
//if requestMessage == nil {
//return err
//}

//request, ok := requestMessage.(*GetRequest)
//if !ok {
//return protobuf.CastError(emptyKeyValueServiceBazYARPCRequest, requestMessage)
//}
//return h.server.Baz(request, &_KeyValueServiceBazYARPCServer{serverStream: serverStream})
//}

//func (h *_KeyValueYARPCHandler) Qux(serverStream *protobuf.ServerStream) error {
//return h.server.Qux(&_KeyValueServiceQuxYARPCServer{serverStream: serverStream})
//}
