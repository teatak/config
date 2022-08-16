package config

import (
	"bytes"
	"io/ioutil"
	"log"
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

var loader = &config{}
var initialized = false

func Load(section Section) {
	if !initialized {
		LoadConfig()
	}
	s := loader.get(section.SectionName())
	data, _ := yaml.Marshal(s)
	_ = yaml.Unmarshal(data, section)
}

func appendByte(buff *bytes.Buffer, b []byte) {
	buff.Write(b)
	if len(b) > 0 && b[len(b)-1] != '\n' {
		buff.WriteByte('\n')
	}
}

func LoadConfig() {
	_path := "./config/app.yml"
	_dir := path.Dir(_path)
	//load config files
	buff := bytes.Buffer{}
	env := os.Getenv("config")
	if env == "" {
		app, e := ioutil.ReadFile(_path)
		if e != nil {
			log.Printf("app file error: %v\n", e)
			// os.Exit(1)
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
				log.Printf("file error: %v\n", e)
			} else {
				appendByte(&buff, b)
			}
		}
	}
	initialized = true
	_ = yaml.Unmarshal(buff.Bytes(), &loader)
}
