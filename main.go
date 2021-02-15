package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type settings map[string]map[string]string

func main() {
	// // mock
	// doy := make(settings)

	// SetSettings("/Users/fr/proj/detl/transformer/transformer", &doy)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// SetSettings returns nothing
func SetSettings(settingsName string, settings *settings) {
	fileData, err := ioutil.ReadFile(settingsName + ".yaml")
	err2 := yaml.Unmarshal(fileData, settings)

	check(err)
	check(err2)
}
