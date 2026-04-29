package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user_uuid := uuid.New()

	params := database.CreateUserParams{
		ID:        user_uuid.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	_, ok := s.dbq.CreateUser(context.Background(), params)
	if ok != nil {
		log.Fatalf("%s already exists", params.Name)
	}
	s.cfg.CurrentUserName = params.Name
	fmt.Printf("user %s was created\n", params.Name)

	ok = s.cfg.SetUser(name)
	if ok != nil {
		return fmt.Errorf("couldn't set current user: %w", ok)
	}
	return nil
}
