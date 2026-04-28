package main

import (
	"log"
	"os"

	"gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	// fmt.Println(cfg)
	p_state := state{
		config: &cfg,
	}
	// fmt.Println(p_state)
	p_cmds := commands{
		map[string]func(*state, command) error{},
	}
	// fmt.Println(p_cmds)
	p_cmds.register("login", handlerLogin)
	os_args := os.Args
	if len(os_args) < 3 {
		os.Exit(1)
	}
	// fmt.Print(os_args)
	cmd := command{
		name: os_args[1],
		args: os_args[2:],
	}
	p_cmds.run(&p_state, cmd)
}
