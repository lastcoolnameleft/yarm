package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Values represents a collection of YARM values.
type Values map[string]interface{}

func main() {
	resourceFile := os.Args[1]
	valuesFile := os.Args[2]

	values, err := ReadValuesFile(valuesFile)
	if err != nil {
		panic(err)
	}

	resourcesString := ""

	resources := values["resources"].(map[string]interface{})
	for resourceType := range resources {
		resourceTypeArr := resources[resourceType].([]interface{})
		for id := range resourceTypeArr {
			var vnet = resourceTypeArr[id].(map[string]interface{})
			resource2 := transformResource(resourceFile, vnet)
			resourcesString = resourcesString + resource2
		}
	}
	values["resources"] = resourcesString

	json2 := transformMaster(values)
	println("----JSON----")
	fmt.Println(json2)

}

func transformResource(resourceFile string, values Values) string {
	resource, err := ioutil.ReadFile(resourceFile)
	if err != nil {
		panic(err)
	}

	resources, err := Transform("resources", resource, values)
	if err != nil {
		panic(err)
	}

	return resources.String()
}

func transformMaster(values Values) string {
	base, err := ioutil.ReadFile("content/base/arm-base.yaml")
	if err != nil {
		panic(err)
	}

	res, err := Transform("base", base, values)
	if err != nil {
		panic(err)
	}

	print(res.String())
	json2, err := ConvertToPrettyJSON(res)
	if err != nil {
		panic(err)
	}

	println("----YAML----")
	println(res.String())

	return json2
}
