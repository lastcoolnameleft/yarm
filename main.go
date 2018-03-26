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
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	values, err := ReadValuesFile(os.Args[2])
	if err != nil {
		panic(err)
	}

	res, err := Transform(source, values)
	if err != nil {
		panic(err)
	}

	json, err := ConvertToJSON(res)
	if err != nil {
		panic(err)
	}

	fmt.Println(json)
}
