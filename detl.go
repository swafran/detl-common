package detl

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Conf hold configuration for each phase of ETL
type Conf struct {
	Settings map[string]string
	Confs    map[string]map[string]map[string]string
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetConf returns configuration from yaml file
func GetConf(confName string) Conf {
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
