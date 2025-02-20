package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jzaager/gator/internal/database"
)

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
	feedsFollowing, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Could not get curent user's followed feeds: %w", err)
	}

	if len(feedsFollowing) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("%s is following these feeds:\n", user.Name)
	for _, feed := range feedsFollowing {
		fmt.Printf("* %s\n", feed.FeedName)
	}

	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <feed_url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Could not get feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Could not create feed follow: %w", err)
	}

	fmt.Println("Feed follow created:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("=========================================================")

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <feed_url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Couldn't get feed: %w", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("Couldn't delete feed follow: %w", err)
	}

	fmt.Printf("%s successfully unfollowed\n", feed.Name)
	return nil
}
