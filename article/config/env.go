package config

import (
	"fmt"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var Env ENV

func init() {
	env := os.Getenv("GO_ENV")

	if err := envconfig.Process(env, &Env); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type ENV struct {
	MySQL
}

type MySQL struct {
	Dsn             string        `envconfig:"MYSQL_DSN"               required:"true"`
	MaxConn         int           `envconfig:"MYSQL_MAX_CONN"          default:"25"`
	MaxIdleConn     int           `envconfig:"MYSQL_MAX_IDLE"          default:"25"`
	MaxConnLifetime time.Duration `envconfig:"MYSQL_MAX_CONN_LIFETIME" default:"300s"`
}
