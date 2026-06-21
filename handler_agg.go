package main

import (
	"context"
	"fmt"
	"time"
)

func handleAggregate(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("agg command requires a time between requests")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	feed, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	feedFetched, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range feedFetched.Channel.Item {
		fmt.Printf("- %v\n", item.Title)
	}

	return nil
}
