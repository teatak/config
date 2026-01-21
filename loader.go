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


var (
	loader   = &config{}
	once     sync.Once
	registry []Section
	mu       sync.RWMutex
)

// UpdateConfig 更新配置数据
// mode: "merge" (默认) - 递归合并新配置到现有配置，数组会覆盖
// mode: "overwrite" - 丢弃除 nacos 以外的所有现有配置，完全使用新配置
func UpdateConfig(data []byte, mode string) error {
	mu.Lock()
	defer mu.Unlock()

	if mode == "overwrite" {
		// 1. 备份 Nacos 配置 (防止断连)
		var nacosBackup interface{}
		if currentNacos := loader.get("nacos"); currentNacos != nil {
			// 深拷贝或直接引用均可，因为我们要创建一个全新的 loader map
			// 但为了安全，最好重新 marshal/unmarshal 或者假设 map 不会被修改
			// 这里简单起见直接引用，因为下面我们创建了 newLoader
			nacosBackup = currentNacos
		}

		// 2. 创建新 Loader (清空操作)
		newLoader := &config{}

		// 3. 恢复 Nacos 配置 (作为基底)
		if nacosBackup != nil {
			(*newLoader)["nacos"] = nacosBackup
		}

		// 4. 解析新配置 (新配置中的 nacos 会覆盖备份的，这是预期的)
		if err := yaml.Unmarshal(data, newLoader); err != nil {
			return err
		}

		// 5. 替换全局 loader
		loader = newLoader
	} else {
		// Default: Merge 模式
		if err := yaml.Unmarshal(data, loader); err != nil {
			return err
		}
	}

	// 刷新所有已注册的 section
	for _, section := range registry {
		reloadSection(section)
	}
	return nil
}

// Load 加载配置到指定的 section 结构体中
func Load(section Section) {
	mu.Lock()
	registry = append(registry, section)
	mu.Unlock()

	once.Do(LoadConfig)
	
	mu.RLock()
	defer mu.RUnlock()
	reloadSection(section)
}

func reloadSection(section Section) {
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

func (c *config) get(key string) interface{} {
	if c == nil || *c == nil {
		return nil
	}
	return (*c)[key]
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
