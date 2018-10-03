package generator

import (
	"bytes"
	"fmt"
	"text/template"
)

var _tmpl = template.Must(
	parseTemplates(
		_baseTemplate,
		_clientTemplate,
		_callerTemplate,
		_serverTemplate,
		_handlerTemplate,
		_parametersTemplate,
		_proceduresTemplate,
	),
)

func parseTemplates(templates ...string) (*template.Template, error) {
	t := template.New(_plugin).Funcs(
		template.FuncMap{
			"goType":        goType,
			"unaryMethods":  unaryMethods,
			"streamMethods": streamMethods,
		},
	)
	for _, tmpl := range templates {
		_, err := t.Parse(tmpl)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func execTemplate(data interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := _tmpl.Execute(buffer, data); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// goType returns a go type name for the message type.
// It prefixes the type name with the package's alias
// if it does not belong to the same package.
func goType(m *Message, pkg string) string {
	if m.Package.GoPackage != pkg {
		return fmt.Sprintf("%s.%s", m.Package.Alias, m.Name)
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
