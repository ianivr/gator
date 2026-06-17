package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ianivr/gator/internal/config"
	"github.com/ianivr/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	newState := &state{cfg: &cfg}

	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)
	newState.db = dbQueries

	cmds := &commands{handlers: make(map[string]func(*state, command) error)}
	cmds.register("login", handleLogin)
	cmds.register("register", handleRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handleUsers)
	cmds.register("agg", handleAggregate)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	cmd := command{name: args[1], args: args[2:]}
	err = cmds.run(newState, cmd)
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
