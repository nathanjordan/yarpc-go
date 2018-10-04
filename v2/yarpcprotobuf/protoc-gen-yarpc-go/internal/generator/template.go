package generator

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"go.uber.org/yarpc/v2/yarpcprotobuf/protoc-gen-yarpc-go/internal/templatedata"
)

var _tmpl = template.Must(
	parseTemplates(
		templatedata.MustAsset("internal/templatedata/base.tmpl"),
		templatedata.MustAsset("internal/templatedata/client.tmpl"),
		templatedata.MustAsset("internal/templatedata/client_impl.tmpl"),
		templatedata.MustAsset("internal/templatedata/client_stream.tmpl"),
		templatedata.MustAsset("internal/templatedata/fx.tmpl"),
		templatedata.MustAsset("internal/templatedata/parameters.tmpl"),
		templatedata.MustAsset("internal/templatedata/server.tmpl"),
		templatedata.MustAsset("internal/templatedata/server_impl.tmpl"),
		templatedata.MustAsset("internal/templatedata/server_stream.tmpl"),
	),
)

func parseTemplates(templates ...[]byte) (*template.Template, error) {
	t := template.New(_plugin).Funcs(
		template.FuncMap{
			"goType":        goType,
			"unaryMethods":  unaryMethods,
			"streamMethods": streamMethods,
		},
	)
	for _, tmpl := range templates {
		_, err := t.Parse(string(tmpl))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func execTemplate(data *Data) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := _tmpl.Execute(buffer, data); err != nil {
		return nil, err
	}
	os.Stderr.WriteString(buffer.String())
	return buffer.Bytes(), nil
}

// goType returns a go type name for the message type.
// It prefixes the type name with the package's alias
// if it does not belong to the same package.
func goType(m *Message, pkg string) string {
	if m.Package.GoPackage != pkg && m.Package.alias != "" {
		return fmt.Sprintf("%s.%s", m.Package.alias, m.Name)
	}
	return m.Name
}

func unaryMethods(s *Service) []*Method {
	methods := make([]*Method, 0, len(s.Methods))
	for _, m := range s.Methods {
		if !m.ClientStreaming && !m.ServerStreaming {
			methods = append(methods, m)
		}
	}
	return methods
}

func streamMethods(s *Service) []*Method {
	methods := make([]*Method, 0, len(s.Methods))
	for _, m := range s.Methods {
		if m.ClientStreaming || m.ServerStreaming {
			methods = append(methods, m)
		}
	}
	return methods
}
