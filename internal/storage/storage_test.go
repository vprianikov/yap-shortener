package storage_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/vprianikov/yap-shortener/internal/models"
	"github.com/vprianikov/yap-shortener/internal/storage"
)

type Suite struct {
	suite.Suite
	storage models.Storage
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	var err error
	s.storage, err = storage.New()

	s.Require().NoError(err, `failed to initialize the storage`)
}

func (s *Suite) TestStorage() {
	s.Run(`invalid key`, func() {
		url, err := s.storage.Get(`bad`)

		//nolint:testifylint // avoid require-error: for error assertions use require
		s.Error(err, `missed error`)
		//nolint:testifylint // avoid require-error: for error assertions use require
		s.EqualError(err, `key has an invalid format`, `bad error message`)
		s.Equal(models.ExternalURL(``), url, `bad url value`)
	})

	s.Run(`not found`, func() {
		url, err := s.storage.Get(`00000000`)

		//nolint:testifylint // avoid require-error: for error assertions use require
		s.Error(err, `missed error`)
		//nolint:testifylint // avoid require-error: for error assertions use require
		s.EqualError(err, `not found`, `bad error message`)
		s.Equal(models.ExternalURL(``), url, `bad url value`)
	})

	s.Run(`success`, func() {
		var urlS models.ExternalURL = `https://practicum.yandex.ru/`

		key, errS := s.storage.Set(urlS)
		urlR, errG := s.storage.Get(key)

		s.NoError(errS, `failed to set`)
		s.NoError(errG, `failed to get`)
		s.Equal(urlS, urlR, `fetched wrong value`)
	})
}
