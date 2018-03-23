package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/ghodss/yaml"
)

// Transform template + values => string
func Transform(content []byte, values Values) (res string, err error) {
	tmpl, err := template.New("test").Parse(string(content))
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, values)

	j2, err := yaml.YAMLToJSON(tpl.Bytes())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	res = string(j2)
	return
}
