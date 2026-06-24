package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ianivr/gator/internal/database"
)

func handleBrowse(s *state, cmd command) error {
	limit := 2
	var err error
	if len(cmd.args) > 0 {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	posts, err := s.db.GetPostsByUser(context.Background(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("\n--- %s ---\n", post.Title)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Printf("Description: %s\n", post.Description)
		fmt.Printf("Published: %v\n", post.PublishedAt)
	}

	return nil
}
