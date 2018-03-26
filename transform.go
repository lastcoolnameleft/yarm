package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/ghodss/yaml"
)

// Transform template + values => string
func Transform(content []byte, values Values) (tpl bytes.Buffer, err error) {
	println("----Content----")
	println(string(content))
	println("----Values----")
	fmt.Printf("%+v\n", values)
	tmpl, err := template.New("test").Parse(string(content))
	err = tmpl.Execute(&tpl, values)
	println("----Post Template----")
	println(tpl.String())
	return
}

// ConvertToJSON Converts bytes to JSON
func ConvertToJSON(content bytes.Buffer) (res string, err error) {
	json, err := yaml.YAMLToJSON(content.Bytes())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	res = string(json)
	return
}
