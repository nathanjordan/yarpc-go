package generator

const _parametersTemplate = `
{{define "parameters" -}}
{{with .File}}
{{$gopkg := .Package.GoPackage}}

{{range .Services -}}
  {{$svc := .Name}}

  {{/* Parameter constructors */}}
  {{range .Methods -}}
    func new{{.RequestType}}()  { return &{{goType .Request $gopkg}}  }
    func new{{.ResponseType}}() { return &{{goType .Response $gopkg}} }
  {{end -}}

  {{/* Empty parameter types */}}
  {{range .Methods -}}
    var (
      _empty{{.RequestType}}  = &{{goType .Request $gopkg}}
      _empty{{.ResponseType}} = &{{goType .Response $gopkg}}
    )
  {{end -}}
{{end -}}

{{end -}}{{end -}}
`
