# config

## yml
|-config<br/>
&nbsp;|-app.yml<br/>
&nbsp;|-common.yml<br/>
&nbsp;|-dev.yml<br/>
&nbsp;|-prod.yml<br/>

app.yml
```yaml
config: common,dev
```

common.yml
```yaml
#common config
riff:
  url: riff://localhost:8610
```

dev.yml
```yaml
#dev config
server:
  name: server-dev
  port: 8080
```

prod.yml
```yaml
#prod config
server:
  name: server-prod
  port: 8080
```

## sections
|-sections<br/>
&nbsp;&nbsp;|-server.go<br/>
```go
package sections

import "github.com/teatak/config"

type server struct {
	Name      string `yaml:"name,omitempty"`
	Port      int    `yaml:"port,omitempty"`
}

func (s *server) SectionName() string {
	return "server"
}

var Server = &server{}

func init() {
	config.Load(Server)
}

```


## run as env
```bash
export config=common,prod && ./test
```