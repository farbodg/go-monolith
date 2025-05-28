package config

import "go-monolith/db"

type (
	DBConfig = db.Config
)

type Config struct {
	DBConfig
	LogLevel       string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	ServerHTTPPort uint16 `envconfig:"SERVER_HTTP_PORT" default:"8000"`
}

func (c *Config) ComputeDependencies() error {
	return c.DBConfig.ComputeDependencies()
}
