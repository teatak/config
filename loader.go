package config

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

type Section interface {
	SectionName() string
}

type config map[string]interface{}

func (c *config) get(key string) interface{} {
	return (*c)[key]
}

var (
	loader = &config{}
	once   sync.Once
)

// Load 加载配置到指定的 section 结构体中
func Load(section Section) {
	once.Do(LoadConfig)
	s := loader.get(section.SectionName())
	if s == nil {
		return
	}
	data, err := yaml.Marshal(s)
	if err != nil {
		log.Printf("marshal section %s error: %v\n", section.SectionName(), err)
		return
	}
	if err := yaml.Unmarshal(data, section); err != nil {
		log.Printf("unmarshal section %s error: %v\n", section.SectionName(), err)
	}
}

func appendByte(buff *bytes.Buffer, b []byte) {
	buff.Write(b)
	if len(b) > 0 && b[len(b)-1] != '\n' {
		buff.WriteByte('\n')
	}
}

// LoadConfig 加载配置文件
// 支持通过环境变量 CONFIG_PATH 指定配置文件路径，默认为 ./config/app.yml
// 支持通过环境变量 config 指定额外加载的配置文件（逗号分隔）
func LoadConfig() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		if _, err := os.Stat("./config/app.yml"); err == nil {
			configPath = "./config/app.yml"
		} else if _, err := os.Stat("./config/app.yaml"); err == nil {
			configPath = "./config/app.yaml"
		} else {
			configPath = "./config/app.yml"
		}
	}
	configDir := filepath.Dir(configPath)

	buff := bytes.Buffer{}
	env := os.Getenv("config")

	if env == "" {
		app, err := os.ReadFile(configPath)
		if err != nil {
			log.Printf("app file error: %v\n", err)
		} else {
			appendByte(&buff, app)
			if err := yaml.Unmarshal(app, &loader); err != nil {
				log.Printf("unmarshal app.yml error: %v\n", err)
			} else if c := loader.get("config"); c != nil {
				if configVal, ok := c.(string); ok {
					env = configVal
				}
			}
		}
	}

	if env != "" {
		for _, file := range strings.Split(env, ",") {
			file = strings.TrimSpace(file)
			if file == "" {
				continue
			}
			filePath := filepath.Join(configDir, file+".yml")
			if _, err := os.Stat(filePath); err != nil {
				filePathYaml := filepath.Join(configDir, file+".yaml")
				if _, err := os.Stat(filePathYaml); err == nil {
					filePath = filePathYaml
				}
			}
			b, err := os.ReadFile(filePath)
			if err != nil {
				log.Printf("file %s error: %v\n", filePath, err)
			} else {
				appendByte(&buff, b)
			}
		}
	}

	if buff.Len() > 0 {
		if err := yaml.Unmarshal(buff.Bytes(), &loader); err != nil {
			log.Printf("unmarshal config error: %v\n", err)
		}
	}
}
