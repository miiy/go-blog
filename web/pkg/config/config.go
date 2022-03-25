package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type AppOptions struct {
	Env             string `yaml:"env"`
	Debug           bool   `yaml:"debug"`
	Locale          string `yaml:"locale"`
	Addr            string `yaml:"addr"`
	Name            string `yaml:"name"`
	Url             string `yaml:"url"`
	FooterCopyright string `yaml:"footer-copyright"`
}

type Config struct {
	App AppOptions `yaml:"app"`
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
