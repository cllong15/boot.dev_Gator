package main

import (
	"context"
	"fmt"
	"log"
)

func handlerAgg(s *state, cmd command) error {
	// fetchFeed and print to console
	// if len(cmd.Args) != 1 {
	// 	log.Fatalf("usage: %v <url>", cmd.Name)
	// }
	// url := cmd.Args[0]
	url := "https://www.wagslane.dev/index.xml"
	feed, ok := fetchFeed(context.Background(), url)
	if ok != nil {
		log.Fatalf("handlerAgg feed: %v", ok)
	}
	fmt.Println(feed)
	return nil
}
