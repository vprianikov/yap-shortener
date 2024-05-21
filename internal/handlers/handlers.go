package handlers

import "github.com/vprianikov/yap-shortener/internal/models"

type Env struct {
	// TODO(SSH): надо сделать неэкспортируемым
	Config models.Config
	// TODO(SSH): надо сделать неэкспортируемым
	Storage models.Storage
}
