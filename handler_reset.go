package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("too many args")
	}

	err := s.db.DeleteAll(context.Background())
	if err != nil {
		return fmt.Errorf("handlerDelete DeleteAll: %w", err)
	}
	return nil
}
