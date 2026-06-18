package main

import (
	"context"
	"fmt"
)

func handleAggregate(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	fmt.Printf("Feed: %v\n", feed)
	return nil
}
