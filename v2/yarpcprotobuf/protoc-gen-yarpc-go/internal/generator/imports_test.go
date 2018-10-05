package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImports(t *testing.T) {
	t.Run("void add", func(t *testing.T) {
		assert.Empty(t, Imports{}.Add(""))
		assert.Empty(t, Imports{}.Add("."))
	})
	t.Run("add with import conflicts", func(t *testing.T) {
		imports := Imports{}
		imports.Add("json")
		imports.Add("encoding/json")
		imports.Add("encodingjson")

		expected := map[string]string{
			"json":          "json",
			"encoding/json": "encodingjson",
			"encodingjson":  "_encodingjson",
		}

		assert.Equal(t, len(expected), len(imports))
		assert.Equal(t, expected, toMap(imports))
	})
}

func toMap(i Imports) map[string]string {
	m := make(map[string]string, len(i))
	for p, a := range i {
		m[p] = a
	}
	return m
}
