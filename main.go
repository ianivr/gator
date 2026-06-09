package main

import (
	"fmt"
	"os"

	"github.com/ianivr/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	newState := &state{cfg: &cfg}

	cmds := &commands{handlers: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	cmd := command{name: args[1], args: args[2:]}
	err = cmds.run(newState, cmd)
	if err != nil {
		fmt.Println("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
