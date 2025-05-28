package db

import (
	"fmt"
	"net/url"
	"time"
)

type Config struct {
	PostgreSQLHost             string        `envconfig:"DB_POSTGRESQL_HOST" default:"localhost"`
	PostgreSQLPort             uint16        `envconfig:"DB_POSTGRESQL_PORT" default:"5432"`
	PostgreSQLUsername         string        `envconfig:"DB_POSTGRESQL_USERNAME" default:"postgres"`
	PostgreSQLPassword         string        `envconfig:"DB_POSTGRESQL_PASSWORD" required:"false"`
	PostgreSQLDatabase         string        `envconfig:"DB_POSTGRESQL_DATABASE" required:"true"`
	PostgreSQLConnectionString string        `envconfig:"DB_POSTGRESQL_CONNECTION_STRING" ignored:"true"`
	PostgreSQLSSL              bool          `envconfig:"DB_POSTGRESQL_SSL" default:"true"`
	RetrySleepTime             time.Duration `envconfig:"DB_RETRY_SLEEP" default:"1s"`
	RetryNumTimes              uint16        `envconfig:"DB_RETRY_TIMES" default:"5"`
}

func (c *Config) ComputeDependencies() error {
	connectionString, err := toURL(
		c.PostgreSQLDatabase,
		c.PostgreSQLHost,
		c.PostgreSQLUsername,
		c.PostgreSQLPassword,
		c.PostgreSQLPort,
		c.PostgreSQLSSL,
	)
	if err != nil {
		return err
	}

	c.PostgreSQLConnectionString = connectionString

	return nil
}

func toURL(db, host, username, password string, port uint16, ssl bool) (string, error) {
	if username == "" {
		username = "postgres"
	}
	if db == "" {
		db = "postgres"
	}
	if host == "" {
		host = "localhost"
	}
	if port == 0 {
		port = 5432
	}

	// URL-encode credentials
	userInfo := url.PathEscape(username)
	if password != "" {
		userInfo = fmt.Sprintf("%s:%s", userInfo, url.PathEscape(password))
	}

	mode := ""
	if !ssl {
		mode = "?sslmode=disable"
	}

	dsn := fmt.Sprintf("postgres://%s@%s:%d/%s%s", userInfo, host, port, url.PathEscape(db), mode)

	if _, err := url.Parse(dsn); err != nil {
		return "", fmt.Errorf("invalid PostgreSQL connection string: %w", err)
	}

	return dsn, nil
}
