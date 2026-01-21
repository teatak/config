package sections

import "github.com/teatak/config/v2"

type mysql struct {
	DSN string `yaml:"dsn,omitempty"`
}

var Mysql = config.RegisterMap[*mysql]("mysql")
