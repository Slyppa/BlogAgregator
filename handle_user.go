package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no arguments found")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %v", err)
	}
	fmt.Printf("set %v as user", cmd.args[0])
	return nil
}
