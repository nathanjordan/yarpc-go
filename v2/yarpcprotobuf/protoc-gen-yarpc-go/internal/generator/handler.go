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
        return h.server.{{.Name}}(&_{{.ServerStream}}{stream: s})
      }
    {{else if .ClientStreaming -}}
      func (h *{{$handler}}) {{.Name}}(s *yarpcprotobuf.ServerStream) error {
        res, err := h.server.{{.Name}}(&_{{.ServerStream}}{server: s})
        if err != nil {
          return err
        }
        return s.Send(res)
      }
    {{else if .ServerStreaming -}}
      func (h *{{$handler}}) {{.Name}}(s *yarpcprotobuf.ServerStream) error {
        recv, err := s.Receive(new{{.RequestType}})
        if err != nil {
          return err
        }
        req, _ := recv.(*{{.RequestType}})
        if req == nil {
          return yarpcprotobuf.CastError(_empty{{.RequestType}}, recv)
        }
        return h.server.{{.Name}}(req, &_{{.ServerStream}}{server: s})
      }
    {{else -}}
      func (h *{{$handler}}) {{.Name}}(ctx context.Context, m proto.Message) (proto.Message, error) {
        req, _ := m.(*GetRequest)
        if req == nil {
          return nil, protobuf.CastError(_empty{{.RequestType}}, m)
        }
        return h.server.{{.Name}}(ctx, req)
      }
    {{end -}}
  {{end -}}
{{end -}}

{{end -}}{{end -}}
`
