package yarm

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// Values represents a collection of YARM values.
type Values map[string]interface{}

// CreateArm will generate the ARM template
func CreateArm(valuesFile, output string) {

	values, err := ReadValuesFile(valuesFile)
	if err != nil {
		panic(err)
	}
	//fmt.Println("values:", values)

	resourcesString := ""

	resources := values["resources"].(map[string]interface{})
	for resourceType := range resources {
		resourceTypeArr := resources[resourceType].([]interface{})
		for id := range resourceTypeArr {
			resourceFile := "content/resources/" + resourceType + ".yarm"
			resource := resourceTypeArr[id].(map[string]interface{})
			resourceString := transformResource(resourceFile, resource)
			resourcesString = resourcesString + resourceString
		}
	}
	values["resources"] = resourcesString

	y := transformMaster(values)
	if output == "yaml" {
		fmt.Println(y.String())
	}
	if output == "json" {
		json2, err := ConvertToPrettyJSON(y)
		if err != nil {
			panic(err)
		}
		fmt.Println(json2)
	}

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

func transformMaster(values Values) bytes.Buffer {
	base, err := ioutil.ReadFile("content/base/arm-base.yaml")
	if err != nil {
		panic(err)
	}

	res, err := Transform("base", base, values)
	if err != nil {
		panic(err)
	}

	return res
}
