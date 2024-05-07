package storage

import (
	"crypto/rand"
	"errors"
	"math/big"
	"regexp"
	"sync"

	"github.com/vprianikov/yap-shortener/internal/models"
)

type storage struct {
	sync.RWMutex
	store map[models.ShortKey]models.ExternalURL
}

func New() (models.Storage, error) {
	var m models.Storage = &storage{
		store: make(map[models.ShortKey]models.ExternalURL),
	}
	return m, nil
}

func (s *storage) Set(url models.ExternalURL) (models.ShortKey, error) {
	key, err := getRandomKey()
	if err != nil {
		return ``, err
	}

	s.Lock()
	_, ok := s.store[key]
	if !ok {
		s.store[key] = url
	}
	s.Unlock()

	if ok {
		return ``, errors.New(`generated key is not unique`)
	}
	return key, nil
}

func (s *storage) Get(key models.ShortKey) (models.ExternalURL, error) {
	if !checkKey(key) {
		return ``, errors.New(`key has an invalid format`)
	}

	s.RLock()
	url, ok := s.store[key]
	s.RUnlock()

	if !ok {
		return ``, errors.New(`not found`)
	}
	return url, nil
}

func checkKey(key models.ShortKey) bool {
	r, e := regexp.MatchString(`\A[A-Za-z0-9]{8}\z`, string(key))
	if e == nil && r {
		return true
	}
	return false
}

func getRandomKey() (models.ShortKey, error) {
	const letters = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`
	const length = 8

	s := ``

	for range length {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ``, err
		}
		s += string(letters[randomIndex.Int64()])
	}

	if !checkKey(models.ShortKey(s)) {
		return ``, errors.New(`generated key has an invalid format`)
	}
	return models.ShortKey(s), nil
}
