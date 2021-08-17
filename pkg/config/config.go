package config

import (
	"github.com/miiy/go-blog/pkg/database"
	"github.com/miiy/go-blog/pkg/jwtauth"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type AppOptions struct {
	Env string
	Locale string
	Debug bool
}

type Config struct {
	App AppOptions            `yaml:"app"`
	Database database.Options `yaml:"database"`
	Jwt jwtauth.Options       `yaml:"jwt"`
}

func NewConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config *Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return config, nil
}
