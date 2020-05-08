package context

import (
	"github.com/krakowitzerm/ghadmin/internal/config"
)

// TODO these are sprinkled across command, context, config, and ghrepo
const defaultHostname = "github.com"

// Context represents the interface for querying information about the current environment
type Context interface {
	AuthToken() (string, error)
	SetAuthToken(string)
	AuthLogin() (string, error)
	Config() (config.Config, error)
}

// New initializes a Context that reads from the filesystem
func New() Context {
	return &fsContext{}
}

// A Context implementation that queries the filesystem
type fsContext struct {
	config    config.Config
	authToken string
}

func (c *fsContext) Config() (config.Config, error) {
	if c.config == nil {
		config, err := config.ParseOrSetupConfigFile(config.ConfigFile())
		if err != nil {
			return nil, err
		}
		c.config = config
		c.authToken = ""
	}
	return c.config, nil
}

func (c *fsContext) AuthToken() (string, error) {
	if c.authToken != "" {
		return c.authToken, nil
	}

	cfg, err := c.Config()
	if err != nil {
		return "", err
	}

	token, err := cfg.Get(defaultHostname, "oauth_token")
	if token == "" || err != nil {
		return "", err
	}

	return token, nil
}

func (c *fsContext) SetAuthToken(t string) {
	c.authToken = t
}

func (c *fsContext) AuthLogin() (string, error) {
	config, err := c.Config()
	if err != nil {
		return "", err
	}

	login, err := config.Get(defaultHostname, "user")
	if login == "" || err != nil {
		return "", err
	}

	return login, nil
}
