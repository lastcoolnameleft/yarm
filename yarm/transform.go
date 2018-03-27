package yarm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/ghodss/yaml"
)

// Transform template + values => string
func Transform(name string, content []byte, values Values) (tpl bytes.Buffer, err error) {
	//	println("----Content----")
	//println(string(content))
	//	println("----Values----")
	//	fmt.Printf("%+v\n", values)
	tmpl, err := template.New(name).Parse(string(content))
	err = tmpl.Execute(&tpl, values)
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

// ConvertToPrettyJSON Converts bytes to JSON
func ConvertToPrettyJSON(content bytes.Buffer) (res string, err error) {
	// Yuck.  This is a pretty ugly process.  Can/will clean up and maybe open PR
	// to https://github.com/ghodss/yaml to support pretty printing
	j, err := yaml.YAMLToJSON(content.Bytes())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(j, &dat); err != nil {
		panic(err)
	}

	j2, err := json.MarshalIndent(dat, "", "    ")
	if err != nil {
		panic(err)
	}

	res = string(j2)
	return
}
