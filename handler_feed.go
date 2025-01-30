package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jzaager/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Usage: %s <name> <url>", cmd.Name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]
	newFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    currentUser.ID,
	})

	if err != nil {
		return fmt.Errorf("Couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(newFeed)
	fmt.Println()
	fmt.Println("==================================================")
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:			%s\n", feed.ID)
	fmt.Printf("* Created:		%v\n", feed.CreatedAt)
	fmt.Printf("* Updated:		%v\n", feed.UpdatedAt)
	fmt.Printf("* Name:			%s\n", feed.Name)
	fmt.Printf("* URL:			%s\n", feed.Url)
	fmt.Printf("* UserID:		%s\n", feed.UserID)
}
