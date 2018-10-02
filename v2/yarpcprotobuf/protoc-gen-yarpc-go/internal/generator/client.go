package generator

const _clientTemplate = `
{{define "client" -}}
{{with .File}}

{{range .Services -}}
{{$svc := .Name}}

  {{/* Client interface */}}

  type {{$svc}}Client interface {

  {{range .Methods -}}
    {{if (not .ClientStreaming) and (not .ServerStreaming) -}}
      {{.Name}}(context.Context, {{.Request}}, ...yarpc.CallOption) ({{.Response}}, error)
    {{else -}}
      {{.Name}}(context.Context, {{if not .ClientStreaming}}{{.Request}},{{end -}} ...yarpc.CallOption) ({{$svc}}{{.Response}}Client, error)
    {{end -}}
  {{end -}}
  }

  {{/* Stream client interfaces */}}

  {{range .Methods -}}
    {{if .ClientStreaming or .ServerStreaming}}
    type {{.Name}}Client interface {
      Context() context.Context

    {{if .ClientStreaming -}}
      Send({{.Request}}, ...yarpc.StreamOption) error
    {{end -}}

    {{if .ServerStreaming -}}
      Recv(...yarpc.StreamOption) ({{.Response}}, error)
      CloseSend(...yarpc.StreamOption) error
    {{end -}}

    {{if .ClientStreaming and (not .ServerStreaming) -}}
      CloseAndRecv(...yarpc.StreamOption) ({{.Response}}, error)
    {{end -}}
    }
    {{end -}}
  {{end -}}
{{end -}}{{end -}}{{end -}}
`
