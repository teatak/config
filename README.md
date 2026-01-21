# Config V2

A flexible, thread-safe configuration loader for Go applications, supporting local YAML files, environment variables, and seamless integration with remote configuration centers like Nacos.

## Features

- **Lazy & Explicit Loading**: Supports zero-config `init` loading or explicit `Init` with custom paths.
- **Dynamic Reloading**: Thread-safe configuration updates at runtime.
- **Remote Config Support**: Easy integration patterns for Nacos, Etcd, etc.
- **Merge & Overwrite Modes**: Choose how remote config interacts with local defaults.
- **Section-based Design**: Modularize configuration into clean, separate structs.

## Installation

```bash
go get github.com/teatak/config/v2
```

## Quick Start

### 1. Define a Section

Create a struct for your configuration section and register it.

```go
package sections

import "github.com/teatak/config/v2"

type ServerConfig struct {
    Port int    `yaml:"port"`
    Name string `yaml:"name"`
}

func (s *ServerConfig) SectionName() string {
    return "server"
}

var Server = &ServerConfig{}

func init() {
    // Register the section to be auto-filled
    config.Load(Server)
}
```

### 2. Create Config File

`config/app.yml`:

```yaml
server:
  port: 8080
  name: "MyApp"
```

### 3. Use It

```go
package main

import (
    "fmt"
    "github.com/teatak/config/v2/sections"
)

func main() {
    // Config is already loaded via init()
    fmt.Printf("Server running on port %d\n", sections.Server.Port)
}
```

## Advanced Usage

### Explicit Initialization

Override default paths or reload explicitly.

```go
func main() {
    // Load specific file
    config.Init("./config/production.yml")
}
```

### Nacos Integration (Example)

Use `UpdateConfig` to integrate with Nacos or any other config source.

**1. Define Nacos Section (`sections/nacos.go`)**

```go
type nacos struct {
    Enable bool   `yaml:"enable"`
    IpAddr string `yaml:"ipAddr"`
    // ...
    Mode   string `yaml:"mode"` // "merge" or "overwrite"
}
// ... implementation ...
```

**2. Bootstrap & Watch**

```go
func main() {
    // 1. Local config loaded automatically
    if sections.Nacos.Enable {
        // 2. Connect to Nacos using sections.Nacos info
        client, _ := clients.NewConfigClient(...)
        
        // 3. Fetch config
        content, _ := client.GetConfig(...)

        // 4. Update Global Config
        // Mode: "merge" (patch) or "overwrite" (replace all except nacos)
        config.UpdateConfig([]byte(content), sections.Nacos.Mode)
    }
    
    // Application logic...
}
```

### Environment Variables

- `CONFIG_PATH`: Explicit path to the main config file.
- `config`: Comma-separated list of extra config files to load (e.g., `config=db,redis` loads `db.yml` and `redis.yml`).

## License

MIT
