package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadConfiguration ...
// Loads configuration to database connection
func (conf *Configuration) LoadConfiguration() *Configuration {
	file, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		return nil
	}

	err = yaml.Unmarshal(file, conf)
	if err != nil {
		return nil
	}

	return conf
}
