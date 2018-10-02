package generator

const _serverTemplate = `
{{define "server" -}}
{{with .File}}

{{range .Services -}}
{{$svc := .Name}}

  {{/* Service interface */}}

  type {{$svc}}Server interface {

  {{range .Methods -}}
    {{if (not .ClientStreaming) and (not .ServerStreaming) -}}
      {{.Name}}(context.Context, {{.Request}}) ({{.Response}}, error)
    {{else -}}
      {{.Name}}({{if not .ClientStreaming}}{{.Request}},{{end -}} {{$svc}}{{.Name}}Server) ({{if not .ServerStreaming}}{{.Response}},{{end -}} error)
    {{end -}}
  {{end -}}
  }

{{/* Stream server interfaces */}}

  {{range .Methods -}}
    {{if .ClientStreaming or .ServerStreaming}}
    type {{.Name}}Server interface {
      Context() context.Context

    {{if .ClientStreaming -}}
      Recv(...yarpc.StreamOption) ({{.Request}}, error)
    {{end -}}

    {{if .ServerStreaming -}}
      Send({{.Response}}, ...yarpc.StreamOption) error
    {{end -}}
    }
	{{end -}}
  {{end -}}
{{end -}}{{end -}}{{end -}}
`
