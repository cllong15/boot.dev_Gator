package main

import (
	"fmt"
	"gator/internal/config"
)

type commands struct {
	cmdName map[string]func(*state, command) error
}
type command struct {
	name string
	args []string
}

type state struct {
	config *config.Config
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("only one username")
	}
	s.config.SetUser(cmd.args[0])
	fmt.Printf("user set to %s\n", s.config.CurrentUserName)
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	ok := c.cmdName[cmd.name](s, cmd)
	if ok != nil {
		return fmt.Errorf("commands run: %v", ok)
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdName[name] = f
}
