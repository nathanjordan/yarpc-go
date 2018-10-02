package generator

const _serverTemplate = `
{{define "server" -}}

{{/* Service interfaces */}}
{{range .Services -}}
{{$svc := .Name}}

type {{$svc}}Server interface {

{{range .UnaryMethods -}}
  {{.Name}}(context.Context, {{.Request}}) ({{.Response}}, error)
{{end -}}

{{range .StreamMethods -}}
  {{.Name}}({{if not .ClientSide}}{{.Request}},{{end -}} {{$svc}}{{.Name}}Server) ({{if not .ServerSide}}{{.Response}},{{end -}} error)
{{end -}}

}
{{end -}}

{{/* Stream server interfaces */}}
{{range .Services -}}

{{range .StreamMethods -}}
type {{.Name}}Server interface {
  Context() context.Context

  {{if .ClientSide -}}
  Recv(...yarpc.StreamOption) ({{.Request}}, error)
  {{end -}}

  {{if .ServerSide -}}
  Send({{.Response}}, ...yarpc.StreamOption) error
  {{end -}}
}
{{end -}}{{end -}}{{end -}}
`
