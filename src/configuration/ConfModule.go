package configuration

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
	Key        string `yaml:"key"`
}

var expath string

func init() {
	ex, _ := os.Executable()
	expath = filepath.Dir(ex)
}

// LoadConfiguration ...
// Loads configuration to database connection
func (conf *Configuration) LoadConfiguration() error {
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/configuration/conf.yaml", expath))
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, conf)
	if err != nil {
		return err
	}

	return nil
}

//SaveConfiguration ...
//Saves configuration to database connection
func (conf *Configuration) SaveConfiguration() error {
	out, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fmt.Sprintf("%s/configuration/conf.yaml", expath), out, 0644)
}
