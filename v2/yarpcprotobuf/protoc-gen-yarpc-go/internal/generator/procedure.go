package generator

const _proceduresTemplate = `
{{define "procedures" -}}
{{with .File}}

{{range .Services -}}
  {{$svc := .Name}}

  {{/* Procedure construction */}}

  // Build{{$svc}}Procedures constructs the YARPC procedures for the {{$svc}} service.
  func Build{{$svc}}Procedures(s {{$svc}}Server) []yarpc.Procedure {
    h := &_{{$svc}}Handler{server: s}
    return yarpcprotobuf.Procedures(
      yarpcprotobuf.ProceduresParams{
        Service: {{printf "%q" $svc}},
        Unary: []yarpcprotobuf.UnaryProceduresParams{
          {{range unaryMethods . -}}
          {
            MethodName: {{printf "%q" .Name}},
            Handler: yarpcprotobuf.NewUnaryHandler{
              yarpcprotobuf.UnaryHandlerParams{
                Handle: h.{{.Name}},
                Create: new{{$svc}}{{.Name}}Request(),
              },
            },
          },
          {{end -}}
        },
        Stream: []yarpcprotobuf.StreamProceduresParams{
          {{range streamMethods . -}}
          {
            MethodName: {{printf "%q" .Name}},
            Handler: yarpcprotobuf.NewStreamHandler{
              yarpcprotobuf.StreamHandlerParams{
                Handle: h.{{.Name}},
              },
            },
          },
          {{end -}}
        },
      },
    )
  }
{{end -}}

{{end -}}{{end -}}
`
