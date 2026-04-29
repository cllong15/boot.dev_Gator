package main

import (
	"context"
	"database/sql"
	"errors"
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

	params := database.CreateUserParams{
		ID:        uuid.New(),
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.dbq.GetUser(context.Background(), name)
	if errors.Is(err, sql.ErrNoRows) {
		log.Fatal("No user exists")
	} else if err != nil {
		log.Fatalf("error finding user: %v", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}
