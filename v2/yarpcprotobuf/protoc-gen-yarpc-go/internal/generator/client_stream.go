package generator

const _clientStreamTemplate = `
{{define "clientStream" -}}
{{with .File}}
{{$gopkg := .Package.GoPackage}}

{{range .Services -}}
  {{$svc := .Name}}

  {{/* Client stream interfaces */}}

  {{range .Methods -}}
    {{if or .ClientStreaming .ServerStreaming}}
    // {{.ClientStream}} is a streaming interface used in the {{$svc}}}Client interface.
    type {{.ClientStream}} interface {
      Context() context.Context
    {{if .ClientStreaming -}}
      Send(*{{goType .Request $gopkg}}, ...yarpc.StreamOption) error
    {{end -}}
    {{if .ServerStreaming -}}
      Recv(...yarpc.StreamOption) (*{{goType .Response $gopkg}}, error)
      CloseSend(...yarpc.StreamOption) error
    {{end -}}
    {{if and .ClientStreaming (not .ServerStreaming) -}}
      CloseAndRecv(...yarpc.StreamOption) (*{{goType .Response $gopkg}}, error)
    {{end -}}
    }
    {{end -}}
  {{end -}}

  {{/* Client stream implementations */}}

  {{range .Methods -}}
    {{if or .ClientStreaming .ServerStreaming}}
    type _{{.ClientStream}} struct {
      stream *yarpcprotobuf.ClientStream
    }

    func (c *_{{.ClientStream}}) Context() context.Context {
      return c.stream.Context()
    }

    {{if .ClientStreaming}}
    func (c *_{{.ClientStream}}) Send(req *{{goType .Request $gopkg}}, opts ...yarpc.StreamOption) error {
      return c.stream.Send(req, opts...)
    }
    {{end -}}

    {{if .ServerStreaming}}
    func (c *_{{.ClientStream}}) Recv(opts ...yarpc.StreamOption) (*{{goType .Response $gopkg}}, error) {
      msg, err := c.stream.Receive(new{{.ResponseType}}, opts...)
      if err != nil {
        return nil, err
      }
      res, ok := msg.(*{{.ResponseType}})
      if !ok {
        return nil, yarpcprotobuf.CastError(_empty{{.ResponseType}}, msg)
      }
      return res, nil
    }

    func (c *_{{.ClientStream}}) CloseSend(opts ...yarpc.StreamOption) error {
      return c.stream.Close(opts...)
    }
    {{end -}}

    {{if and .ClientStreaming (not .ServerStreaming)}}
    func (c *_{{.ClientStream}}) CloseAndRecv(opts ...yarpc.StreamOption) (*{{goType .Request $gopkg}}, error) {
      if err := c.stream.Close(opts...); err != nil {
        return nil, err
      }
      msg, err := c.stream.Receive(new{{.ResponseType}}, opts...)
      if err != nil {
        return nil, err
      }
      res, ok := msg.(*{{goType .Response $gopkg}})
      if !ok {
        return nil, protobuf.CastError(_empty{{.ResponseType}}, msg)
      }
      return res, err
    }
    {{end -}}

    {{end -}}
  {{end -}}

{{end -}}

{{end -}}{{end -}}
`
