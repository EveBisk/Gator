package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if cmd.args == nil {
		return errors.New("command expects at least one argument")
	}

	err := s.cfg.SetUser(cmd.args[0])

	if err != nil {
		return err
	}

	fmt.Printf("User has been set")

	return nil
}
