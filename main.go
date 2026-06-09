package main

import (
	"fmt"

	"github.com/ianivr/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	err = cfg.SetUser("Ianiv")
	if err != nil {
		panic(err)
	}

	cfg, err = config.Read()
	if err != nil {
		panic(err)
	}

	fmt.Println("Current config:", cfg)
}
