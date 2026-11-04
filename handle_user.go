package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %v", err)
	}

	fmt.Printf("User switched to %v successfully!\n", cmd.Args[0])
	return nil
}
