package generator

import (
	"bytes"
	"fmt"
	"text/template"
)

const (
	_clientStreamInterface = "ClientStream"
	_serverStreamInterface = "ServerStream"
)

var _tmpl = template.Must(
	parseTemplates(
		_baseTemplate,
		_clientTemplate,
		_serverTemplate,
	),
)

func parseTemplates(templates ...string) (*template.Template, error) {
	t := template.New(_plugin).Funcs(
		template.FuncMap{
			"goType":       goType,
			"clientStreamInterface": clientStreamInterface,
			"serverStreamInterface": serverStreamInterface,
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

// clientStreamInterface returns the name of the client stream
// interface for the given method.
func clientStreamInterface(m *Method) string {
	return fmt.Sprintf("%s%s%s", m.Service, m.Name, _clientStreamInterface)
}

// serverStreamInterface returns the name of the server stream
// interface for the given method.
func serverStreamInterface(m *Method) string {
	return fmt.Sprintf("%s%s%s", m.Service, m.Name, _serverStreamInterface)
}

// goType returns a go type name for the message type.
// It prefixes the type name with the package's alias
// if it does not belong to the same package.
func goType(m *Message, pkg string) string {
	name := m.Name
	if m.Package.GoPackage != pkg {
		name = fmt.Sprintf("%s.%s", m.Package.Alias, m.Name)
	}
	return fmt.Sprintf("*%s", name)
}
