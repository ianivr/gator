package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ianivr/gator/internal/database"

	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("addfeed command requires a name and url argument")
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
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
		UserID:    user.ID,
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

func handleFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("follow command requires a feed url argument")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to get feed by url: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	fmt.Printf("User %s is now following feed %s.\n", user.Name, feed.Name)
	return nil
}

func handleFollowing(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get feed follows for user: %w", err)
	}

	fmt.Printf("User %s is following:\n", user.Name)
	for _, follow := range follows {
		fmt.Printf("- '%s'\n", follow.FeedName)
	}
	return nil
}

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("unfollow command requires a feed url argument")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("User %s has unfollowed feed %s.\n", user.Name, feed.Name)
	return nil
}
