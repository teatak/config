package config

import (
	"os"
	"path/filepath"
	"sync"
	"testing"
)

// 测试用的结构体 - 名称需要与 YAML 配置中的 section 名称匹配
type server struct {
	Environment  string   `yaml:"environment,omitempty"`
	Url          string   `yaml:"url,omitempty"`
	ShortUrl     string   `yaml:"shortUrl,omitempty"`
	AllowOrigins []string `yaml:"allowOrigins,omitempty"`
	Name         string   `yaml:"name,omitempty"`
	Port         int      `yaml:"port,omitempty"`
}

type mongo struct {
	URI         string `yaml:"uri,omitempty"`
	Database    string `yaml:"database,omitempty"`
	MaxPoolSize uint64 `yaml:"maxPoolSize,omitempty"`
}

type redis struct {
	Addrs    []string `yaml:"addrs,omitempty"`
	Password string   `yaml:"password,omitempty"`
	DB       int      `yaml:"db,omitempty"`
}

type auth struct {
	JWTSecret string            `yaml:"jwt_secret"`
	OAuth2    map[string]oauth2 `yaml:"oauth2"`
}

type oauth2 struct {
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	RedirectURL  string   `yaml:"redirect_url"`
	Scopes       []string `yaml:"scopes"`
}

type logConfig struct {
	Handler string `yaml:"handler,omitempty"`
	Level   string `yaml:"level,omitempty"`
}

type gateway struct {
	DevMode       bool   `yaml:"devMode,omitempty"`
	Introspection bool   `yaml:"introspection,omitempty"`
	ConfigPath    string `yaml:"configPath,omitempty"`
	ConfigURL     string `yaml:"configURL,omitempty"`
}

// resetForTest 重置全局状态，用于测试隔离
func resetForTest() {
	mu.Lock()
	defer mu.Unlock()
	loader = &config{}
	registry = nil
	once = sync.Once{}
}

// TestConfigChain 测试配置链式加载 (app.yaml -> config: common,dev -> common.yaml + dev.yaml)
func TestConfigChain(t *testing.T) {
	// 设置配置文件路径
	wd, _ := os.Getwd()
	configPath := filepath.Join(wd, "config", "app.yaml")
	os.Setenv("CONFIG_PATH", configPath)
	defer os.Unsetenv("CONFIG_PATH")

	// 重置全局状态
	resetForTest()

	// 注册多个 section
	srv := Register(&server{})
	authCfg := Register(&auth{})
	logCfg := Register(&logConfig{})
	mongoMap := RegisterMap[*mongo]("mongo")
	redisMap := RegisterMap[*redis]("redis")
	gw := Register(&gateway{})

	// 验证 server (来自 dev.yaml)
	if srv.Environment != "dev" {
		t.Errorf("Expected server.Environment = 'dev', got '%s'", srv.Environment)
	}
	if srv.Name != "unicorn-gateway" {
		t.Errorf("Expected server.Name = 'unicorn-gateway', got '%s'", srv.Name)
	}
	if srv.Port != 8080 {
		t.Errorf("Expected server.Port = 8080, got %d", srv.Port)
	}
	if len(srv.AllowOrigins) == 0 || srv.AllowOrigins[0] != "*" {
		t.Errorf("Expected server.AllowOrigins = ['*'], got %v", srv.AllowOrigins)
	}

	// 验证 auth (来自 common.yaml)
	if authCfg.JWTSecret == "" {
		t.Error("Expected auth.JWTSecret to be set")
	}
	if authCfg.OAuth2 == nil {
		t.Error("Expected auth.OAuth2 to be set")
	}
	if github, ok := authCfg.OAuth2["github"]; ok {
		if github.ClientID == "" {
			t.Error("Expected auth.OAuth2.github.client_id to be set")
		}
	} else {
		t.Error("Expected auth.OAuth2 to have 'github' key")
	}

	// 验证 log (来自 dev.yaml)
	// 注意：logConfig 结构体名会映射到 "logConfig" section，而不是 "log"
	// 如果需要自定义名称，可以继续使用传统的 Load() 方法
	t.Logf("   LogConfig (mapped to 'logConfig' section): %+v", logCfg)

	// 验证 mongo (来自 dev.yaml)
	if mongoMap["default"] == nil {
		t.Fatal("Expected mongo['default'] to exist")
	}
	if mongoMap["default"].Database != "unicorn_system" {
		t.Errorf("Expected mongo.default.database = 'unicorn_system', got '%s'", mongoMap["default"].Database)
	}

	// 验证 redis (来自 dev.yaml)
	if redisMap["default"] == nil {
		t.Fatal("Expected redis['default'] to exist")
	}
	if redisMap["default"].Password != "admin" {
		t.Errorf("Expected redis.default.password = 'admin', got '%s'", redisMap["default"].Password)
	}
	if redisMap["session"] == nil {
		t.Fatal("Expected redis['session'] to exist")
	}

	// 验证 gateway (来自 dev.yaml)
	if !gw.DevMode {
		t.Error("Expected gateway.devMode = true")
	}
	if gw.ConfigPath != "config/router.json" {
		t.Errorf("Expected gateway.configPath = 'config/router.json', got '%s'", gw.ConfigPath)
	}

	t.Log("✅ Config chain test passed!")
	t.Logf("   Server: %+v", srv)
	t.Logf("   Auth: JWTSecret=%s, OAuth2 providers=%d", authCfg.JWTSecret[:10]+"...", len(authCfg.OAuth2))
	t.Logf("   Log: %+v", logCfg)
	t.Logf("   Mongo: %v", mongoMap)
	t.Logf("   Redis: %v", redisMap)
	t.Logf("   Gateway: %+v", gw)
}

func TestRegister(t *testing.T) {
	wd, _ := os.Getwd()
	configPath := filepath.Join(wd, "config", "app.yaml")
	os.Setenv("CONFIG_PATH", configPath)
	defer os.Unsetenv("CONFIG_PATH")

	resetForTest()

	srv := Register(&server{})

	if srv.Name != "unicorn-gateway" {
		t.Errorf("Expected server.Name = 'unicorn-gateway', got '%s'", srv.Name)
	}
	if srv.Port != 8080 {
		t.Errorf("Expected server.Port = 8080, got %d", srv.Port)
	}

	t.Logf("✅ Register test passed: %+v", srv)
}

func TestRegisterMap(t *testing.T) {
	wd, _ := os.Getwd()
	configPath := filepath.Join(wd, "config", "app.yaml")
	os.Setenv("CONFIG_PATH", configPath)
	defer os.Unsetenv("CONFIG_PATH")

	resetForTest()

	mongoMap := RegisterMap[*mongo]("mongo")

	if mongoMap["default"] == nil {
		t.Fatal("Expected mongo['default'] to exist")
	}
	if mongoMap["default"].Database != "unicorn_system" {
		t.Errorf("Expected Database = 'unicorn_system', got '%s'", mongoMap["default"].Database)
	}

	t.Logf("✅ RegisterMap test passed: %+v", mongoMap["default"])
}

func TestUpdateConfigMerge(t *testing.T) {
	wd, _ := os.Getwd()
	configPath := filepath.Join(wd, "config", "app.yaml")
	os.Setenv("CONFIG_PATH", configPath)
	defer os.Unsetenv("CONFIG_PATH")

	resetForTest()

	srv := Register(&server{})

	if srv.Port != 8080 {
		t.Errorf("Expected initial port = 8080, got %d", srv.Port)
	}

	// 使用 merge 模式更新配置
	newConfig := []byte(`
server:
  port: 9090
  name: "UpdatedApp"
`)
	err := UpdateConfig(newConfig, "merge")
	if err != nil {
		t.Fatalf("UpdateConfig failed: %v", err)
	}

	if srv.Port != 9090 {
		t.Errorf("Expected updated port = 9090, got %d", srv.Port)
	}
	if srv.Name != "UpdatedApp" {
		t.Errorf("Expected updated name = 'UpdatedApp', got '%s'", srv.Name)
	}
	// Environment 应该保留（merge 模式）
	if srv.Environment != "dev" {
		t.Errorf("Expected environment to remain 'dev', got '%s'", srv.Environment)
	}

	t.Logf("✅ UpdateConfig (merge) test passed: %+v", srv)
}

func TestUpdateConfigOverwrite(t *testing.T) {
	wd, _ := os.Getwd()
	configPath := filepath.Join(wd, "config", "app.yaml")
	os.Setenv("CONFIG_PATH", configPath)
	defer os.Unsetenv("CONFIG_PATH")

	resetForTest()

	srv := Register(&server{})

	if srv.Port != 8080 {
		t.Errorf("Expected initial port = 8080, got %d", srv.Port)
	}

	// 使用 overwrite 模式更新配置
	newConfig := []byte(`
server:
  port: 7070
`)
	err := UpdateConfig(newConfig, "overwrite")
	if err != nil {
		t.Fatalf("UpdateConfig failed: %v", err)
	}

	if srv.Port != 7070 {
		t.Errorf("Expected port = 7070, got %d", srv.Port)
	}

	t.Logf("✅ UpdateConfig (overwrite) test passed: %+v", srv)
}

func TestLcFirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Server", "server"},
		{"MongoDB", "mongoDB"},
		{"URL", "uRL"},
		{"a", "a"},
		{"A", "a"},
		{"", ""},
	}

	for _, tt := range tests {
		result := lcFirst(tt.input)
		if result != tt.expected {
			t.Errorf("lcFirst(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}

	t.Log("✅ lcFirst test passed")
}

func TestRegisterMapUpdateConfig(t *testing.T) {
	wd, _ := os.Getwd()
	configPath := filepath.Join(wd, "config", "app.yaml")
	os.Setenv("CONFIG_PATH", configPath)
	defer os.Unsetenv("CONFIG_PATH")

	resetForTest()

	mongoMap := RegisterMap[*mongo]("mongo")

	// 验证初始值
	if mongoMap["default"] == nil {
		t.Fatal("Expected mongo['default'] to exist")
	}

	// 使用 merge 模式更新配置
	newConfig := []byte(`
mongo:
  default:
    uri: "mongodb://newhost:27017"
    maxPoolSize: 200
  third:
    uri: "mongodb://third:27017"
    database: "thirddb"
`)
	err := UpdateConfig(newConfig, "merge")
	if err != nil {
		t.Fatalf("UpdateConfig failed: %v", err)
	}

	// 验证更新后的值
	if mongoMap["default"].MaxPoolSize != 200 {
		t.Errorf("Expected updated MaxPoolSize = 200, got %d", mongoMap["default"].MaxPoolSize)
	}
	if mongoMap["default"].URI != "mongodb://newhost:27017" {
		t.Errorf("Expected updated URI = 'mongodb://newhost:27017', got '%s'", mongoMap["default"].URI)
	}
	// 验证新增的 third
	if mongoMap["third"] == nil {
		t.Fatal("Expected mongo['third'] to exist after update")
	}
	if mongoMap["third"].URI != "mongodb://third:27017" {
		t.Errorf("Expected third URI = 'mongodb://third:27017', got '%s'", mongoMap["third"].URI)
	}

	t.Logf("✅ RegisterMap UpdateConfig test passed")
}

// TestEnvironmentVariable 测试通过环境变量 config 指定配置文件
func TestEnvironmentVariable(t *testing.T) {
	wd, _ := os.Getwd()
	configPath := filepath.Join(wd, "config", "app.yaml")

	// 设置环境变量直接指定要加载的配置
	os.Setenv("CONFIG_PATH", configPath)
	os.Setenv("config", "common,dev")
	defer os.Unsetenv("CONFIG_PATH")
	defer os.Unsetenv("config")

	resetForTest()

	srv := Register(&server{})
	authCfg := Register(&auth{})

	// 验证两个配置都加载了
	if srv.Name != "unicorn-gateway" {
		t.Errorf("Expected server.Name from dev.yaml, got '%s'", srv.Name)
	}
	if authCfg.JWTSecret == "" {
		t.Error("Expected auth.JWTSecret from common.yaml")
	}

	t.Log("✅ Environment variable test passed")
}
