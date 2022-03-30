package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"path"
	"strings"
)

type Config struct {
	App App `yaml:"app"`
}

type App struct {
	Env             string `yaml:"env"`
	Debug           bool   `yaml:"debug"`
	Locale          string `yaml:"locale"`
	Addr            string `yaml:"addr"`
	Name            string `yaml:"name"`
	Url             string `yaml:"url"`
	FooterCopyright string `yaml:"footer-copyright"`
}

var config *Config
var v *viper.Viper

func NewConfig(filename string) (*Config, error) {
	v = viper.New()
	v.SetConfigName(strings.TrimRight(path.Base(filename), path.Ext(filename))) // name of config file (without extension)
	v.SetConfigType(strings.TrimLeft(path.Ext(filename), "."))           // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath(path.Dir(filename))                                         // optionally look for config in the working directory
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Fatal error config file: %w \n", err)
	}

	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v \n", err)
	}

	return config, nil
}

var ProviderSet = wire.NewSet(NewConfig)

func GetString(key string) string {
	return v.GetString(key)
}