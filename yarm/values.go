package yarm

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// https://github.com/kubernetes/helm/blob/e8d80729ac20f402a80c5107b8b2513008de13fc/pkg/chartutil/values.go#L117

// ReadValues will parse YAML byte data into a Values.
func ReadValues(data []byte) (vals Values, err error) {
	err = yaml.Unmarshal(data, &vals)
	if len(vals) == 0 {
		vals = Values{}
	}
	return
}

// ReadValuesFile will parse a YAML file into a map of values.
func ReadValuesFile(filename string) (Values, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return ReadValues(data)
}
