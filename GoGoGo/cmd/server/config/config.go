package config

import (
	"errors"
)

type Config struct {
	PublicConfig  PublicConfig
	PrivateConfig PrivateConfig
}

type PublicConfig struct {
	PublicKey string
}

type PrivateConfig struct {
	SecretKey string
}

var (
	envs = map[string]PublicConfig{
		"local": {
			PublicKey: "localAdmin",
		},
		"dev": {
			PublicKey: "devAdmin",
		},
		"prod": {
			PublicKey: "prodAdmin",
		},
	}
)

func NewConfig(env string) (Config, error) {

	publicConfig, exists := envs[env]
	if !exists {
		return Config{}, errors.New("env doest not exists")
	}

	return Config{
		PublicConfig: publicConfig,
	}, nil

}
