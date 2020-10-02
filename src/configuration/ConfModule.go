package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Configuration ...
// Configuration file structure to read yaml configuration file
type Configuration struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
	Cost     int    `yaml:"cost"`
}

// LoadConfiguration ...
// Loads configuration to database connection
func (conf *Configuration) LoadConfiguration() error {
	file, err := ioutil.ReadFile("configuration/conf.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, conf)
	if err != nil {
		return err
	}

	return nil
}
