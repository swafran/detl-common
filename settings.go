package settings

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Settings for each phase of ETL
type Settings map[string]interface{}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetSettings returns Settings from yaml file
func GetSettings(settingsName string) Settings {
	var settings Settings

	fileData, err := ioutil.ReadFile(settingsName + ".yaml")
	err2 := yaml.Unmarshal(fileData, &settings)

	check(err)
	check(err2)

	return settings
}
