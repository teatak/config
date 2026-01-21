package sections

import "github.com/teatak/config/v2"

type mysql struct {
	DSN string `yaml:"dsn,omitempty"`
}

type mysqlSection map[string]*mysql

func (s mysqlSection) Default() *mysql {
	return s["default"]
}

var Mysql = mysqlSection(config.RegisterMap[string, *mysql]("mysql"))
