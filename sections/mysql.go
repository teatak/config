package sections

import "github.com/teatak/config/v2"

type mysql struct {
	Uri      string `yaml:"uri,omitempty"`
	Database string `yaml:"database,omitempty"`
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
