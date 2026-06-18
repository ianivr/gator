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

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
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

func handleFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("follow command requires a feed url argument")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to get feed by url: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	fmt.Printf("User %s is now following feed %s.\n", currentUser.Name, feed.Name)
	return nil
}

func handleFollowing(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("failed to get feed follows for user: %w", err)
	}

	fmt.Printf("User %s is following:\n", currentUser.Name)
	for _, follow := range follows {
		fmt.Printf("- '%s'\n", follow.FeedName)
	}
	return nil
}
