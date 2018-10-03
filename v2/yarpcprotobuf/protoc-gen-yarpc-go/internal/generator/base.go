package generator

const _baseTemplate = `
// Code generated by protoc-gen-yarpc-go
// source: {{.File.Name}}
// DO NOT EDIT!

package {{.File.Package.GoPackage}}

{{if .File.Services -}}
  import (
    {{range $importPath, $alias := .Imports -}}
      {{$alias}} "{{$importPath}}"
    {{end -}}
  )
{{end -}}

{{template "client" . -}}
{{template "caller" . -}}
{{template "server" . -}}
{{template "handler" . -}}
{{template "procedures" . -}}
{{template "parameters" . -}}
`
