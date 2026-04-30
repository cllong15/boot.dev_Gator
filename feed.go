package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"log"
	"net/http"
	"time"
)

type RSSFeed struct {
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	Description   string    `xml:"description"`
	Generator     string    `xml:"generator"`
	Language      string    `xml:"language"`
	LastBuildDate time.Time `xml:"lastBuildDate"`
	Atom          string    `xml:"atom,attr"`
	Channel       struct {
		ItemAttr string `xml:"item,attr"`
		Item     []struct {
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Pubdate     string `xml:"pubdate"`
			Guid        string `xml:"guid"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, ok := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if ok != nil {
		log.Fatalf("fetchFeed req: %v", ok)
	}
	req.Header.Set("User-Agent", "gator")
	client := &http.Client{}

	res, ok := client.Do(req)
	if ok != nil {
		log.Fatalf("fetchFeed res: %v", ok)
	}
	defer res.Body.Close()

	data, ok := io.ReadAll(res.Body)
	if ok != nil {
		log.Fatalf("fetchFeed data: %v", ok)
	}

	// fmt.Println(string(data))
	feed := RSSFeed{}
	ok = xml.Unmarshal(data, &feed)
	if ok != nil {
		log.Fatalf("fetchFeed, Unmarshal: %v", ok)
	}

	// fmt.Println(feed)

	feed.Title = html.UnescapeString(feed.Title)
	feed.Description = html.UnescapeString(feed.Description)
	for _, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return &feed, nil
}
