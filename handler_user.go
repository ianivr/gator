package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("login command requires a username argument")
	}

	username := cmd.args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}

	fmt.Printf("User %s has been set\n", username)
	return nil
}
