package main

import (
	"github.com/vprianikov/yap-shortener/internal/config"
	"github.com/vprianikov/yap-shortener/internal/handlers"
	"github.com/vprianikov/yap-shortener/internal/server"
	"github.com/vprianikov/yap-shortener/internal/storage"
)

func main() {
	c, errC := config.New()
	if errC != nil {
		panic(errC)
	}

	s, errS := storage.New()
	if errS != nil {
		panic(errS)
	}

	l, errL := server.New(&handlers.Env{
		Config:  c,
		Storage: s,
	})
	if errL != nil {
		panic(errL)
	}

	if err := l.ListenAndServe(); err != nil {
		panic(err)
	}
}
