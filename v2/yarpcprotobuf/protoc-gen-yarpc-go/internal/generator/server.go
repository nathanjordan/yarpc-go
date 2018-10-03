package generator

const _serverTemplate = `
{{define "server" -}}
{{with .File}}
{{$gopkg := .Package.GoPackage}}

{{range .Services -}}
  {{$svc := .Name}}

  {{/* Service interface */}}

  // {{$svc}}Server is the {{$svc}} service's server interface.
  type {{$svc}}Server interface {
    {{range .Methods -}}
      {{if or .ClientStreaming .ServerStreaming -}}
        {{.Name}}(
          {{if not .ClientStreaming -}}
            {{goType .Request $gopkg}},
          {{end -}}
          {{.ServerStream}},
        ) {{if not .ServerStreaming}} ({{.ServerStream}}, error) {{else}} error {{end}}
      {{else -}}
        {{.Name}}(
          context.Context,
          {{goType .Request $gopkg}},
        ) ({{goType .Response $gopkg}}, error)
      {{end -}}
    {{end -}}
  }


  {{/* Stream server interfaces */}}

  {{range .Methods -}}
    {{if or .ClientStreaming .ServerStreaming}}
    // {{.ServerStream}} is a streaming interface used in the {{$svc}}}Server interface.
    type {{.ServerStream}} interface {
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
