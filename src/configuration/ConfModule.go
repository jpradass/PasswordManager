package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Configuration ...
// Configuration file structure to read yaml configuration file
type Configuration struct {
	Host       string `yaml:"host"`
	User       string `yaml:"user"`
	DB         string `yaml:"db"`
	DbPass     string `yaml:"db_pass"`
	Password   string `yaml:"pass"`
	Cost       int    `yaml:"cost"`
	Collection string `yaml:"collection"`
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

//SetPassword ...
//Set new master password through init command
func (conf *Configuration) SetPassword(pwd string) error {
	conf.Password = pwd
	out, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("configuration/conf.yaml", out, 0644)
}
