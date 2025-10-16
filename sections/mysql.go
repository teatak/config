package sections

import "github.com/teatak/config/v2"

type mysql struct {
	DSN string `yaml:"dsn,omitempty"`
}

type mysqlSection map[string]*mysql

func (s *mysqlSection) SectionName() string {
	return "mysql"
}

func (s *mysqlSection) Default() *mysql {
	return Mysql["default"]
}

var Mysql = mysqlSection{}

func init() {
	config.Load(&Mysql)
}
