package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ianivr/gator/internal/database"
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
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: parsePublishedDate(item.PubDate),
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "23505") {
				continue
			} else {
				log.Printf("Failed to create post for feed %s: %v", feed.Name, err)
			}
			return err
		}
	}

	return nil
}

func parsePublishedDate(dateStr string) time.Time {
	t, err := time.Parse(time.RFC1123Z, dateStr)
	if err != nil {
		t, err = time.Parse(time.RFC1123, dateStr)
		if err != nil {
			return time.Now()
		}
	}
	return t
}
