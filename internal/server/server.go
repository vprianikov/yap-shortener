package server

import (
	"net/http"
	"time"

	"github.com/vprianikov/yap-shortener/internal/handlers"
)

func New(env *handlers.Env) (*http.Server, error) {
	mux := http.NewServeMux()

	return &http.Server{
		Addr:              `0.0.0.0:` + env.Config.Port(),
		Handler:           mux,
		ReadHeaderTimeout: 1 * time.Second,
	}, nil
}
