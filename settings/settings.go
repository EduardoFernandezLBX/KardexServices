package settings

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed settings.yaml
var settingsFile []byte

type databaseConfig struct {
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

type Settings struct {
	Port string         `yaml:"port"`
	DB   databaseConfig `yaml:"database"`
}

func New() (*Settings, error) {
	var s Settings

	err := yaml.Unmarshal(settingsFile, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
