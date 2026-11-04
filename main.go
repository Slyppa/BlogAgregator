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

	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("no command passed")
	}

	err = cmds.run(programState, command{Name: args[1], Args: args[2:]})
	if err != nil {
		log.Fatal(err)
	}
}
