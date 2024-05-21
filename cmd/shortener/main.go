package main

import (
	"os"

	"github.com/vprianikov/yap-shortener/internal/config"
	"github.com/vprianikov/yap-shortener/internal/handlers"
	"github.com/vprianikov/yap-shortener/internal/server"
	"github.com/vprianikov/yap-shortener/internal/storage"
)

func main() {
	// TODO(SSH): errC -> err
	c, errC := config.New(os.Args[0], os.Args[1:])
	if errC != nil {
		// TODO(SSH): https://app.pachca.com/chats/9180496?message=234290731
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
