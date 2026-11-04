package main

import (
	"log"
	"os"

	"github.com/slyppa/blog_agregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := &state{
		cfg: &cfg,
	}

	cmds := commands{
		cmd: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("no command passed")
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}

	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
