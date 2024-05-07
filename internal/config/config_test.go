package config_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/vprianikov/yap-shortener/internal/config"
	"github.com/vprianikov/yap-shortener/internal/models"
)

type Suite struct {
	suite.Suite
	config models.Config
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	var err error
	s.config, err = config.New()

	s.Require().NoError(err, `failed to initialize the configuration`)
}

func (s *Suite) TestConfig() {
	s.Run(`check host`, func() {
		s.Equal(`localhost`, s.config.Host(), `invalid value of Host`)
	})

	s.Run(`check port`, func() {

		s.Equal(`8080`, s.config.Port(), `invalid value of Port`)
	})
}
