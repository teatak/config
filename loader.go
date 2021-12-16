package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"
)

type Section interface {
	SectionName() string
}

type config map[string]interface{}

func (c *config) get(key string) interface{} {
	return (*c)[key]
}

func Load(section Section) {
	s := loader.get(section.SectionName())
	data, _ := yaml.Marshal(s)
	_ = yaml.Unmarshal(data, section)
}

var loader = &config{}

func appendByte(buff *bytes.Buffer, b []byte) {
	buff.Write(b)
	if len(b) > 0 && b[len(b)-1] != '\n' {
		buff.WriteByte('\n')
	}
}

func LoadConfig(configPath string) {
	_path := "./config/app.yml"
	if configPath != "" {
		_path = configPath
	}
	_dir := path.Dir(_path)
	//load config files
	buff := bytes.Buffer{}
	env := os.Getenv("config")
	if env == "" {
		app, e := ioutil.ReadFile(_path)
		if e != nil {
			fmt.Printf("File error: %v\n", e)
			os.Exit(1)
		} else {
			appendByte(&buff, app)
			_ = yaml.Unmarshal(app, &loader)
			if c := loader.get("config"); c != nil {
				env = c.(string)
			}
		}
	}
	if env != "" {
		for _, file := range strings.Split(env, ",") {
			b, e := ioutil.ReadFile(_dir + "/" + file + ".yml")
			if e != nil {
				fmt.Printf("File error: %v\n", e)
			} else {
				appendByte(&buff, b)
			}
		}
	}
	_ = yaml.Unmarshal(buff.Bytes(), &loader)
}
