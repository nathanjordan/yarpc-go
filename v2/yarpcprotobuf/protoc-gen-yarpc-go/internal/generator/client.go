package generator

const _clientTemplate = `
{{define "client" -}}
{{with .File}}

{{range .Services -}}
  {{$svc := .Name}}

  {{/* Client interface */}}

  type {{$svc}}Client interface {
    {{range .Methods -}}
      {{if and (not .ClientStreaming) (not .ServerStreaming) -}}
        {{.Name}}(context.Context, {{.Request.Name}}, ...yarpc.CallOption) ({{.Response.Name}}, error)
      {{else -}}
        {{.Name}}(context.Context, {{if not .ClientStreaming}}{{.Request.Name}},{{end -}} ...yarpc.CallOption) ({{$svc}}{{.Response.Name}}Client, error)
      {{end -}}
    {{end -}}
  }

  {{/* Stream client interfaces */}}

  {{range .Methods -}}
    {{if or .ClientStreaming .ServerStreaming}}
    type {{.Name}}Client interface {
      Context() context.Context

    {{if .ClientStreaming -}}
      Send({{.Request.Name}}, ...yarpc.StreamOption) error
    {{end -}}

    {{if .ServerStreaming -}}
      Recv(...yarpc.StreamOption) ({{.Response.Name}}, error)
      CloseSend(...yarpc.StreamOption) error
    {{end -}}

    {{if and .ClientStreaming (not .ServerStreaming) -}}
      CloseAndRecv(...yarpc.StreamOption) ({{.Response.Name}}, error)
    {{end -}}
    }
    {{end -}}
  {{end -}}
{{end -}}

{{end -}}{{end -}}
`
