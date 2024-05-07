package config

import (
	"github.com/vprianikov/yap-shortener/internal/models"
)

type config struct {
	host string
	port string
}

func New() (models.Config, error) {
	var m models.Config = &config{
		host: `localhost`,
		port: `8080`,
	}
	return m, nil
}

func (c *config) Host() string {
	return c.host
}

func (c *config) Port() string {
	return c.port
}
