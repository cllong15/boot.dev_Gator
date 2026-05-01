package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	cur_user, ok := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if ok != nil {
		log.Fatalf("could not get user: %v", ok)
	}
	feed, ok := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    cur_user.ID,
	})
	if ok != nil {
		log.Fatalf("could not create feed: %v", ok)
	}
	fmt.Println(feed)
	return nil
}
