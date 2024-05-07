package config

import (
	"flag"

	"github.com/vprianikov/yap-shortener/internal/models"
)

type config struct {
	serverAddress string
	baseURL       string
}

func New(progName string, args []string) (models.Config, error) {
	var c config
	flags := flag.NewFlagSet(progName, flag.ContinueOnError)
	flags.StringVar(&c.serverAddress, `a`, `0.0.0.0:8080`, ``)
	flags.StringVar(&c.baseURL, `b`, `http://localhost:8080`, ``)

	if err := flags.Parse(args); err != nil {
		return nil, err
	}

	var m models.Config = &c
	return m, nil
}

func (c *config) ServerAddress() string {
	return c.serverAddress
}

func (c *config) BaseURL() string {
	return c.baseURL
}
