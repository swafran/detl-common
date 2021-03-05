package settings

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config for each phase of ETL
type Conf struct {
	Settings map[string]string
	Confs    map[string]map[string]map[string]string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetConf returns Conf from yaml file
func GetConf(confName string) Settings {
	var conf Conf

	fileData, err := ioutil.ReadFile(confName + ".yaml")
	err2 := yaml.Unmarshal(fileData, &conf)

	check(err)
	check(err2)

	return conf
}

// GetArbitraryYaml returns arbitrary data from yaml file
func GetArbitraryYaml(settingsName string) map[string]interface{} {
	var settings map[string]interface{}

	fileData, err := ioutil.ReadFile(settingsName + ".yaml")
	err2 := yaml.Unmarshal(fileData, &settings)

	check(err)
	check(err2)

	return settings
}
