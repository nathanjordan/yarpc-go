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
		_serverTemplate,
	),
)

func parseTemplates(templates ...string) (*template.Template, error) {
	t := template.New(_plugin).Funcs(
		template.FuncMap{
			"goType": goType,
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
//
//  Ex:
//  *Foo
//  *foo.Foo
func goType(m *Message, pkg string) string {
	if m.Package.GoPackage == pkg {
		return m.Name
	}
	return fmt.Sprintf("%s.%s", m.Package.Alias, m.Name)
}
