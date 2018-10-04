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

{{end -}}

{{end -}}{{end -}}
`
