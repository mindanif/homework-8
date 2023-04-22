package config

import "github.com/kelseyhightower/envconfig"

const envPrefix = "QA"

type Config struct {
	Host   string `split_words:"true" default:"localhost:6000"`
	DbHost string `split_words:"true" default:"localhost"`
	DbPort string `split_words:"true" default:"5432"`
}

func FromEnv() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process(envPrefix, cfg)
	return cfg, err
}
