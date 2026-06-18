package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ianivr/gator/internal/database"

	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("addfeed command requires a name and url argument")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		Name:      cmd.args[0],
		Url:       cmd.args[1],
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	fmt.Printf("Feed %s has been created.\nData: %+v\n", feed.Name, feed)
	return nil
}

func handleFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %w", err)
	}

	fmt.Printf("Feeds:\n")
	for _, feed := range feeds {
		userName, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("failed to get user for feed %s: %w", feed.Name, err)
		}
		fmt.Printf("- %s (URL: %s, User: %s)\n", feed.Name, feed.Url, userName)
	}
	return nil
}
