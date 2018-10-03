package generator

const _clientTemplate = `
{{define "client" -}}
{{with .File}}
{{$gopkg := .Package.GoPackage}}

{{range .Services -}}
  {{$svc := .Name}}

  {{/* Client interface */}}

  type {{$svc}}Client interface {
    {{range .Methods -}}
      {{if and (not .ClientStreaming) (not .ServerStreaming) -}}
        {{.Name}}(
          context.Context,
          {{goType .Request $gopkg}},
          ...yarpc.CallOption,
        ) ({{goType .Response $gopkg}}, error)
      {{else -}}
        {{.Name}}(
          context.Context,
          {{if not .ClientStreaming -}}
            {{goType .Request $gopkg}},
          {{end -}}
          ...yarpc.CallOption,
        ) ({{.ClientStream}}, error)
      {{end -}}
    {{end -}}
  }

  {{/* Stream client interfaces */}}

  {{range .Methods -}}
    {{if or .ClientStreaming .ServerStreaming}}
    type {{.ClientStream}} interface {
      Context() context.Context
    {{if .ClientStreaming -}}
      Send({{goType .Request $gopkg}}, ...yarpc.StreamOption) error
    {{end -}}
    {{if .ServerStreaming -}}
      Recv(...yarpc.StreamOption) ({{goType .Response $gopkg}}, error)
      CloseSend(...yarpc.StreamOption) error
    {{end -}}
    {{if and .ClientStreaming (not .ServerStreaming) -}}
      CloseAndRecv(...yarpc.StreamOption) ({{goType .Response $gopkg}}, error)
    {{end -}}
    }
    {{end -}}
  {{end -}}
{{end -}}

{{end -}}{{end -}}
`
