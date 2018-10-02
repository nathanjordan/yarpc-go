package template

// Client represents the generated client code.
const Client = `
{{define "client" -}}

{{/* Client interfaces */}}
{{range .Services -}}
{{$svc := .Name}}

type {{$svc}}Client interface {

{{range .UnaryMethods -}}
  {{.Name}}(context.Context, {{.Request}}, ...yarpc.CallOption) ({{.Response}}, error)
{{end -}}

{{range .StreamMethods -}}
  {{.Name}}(context.Context, {{if not .ClientSide}}{{.Request}},{{end -}} ...yarpc.CallOption) ({{$svc}}{{.Response}}Client, error)
{{end -}}

}
{{end -}}

{{/* Stream client interfaces */}}
{{range .Services -}}

{{range .StreamMethods -}}
type {{.Name}}Client interface {
  Context() context.Context

  {{if .ClientSide -}}
  Send({{.Request}}, ...yarpc.StreamOption) error
  {{end -}}

  {{if .ServerSide -}}
  Recv(...yarpc.StreamOption) ({{.Response}}, error)
  CloseSend(...yarpc.StreamOption) error
  {{end -}}

  {{if (not .ClientSide) and (not .ServerSide) -}}
  CloseAndRecv(...yarpc.StreamOption) ({{.Response}}, error)
  {{end -}}
}
{{end -}}{{end -}}{{end -}}
`
