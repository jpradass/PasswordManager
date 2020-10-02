package configuration

// Configuration ...
// Configuration file structure to read yaml configuration file
type Configuration struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
}
