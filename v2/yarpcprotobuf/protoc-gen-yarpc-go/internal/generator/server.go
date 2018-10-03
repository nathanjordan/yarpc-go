package generator

const _serverTemplate = `
{{define "server" -}}
{{with .File}}
{{$gopkg := .Package.GoPackage}}

{{range .Services -}}
  {{$svc := .Name}}

  {{/* Service interface */}}

  type {{$svc}}Server interface {
    {{range .Methods -}}
      {{if and (not .ClientStreaming) (not .ServerStreaming) -}}
        {{.Name}}(
          context.Context,
          {{goType .Request $gopkg}},
        ) ({{goType .Response $gopkg}}, error)
      {{else -}}
        {{.Name}}(
          {{if not .ClientStreaming -}}
            {{goType .Request $gopkg}},
          {{end -}}
          {{serverStream .}},
        ) {{if not .ServerStreaming}} ({{serverStream .}}, error) {{else}} error {{end}}
      {{end -}}
    {{end -}}
  }


  {{/* Stream server interfaces */}}

  {{range .Methods -}}
    {{if or .ClientStreaming .ServerStreaming}}
    type {{serverStream .}} interface {
      Context() context.Context
    {{if .ClientStreaming -}}
      Recv(...yarpc.StreamOption) ({{goType .Request $gopkg}}, error)
    {{end -}}
    {{if .ServerStreaming -}}
      Send({{goType .Response $gopkg}}, ...yarpc.StreamOption) error
    {{end -}}
    }
    {{end -}}
  {{end -}}
{{end -}}

{{end -}}{{end -}}
`
