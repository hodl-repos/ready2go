package config

type Config struct {
	DeveloperToken string `toml:"developertoken"`
	AccountToken   string `toml:"accounttoken"`
}

// c stores the current config value
var c *Config

func init() {
	c = &Config{"", ""}
}

func SetDeveloperToken(t string) {
	if t != "" {
		c.DeveloperToken = t
	}
}

func DeveloperToken() string {
	return c.DeveloperToken
}

func SetAccountToken(t string) {
	if t != "" {
		c.AccountToken = t
	}
}
