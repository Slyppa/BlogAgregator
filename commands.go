package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	cmd map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmd[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	value, ok := c.cmd[cmd.name]
	if !ok {
		fmt.Errorf("fail to run command")
	}
	// runs command
	return value(s, cmd)
}
