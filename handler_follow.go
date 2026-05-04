package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		log.Fatalf("usage: %v <url>", cmd.Name)
	}
	url := cmd.Args[0]
	feed, ok := s.db.GetFeedByURL(context.Background(), url)
	if ok != nil {
		return fmt.Errorf("feed: %v", ok)
	}
	cur_user, ok := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if ok != nil {
		return fmt.Errorf("cur_user: %v", ok)
	}
	feed_follow, ok := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    cur_user.ID,
		FeedID:    feed.ID,
	})
	if ok != nil {
		return fmt.Errorf("feed_follow: %v", ok)
	}
	fmt.Println(feed_follow)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	// It should print all the names of the feeds the current user is following.
	cur_user := s.cfg.CurrentUserName
	feeds, ok := s.db.GetFeedFollowsForUser(context.Background(), cur_user)
	if ok != nil {
		return ok
	}
	for _, feed := range feeds {
		fmt.Println(feed)
	}
	return nil
}
