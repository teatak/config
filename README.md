# Config V2

A flexible, thread-safe configuration loader for Go applications, supporting local YAML files, environment variables, and seamless integration with remote configuration centers like Nacos.

## Features

- **Generic Registration**: Use `Register[T]` and `RegisterMap[K,V]` for type-safe, boilerplate-free config loading
- **Dynamic Reloading**: Thread-safe configuration updates at runtime
- **Remote Config Support**: Easy integration patterns for Nacos, Etcd, etc.
- **Merge & Overwrite Modes**: Choose how remote config interacts with local defaults
- **Chain Loading**: Support `config: common,dev` to load multiple config files in order

## Installation

```bash
go get github.com/teatak/config/v2
```

## Quick Start

### 1. Define a Section (Simplified API ✨)

```go
package sections

import "github.com/teatak/config/v2"

type server struct {
    Port int    `yaml:"port,omitempty"`
    Name string `yaml:"name,omitempty"`
}

// One line registration - struct name "server" maps to YAML section "server"
var Server = config.Register(&server{})
```

### 2. Define Multi-Instance Section (e.g., Database Pools)

```go
package sections

import "github.com/teatak/config/v2"

type mongo struct {
    URI      string `yaml:"uri,omitempty"`
    Database string `yaml:"database,omitempty"`
}

// RegisterMap for map[string]*T configurations
var Mongo = config.RegisterMap[string, *mongo]("mongo")

// Helper function to get default instance
func MongoDefault() *mongo {
    return Mongo["default"]
}
```

### 3. Create Config Files

`config/app.yaml`:
```yaml
# Entry point - specify which config files to load (merged in order)
config: common,dev
```

`config/common.yaml`:
```yaml
auth:
  jwt_secret: "your-secret-key"
```

`config/dev.yaml`:
```yaml
server:
  port: 8080
  name: "MyApp"

mongo:
  default:
    uri: "mongodb://localhost:27017"
    database: "mydb"
  analytics:
    uri: "mongodb://localhost:27018"
    database: "analytics"
```

### 4. Use It

```go
package main

import (
    "fmt"
    "github.com/teatak/config/v2/sections"
)

func main() {
    // Config is already loaded via package init
    fmt.Printf("Server: %s on port %d\n", sections.Server.Name, sections.Server.Port)
    
    // Access multi-instance configs
    fmt.Printf("MongoDB: %s\n", sections.Mongo["default"].URI)
    fmt.Printf("Analytics DB: %s\n", sections.Mongo["analytics"].URI)
}
```

## API Reference

### Register[T any](ptr *T) *T

Registers a struct pointer as a config section. The section name is automatically derived from the struct type name (lowercase first letter).

```go
type server struct { ... }
var Server = config.Register(&server{})  // maps to YAML "server:" section
```

### RegisterMap[K comparable, V any](name string) map[K]V

Registers a map type config section. Useful for multi-instance configurations.

```go
var Mongo = config.RegisterMap[string, *mongo]("mongo")
// maps to YAML:
// mongo:
//   default: { ... }
//   secondary: { ... }
```

### Load(section Section)

Traditional API for custom section names (implements `SectionName() string`).

```go
type myLog struct { ... }
func (l *myLog) SectionName() string { return "log" }  // Custom name

var Log = &myLog{}
func init() { config.Load(Log) }
```

### UpdateConfig(data []byte, mode string) error

Updates configuration at runtime. Used for integration with remote config centers.

- `mode: "merge"` - Recursively merge new config into existing (default)
- `mode: "overwrite"` - Replace all config except `nacos` section

```go
// Example: Nacos integration
content, _ := nacosClient.GetConfig(...)
config.UpdateConfig([]byte(content), sections.Nacos.Mode)
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `CONFIG_PATH` | Explicit path to main config file (default: `./config/app.yml`) |
| `config` | Comma-separated list of config files to load (e.g., `common,dev`) |

## Config Loading Order

1. Read `CONFIG_PATH` (or default `./config/app.yml`)
2. Parse `config:` field to get file list
3. Load and merge each file in order
4. Later files override earlier ones

## Thread Safety

All config operations are protected by `sync.RWMutex`:
- Registration: Uses write lock
- Reading config values: Uses read lock  
- UpdateConfig: Uses write lock, then refreshes all registered sections

## License

MIT
