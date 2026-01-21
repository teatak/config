package sections_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/teatak/config/v2/sections"
)

func init() {
	// 设置配置文件路径
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filepath.Dir(filename))
	configPath := filepath.Join(dir, "config", "app.yaml")
	os.Setenv("CONFIG_PATH", configPath)
}

func TestServerSection(t *testing.T) {
	if sections.Server == nil {
		t.Fatal("Expected Server to be initialized")
	}
	t.Logf("Server: %+v", sections.Server)
}

func TestMongoSection(t *testing.T) {
	if sections.Mongo == nil {
		t.Fatal("Expected Mongo to be initialized")
	}
	if sections.Mongo.Default() == nil {
		t.Log("Mongo.Default() is nil (may not be configured)")
	} else {
		t.Logf("Mongo default: %+v", sections.Mongo.Default())
	}
}

func TestRedisSection(t *testing.T) {
	if sections.Redis == nil {
		t.Fatal("Expected Redis to be initialized")
	}
	if sections.Redis.Default() == nil {
		t.Log("Redis.Default() is nil (may not be configured)")
	} else {
		t.Logf("Redis default: %+v", sections.Redis.Default())
	}
}

func TestAuthSection(t *testing.T) {
	if sections.Auth == nil {
		t.Fatal("Expected Auth to be initialized")
	}
	t.Logf("Auth: JWTSecret=%s..., OAuth2 providers=%d",
		sections.Auth.JWTSecret[:min(10, len(sections.Auth.JWTSecret))],
		len(sections.Auth.OAuth2))
}

func TestAllSectionsInitialized(t *testing.T) {
	// 验证所有 section 都已初始化
	checks := []struct {
		name string
		ok   bool
	}{
		{"Server", sections.Server != nil},
		{"Nacos", sections.Nacos != nil},
		{"Log", sections.Log != nil},
		{"Auth", sections.Auth != nil},
		{"Smtp", sections.Smtp != nil},
		{"Mongo", sections.Mongo != nil},
		{"Redis", sections.Redis != nil},
		{"Mysql", sections.Mysql != nil},
		{"Alipay", sections.Alipay != nil},
		{"Aliyun", sections.Aliyun != nil},
		{"Wechat", sections.Wechat != nil},
		{"WechatPay", sections.WechatPay != nil},
		{"Github", sections.Github != nil},
		{"Gitlab", sections.Gitlab != nil},
		{"Riff", sections.Riff != nil},
	}

	for _, c := range checks {
		if !c.ok {
			t.Errorf("Section %s is not initialized", c.name)
		}
	}

	t.Log("✅ All sections initialized successfully")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
