# Config

[![Go Reference](https://pkg.go.dev/badge/github.com/teatak/config/v2.svg)](https://pkg.go.dev/github.com/teatak/config/v2)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lightweight, flexible YAML configuration loader for Go applications with support for environment-based configuration files and section-based loading.

## Features

- 🔧 **Environment-based configuration** - Load different configs for dev, prod, etc.
- 📦 **Section-based loading** - Define strongly-typed config sections
- 🔀 **Config merging** - Automatically merge multiple YAML files
- 🔒 **Thread-safe** - Safe for concurrent access using `sync.Once`
- 📁 **Flexible paths** - Support custom config paths via environment variables
- 🏗️ **Built-in sections** - Pre-defined sections for common services (Redis, MongoDB, MySQL, etc.)

## Installation

```bash
go get github.com/teatak/config/v2
```

## Quick Start

### 1. Create your config files

```
config/
├── app.yml      # Main config file
├── common.yml   # Shared configuration
├── dev.yml      # Development environment
└── prod.yml     # Production environment
```

**app.yml** - Specifies which config files to load:
```yaml
config: common,dev
```

**common.yml** - Shared configuration:
```yaml
server:
  name: my-app
  port: 8080

redis:
  default:
    addr: localhost:6379
```

**dev.yml** - Development overrides:
```yaml
server:
  environment: development
```

**prod.yml** - Production overrides:
```yaml
server:
  environment: production
```

### 2. Define a config section

Create a struct that implements the `Section` interface:

```go
package sections

import "github.com/teatak/config/v2"

type server struct {
    Environment string   `yaml:"environment,omitempty"`
    Name        string   `yaml:"name,omitempty"`
    Port        int      `yaml:"port,omitempty"`
}

func (s *server) SectionName() string {
    return "server"
}

var Server = &server{}

func init() {
    config.Load(Server)
}
```

### 3. Use the configuration

```go
package main

import (
    "fmt"
    "yourapp/sections"
)

func main() {
    fmt.Printf("Server: %s running on port %d\n", 
        sections.Server.Name, 
        sections.Server.Port)
}
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `CONFIG_PATH` | Path to the main config file | `./config/app.yml` |
| `config` | Comma-separated list of config files to load (overrides `config` key in app.yml) | - |

### Examples

```bash
# Use default config path with production configs
export config=common,prod && ./myapp

# Use custom config path
export CONFIG_PATH=/etc/myapp/app.yml && ./myapp

# Combine both
export CONFIG_PATH=/etc/myapp/app.yml config=common,prod && ./myapp
```

## Built-in Sections

The package includes pre-defined sections for common services:

| Section | Description |
|---------|-------------|
| `Server` | Server configuration (name, port, environment) |
| `Redis` | Redis connection settings |
| `Mongo` | MongoDB connection settings |
| `MySQL` | MySQL database configuration |
| `Log` | Logging configuration |
| `SMTP` | Email server settings |
| `Alipay` | Alipay payment integration |
| `WechatPay` | WeChat Pay integration |
| `Wechat` | WeChat integration |
| `Aliyun` | Aliyun cloud services |
| `GitHub` | GitHub OAuth settings |
| `GitLab` | GitLab OAuth settings |
| `Riff` | Riff service discovery |

### Using built-in sections

```go
import "github.com/teatak/config/v2/sections"

func main() {
    // Access server config
    fmt.Println(sections.Server.Port)
    
    // Access Redis config (supports multiple instances)
    fmt.Println(sections.Redis.Default().Addr)
    
    // Access MongoDB config
    fmt.Println(sections.Mongo.Default().URI)
}
```

## Advanced Usage

### Map-based sections (multiple instances)

For services that may have multiple instances (like databases), use a map type:

```go
type mongo struct {
    URI            string `yaml:"uri,omitempty"`
    Database       string `yaml:"database,omitempty"`
    ConnectTimeout uint64 `yaml:"connect_timeout,omitempty"`
    MaxPoolSize    uint64 `yaml:"max_pool_size,omitempty"`
}

type mongoSection map[string]*mongo

func (s *mongoSection) SectionName() string {
    return "mongo"
}

func (s *mongoSection) Default() *mongo {
    return (*s)["default"]
}

var Mongo = mongoSection{}

func init() {
    config.Load(&Mongo)
}
```

**YAML configuration:**
```yaml
mongo:
  default:
    uri: mongodb://localhost:27017
    database: myapp
    max_pool_size: 100
  analytics:
    uri: mongodb://analytics-host:27017
    database: analytics
```

**Usage:**
```go
// Access default instance
db := sections.Mongo.Default()

// Access named instance
analyticsDB := sections.Mongo["analytics"]
```

## How It Works

1. On first call to `config.Load()`, the loader reads the main config file (`app.yml` by default)
2. It looks for a `config` key specifying additional files to load
3. Additional files are loaded and merged in order (later files override earlier ones)
4. The merged configuration is cached for subsequent `Load()` calls
5. Each section struct receives only its relevant portion of the configuration

## License

MIT License - see [LICENSE](LICENSE) for details.