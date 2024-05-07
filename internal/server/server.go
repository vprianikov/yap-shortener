package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/vprianikov/yap-shortener/internal/handlers"
)

func New(env *handlers.Env) (*http.Server, error) {
	router := gin.New()
	router.Use(gin.Recovery())
	router.HandleMethodNotAllowed = true

	router.POST(`/`, env.Shorten)
	router.GET(`/:shortKey`, env.Expand)

	return &http.Server{
		Addr:              env.Config.ServerAddress(),
		Handler:           router,
		ReadHeaderTimeout: 1 * time.Second,
	}, nil
}
