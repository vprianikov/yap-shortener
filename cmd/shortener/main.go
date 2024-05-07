package main

import (
	"fmt"

	"github.com/vprianikov/yap-shortener/internal/config"
)

func main() {
	c, errC := config.New()
	if errC != nil {
		panic(errC)
	}

	fmt.Printf("Server must be started on http://%s:%s\n", c.Host(), c.Port())
}
