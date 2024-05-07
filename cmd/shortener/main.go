package main

import (
	"fmt"

	"github.com/vprianikov/yap-shortener/internal/config"
	"github.com/vprianikov/yap-shortener/internal/storage"
)

func main() {
	c, errC := config.New()
	if errC != nil {
		panic(errC)
	}

	fmt.Printf("Server must be started on http://%s:%s\n", c.Host(), c.Port())

	s, errS := storage.New()
	if errS != nil {
		panic(errS)
	}

	key, errK := s.Set(`https://ya.ru`)
	if errK != nil {
		fmt.Println(errK)
	} else {
		fmt.Println(key)
	}

	url, errU := s.Get(key)
	if errU != nil {
		fmt.Println(errU)
	} else {
		fmt.Println(url)
	}
}
