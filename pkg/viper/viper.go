package viper

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func NewViper(name, cType string, paths []string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(name) // name of config file (without extension)
	v.SetConfigType(cType) // REQUIRED if the config file does not have the extension in the name
	for _, path := range paths {
		v.AddConfigPath(path)               // optionally look for config in the working directory
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

var ProviderSet = wire.NewSet(NewViper)