package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

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
