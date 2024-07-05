package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Version    string `envconfig:"VERSION"`
	Env        string `envconfig:"ENV"`
	AppEnv     string `envconfig:"APP_ENV"`
	Port       string `envconfig:"PORT"`
	DbHost     string `envconfig:"DB_HOST"`
	DbDriver   string `envconfig:"DB_DRIVER"`
	DbUser     string `envconfig:"DB_USER"`
	DbPassword string `envconfig:"DB_PASSWORD"`
	DbName     string `envconfig:"DB_NAME"`
	DbPort     string `envconfig:"DB_PORT"`
}

var cfg *Config

func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		Env:        "local",
		AppEnv:     "",
		Port:       "4000",
		Version:    "0.0.1",
		DbHost:     "host.docker.internal",
		DbDriver:   "postgres",
		DbUser:     "tp_user",
		DbPassword: "tp_pword",
		DbName:     "tp_db",
		DbPort:     "5432",
	}

	err := envconfig.Process("", cfg)

	if err != nil {
		return cfg, fmt.Errorf("unable to load config: %s", err.Error())
	}

	return cfg, nil
}
