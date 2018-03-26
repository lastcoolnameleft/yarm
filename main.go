package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Values represents a collection of YARM values.
type Values map[string]interface{}

func main() {
	filename := os.Args[1]
	resource, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	base, err := ioutil.ReadFile("content/base/arm-base.yaml")
	if err != nil {
		panic(err)
	}

	values, err := ReadValuesFile(os.Args[2])
	if err != nil {
		panic(err)
	}

	resources, err := Transform("resources", resource, values)
	if err != nil {
		panic(err)
	}

	values["resources"] = resources.String()

	res, err := Transform("base", base, values)
	if err != nil {
		panic(err)
	}

	println("----YAML----")
	println(res.String())

	json2, err := ConvertToPrettyJSON(res)
	if err != nil {
		panic(err)
	}

	println("----JSON----")
	fmt.Println(json2)
}
