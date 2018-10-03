package generator

const _callerTemplate = `
{{define "caller" -}}
{{with .File}}
{{$gopkg := .Package.GoPackage}}

{{range .Services -}}
  {{$svc := .Name}}
  {{$caller := printf "_%s%s" $svc "Caller"}}

  {{/* Caller implementation */}}

  type {{$caller}} struct {
    stream yarpcprotobuf.StreamClient
  }

  {{range .Methods}}
    {{if and .ClientStreaming .ServerStreaming -}}
      func (c *{{$caller}}) {{.Name}}(ctx context.Context, opts ...yarpc.CallOption) ({{.ClientStream}}, error) {
        s, err := c.stream.CallStream(ctx, {{printf "%q" .Name}}, opts...)
        if err != nil {
          return nil, err
        }
        return &_{{.ClientStream}}{stream: s}, nil
      }
    {{else if .ClientStreaming -}}
      func (c *{{$caller}}) {{.Name}}(ctx context.Context, opts ...yarpc.CallOption) ({{.ClientStream}}, error) {
        s, err := c.stream.CallStream(ctx, {{printf "%q" .Name}}, opts...)
        if err != nil {
          return err
        }
        return &_{{.ClientStream}}{stream: s}, nil
      }
    {{else if .ServerStreaming -}}
      func (c *{{$caller}}) {{.Name}}(ctx context.Context, req *{{goType .Request $gopkg}}, opts ...yarpc.CallOption) ({{.ClientStream}}, error) {
        s, err := c.stream.CallStream(ctx, {{printf "%q" .Name}}, opts...)
        if err != nil {
          return err
        }
        if err := s.Send(req); err != nil {
          return nil, err
        }
        return &_{{.ClientStream}}{stream: s}, nil
      }
    {{else -}}
      func (c *{{$caller}}) {{.Name}}(ctx context.Context, req *{{goType .Request $gopkg}}, opts ...yarpc.CallOption) (*{{goType .Response $gopkg}}, error) {
        msg, err := c.stream.Call(ctx, {{printf "%q" .Name}}, req, new{{.ResponseType}}, opts...)
        if err != nil {
          return nil, err
        }
        res, ok := msg.(*{{goType .Response $gopkg}})
        if !ok {
          return nil, yarpcprotobuf.CastError(_empty{{.ResponseType}}, res)
        }
        return res, nil
      }
    {{end -}}
  {{end -}}
{{end -}}

{{end -}}{{end -}}
`
