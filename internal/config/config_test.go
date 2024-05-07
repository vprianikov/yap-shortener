package config_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/vprianikov/yap-shortener/internal/config"
)

type Suite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestConfig() {
	s.Run(`check default serverAddress`, func() {
		c, err := config.New(`progName`, []string{})

		s.Require().NoError(err, `failed to initialize the configuration`)
		s.Equal(`0.0.0.0:8080`, c.ServerAddress(), `invalid value of ServerAddress`)
	})

	s.Run(`check default baseURL`, func() {
		c, err := config.New(`progName`, []string{})

		s.Require().NoError(err, `failed to initialize the configuration`)
		s.Equal(`http://localhost:8080`, c.BaseURL(), `invalid value of BaseURL`)
	})

	s.Run(`check -a serverAddress`, func() {
		c, err := config.New(`progName`, []string{`-a`, `localhost:8888`})

		s.Require().NoError(err, `failed to initialize the configuration`)
		s.Equal(`localhost:8888`, c.ServerAddress(), `invalid value of ServerAddress`)
		s.Equal(`http://localhost:8080`, c.BaseURL(), `BaseURL must be default`)
	})

	s.Run(`check -b serverAddress`, func() {
		c, err := config.New(`progName`, []string{`-b`, `localhost:8000`})

		s.Require().NoError(err, `failed to initialize the configuration`)
		s.Equal(`0.0.0.0:8080`, c.ServerAddress(), `ServerAddress must be default`)
		s.Equal(`localhost:8000`, c.BaseURL(), `invalid value of BaseURL`)
	})
}
