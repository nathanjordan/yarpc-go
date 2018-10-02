package generator

import (
	"bytes"
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
	t := template.New(_plugin)
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
