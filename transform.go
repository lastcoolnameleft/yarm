package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/ghodss/yaml"
)

// Transform template + values => string
func Transform(content []byte, values Values) (res string, err error) {
	println("----Content----")
	println(string(content))
	println("----Values----")
	fmt.Printf("%+v\n", values)
	tmpl, err := template.New("test").Parse(string(content))
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, values)
	println("----Post Template----")
	println(tpl.String())

	j2, err := yaml.YAMLToJSON(tpl.Bytes())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	res = string(j2)
	return
}
