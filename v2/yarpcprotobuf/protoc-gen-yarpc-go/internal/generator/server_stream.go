package generator

const _serverStreamTemplate = `
{{define "serverStream" -}}
{{with .File}}
{{$gopkg := .Package.GoPackage}}

{{range .Services -}}
  {{$svc := .Name}}

  {{/* Server stream interfaces */}}

  {{range .Methods -}}
    {{if or .ClientStreaming .ServerStreaming}}
    // {{.ServerStream}} is a streaming interface used in the {{$svc}}}Server interface.
    type {{.ServerStream}} interface {
      Context() context.Context
    {{if .ClientStreaming -}}
      Recv(...yarpc.StreamOption) (*{{goType .Request $gopkg}}, error)
    {{end -}}
    {{if .ServerStreaming -}}
      Send(*{{goType .Response $gopkg}}, ...yarpc.StreamOption) error
    {{end -}}
    }
    {{end -}}
  {{end -}}

  {{/* Server stream implementations */}}

  {{range .Methods -}}
    {{if or .ClientStreaming .ServerStreaming}}
    type _{{.ServerStream}} struct {
      stream *yarpcprotobuf.ServerStream
    }

    func (s *_{{.ServerStream}}) Context() context.Context {
      return s.stream.Context()
    }

    {{if .ClientStreaming}}
    func (s *_{{.ServerStream}}) Recv(opts ...yarpc.StreamOption) (*{{goType .Request $gopkg}}, error) {
      msg, err := s.stream.Receive(new{{.RequestType}}, opts...)
      if err != nil {
        return nil, err
      }
      req, ok := msg.(*{{.RequestType}})
      if !ok {
        return nil, yarpcprotobuf.CastError(_empty{{.RequestType}}, msg)
      }
      return req, nil
    }
    {{end -}}

    {{if .ServerStreaming}}
    func (s *_{{.ServerStream}}) Send(res *{{goType .Response $gopkg}}, opts ...yarpc.StreamOption) error {
      return s.stream.Send(res, opts...)
    }
    {{end -}}

    {{end -}}
  {{end -}}

{{end -}}

{{end -}}{{end -}}
`
