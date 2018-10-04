package generator

const _clientTemplate = `
{{define "client" -}}
{{with .File}}
{{$gopkg := .Package.GoPackage}}

{{range .Services -}}
  {{$svc := .Name}}

  {{/* Client interface */}}

  // {{$svc}}Client is the {{$svc}} service's client interface.
  type {{$svc}}Client interface {
    {{range .Methods -}}
      {{if or .ClientStreaming .ServerStreaming -}}
        {{.Name}}(
          context.Context,
          {{if not .ClientStreaming -}}
            *{{goType .Request $gopkg}},
          {{end -}}
          ...yarpc.CallOption,
        ) ({{.ClientStream}}, error)
      {{else -}}
        {{.Name}}(
          context.Context,
          *{{goType .Request $gopkg}},
          ...yarpc.CallOption,
        ) (*{{goType .Response $gopkg}}, error)
      {{end -}}
    {{end -}}
  }

  {{/* Client construction */}}

  // New{{$svc}}Client builds a new YARPC client for the {{$svc}} service.
  func New{{$svc}}Client(c yarpc.Client, opts ...yarpcprotobuf.ClientOption) {{$svc}}Client {
    return &_{{$svc}}Caller{stream: yarpcprotobuf.NewStreamClient(c, {{printf "%q" $svc}}, opts...)}
  }

{{end -}}

{{end -}}{{end -}}
`
