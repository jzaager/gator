package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	const wagslaneURL = "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), wagslaneURL)
	if err != nil {
		return fmt.Errorf("Couldn't fetch feed: %w", err)
	}

	fmt.Printf("Feed: %v\n", feed)
	return nil
}
