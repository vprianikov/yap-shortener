package handlers

import "github.com/vprianikov/yap-shortener/internal/models"

type Env struct {
	Config  models.Config
	Storage models.Storage
}
