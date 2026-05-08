package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"gator/internal/database"

	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}
	// var ok error
	for _, item := range feedData.Channel.Item {
		// fmt.Println(i)
		ok := savePost(db, item, feed)
		if ok != nil {
			fmt.Println(ok)
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}

func savePost(db *database.Queries, rssItem RSSItem, feed database.Feed) error {
	pubtime, ok := time.Parse(time.RFC1123Z, rssItem.PubDate)
	if ok != nil {
		return fmt.Errorf("couldn't parse pubdate: %v", ok)
	}
	post, ok := db.CreatePost(context.Background(), database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       rssItem.Title,
		Url:         rssItem.Link,
		Description: rssItem.Description,
		PublishedAt: pubtime,
		FeedID:      feed.ID,
	})
	if ok != nil {
		if ok != sql.ErrNoRows {
			return fmt.Errorf("%w", ok)
		}
		// fmt.Printf("%v", ok)
	}
	if post.Title != "" {
		fmt.Printf("saved '%v' to database\n", post.Title)
	}
	return nil
}
