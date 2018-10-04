package generator

import (
	"fmt"
	"strings"
)

// Imports is used to maintain the set of aliases and map them back
// to the package path.
type Imports map[string]string

// newImports returns a new Imports struct with the given paths.
func newImports(paths ...string) Imports {
	imports := Imports{}
	for _, path := range paths {
		imports.Add(path)
	}
	return imports
}

// Add adds the path to the imports map, initially using the base directory
// as the package alias. If this alias is already in use, we continue to
// prepend the remaining filepath elements until we have receive a unique
// alias. If all of the path elements are exhausted, a '_' is continually used
// until we create a unique alias.
//
//  Ex:
//   imports := NewImports("json")
//   imports.Add("encoding/json") -> "encodingjson"
//   imports.Add("encodingjson")  -> "_encodingjson"
func (imp Imports) Add(path string) string {
	if path == "" || path == "." {
		return ""
	}
	if alias, ok := imp[path]; ok {
		return alias
	}

	elems := strings.Split(path, "/")
	var alias string
	for i := 1; i <= len(elems); i++ {
		alias = strings.Join(elems[len(elems)-i:], "")
		if !imp.contains(alias) {
			break
		}
	}
	for imp.contains(alias) {
		alias = fmt.Sprintf("_%s", alias)
	}
	imp[path] = alias
	return alias
}

// contains determines whether the given alias is already
// registered in the import map.
func (imp Imports) contains(alias string) bool {
	for _, i := range imp {
		if i == alias {
			return true
		}
	}
	return false
}
