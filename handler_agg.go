package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %v <time between requests>", cmd.Name)
	}
	time_between_reqs, ok := time.ParseDuration(cmd.Args[0])
	if ok != nil {
		return fmt.Errorf("handleragg: could not parse duration: %v", ok)
	}
	ticker := time.NewTicker(time_between_reqs)
	fmt.Printf("Collecting feeds every %v", time_between_reqs.String())
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	nextfeed, ok := s.db.GetNextFeedToFetch(context.Background())
	if ok != nil {
		return fmt.Errorf("scrapefeeds: couldn't get next feed: %v", ok)
	}
	ok = s.db.UpdateFeed(context.Background(), database.UpdateFeedParams{
		UpdatedAt: time.Now(),
		ID:        nextfeed.ID,
	})
	if ok != nil {
		return fmt.Errorf("scrapefeeds: could not update feed: %v", ok)
	}
	feed, ok := fetchFeed(context.Background(), nextfeed.Url)
	if ok != nil {
		return fmt.Errorf("scrapefeeds: could not fetchfeed: %v", ok)
	}
	fmt.Printf("feed: %s\n", feed.Channel.Title)

	return nil
}
