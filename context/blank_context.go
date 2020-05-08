package context

import (
	"fmt"

	"github.com/krakowitzerm/ghadmin/internal/config"
)

// NewBlank initializes a blank Context suitable for testing
func NewBlank() *blankContext {
	return &blankContext{}
}

// A Context implementation that queries the filesystem
type blankContext struct {
	authToken string
	authLogin string
}

func (c *blankContext) Config() (config.Config, error) {
	cfg, err := config.ParseConfig("boom.txt")
	if err != nil {
		panic(fmt.Sprintf("failed to parse config during tests. did you remember to stub? error: %s", err))
	}
	return cfg, nil
}

func (c *blankContext) AuthToken() (string, error) {
	return c.authToken, nil
}

func (c *blankContext) SetAuthToken(t string) {
	c.authToken = t
}

func (c *blankContext) SetAuthLogin(login string) {
	c.authLogin = login
}

func (c *blankContext) AuthLogin() (string, error) {
	return c.authLogin, nil
}
