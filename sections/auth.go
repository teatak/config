package sections

import "github.com/teatak/config/v2"

type auth struct {
	JWTSecret string            `yaml:"jwt_secret" json:"jwt_secret"`
	OAuth2    map[string]OAuth2 `yaml:"oauth2" json:"oauth2"`
}

func (c *auth) SectionName() string {
	return "auth"
}

type OAuth2 struct {
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	RedirectURL  string   `yaml:"redirect_url"`
	Scopes       []string `yaml:"scopes"`
}

var Auth = auth{}

func init() {
	config.Load(&Auth)
}
