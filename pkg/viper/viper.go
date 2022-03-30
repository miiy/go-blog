package viper

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"path"
	"strings"
)

func NewViper(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(strings.TrimRight(path.Base(filename), path.Ext(filename))) // name of config file (without extension)
	v.SetConfigType(strings.TrimLeft(path.Ext(filename), "."))           // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath(path.Dir(filename))                                         // optionally look for config in the working directory
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Fatal error config file: %w \n", err)
	}

	return v, nil
}

var ProviderSet = wire.NewSet(NewViper)