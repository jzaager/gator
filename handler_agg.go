package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jzaager/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *state) {
	// get next nextFeed to fetch from DB
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch: ", err)
	}

	log.Println("Found a feed to fetch!")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	// mark feed as fetched
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	// fetch the feed using it's URL
	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
	}

	// iterate over items in feed & save posts to DB
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)

		postParams, err := getPostParams(item, feed)
		if err != nil {
			log.Printf("Couldn't get post params: %v", err)
		}

		_, err = db.CreatePost(context.Background(), postParams)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}

	log.Printf("Feed %s collected, %v posts found\n", feed.Name, len(feedData.Channel.Item))
}

func getPostParams(postData RSSItem, feed database.Feed) (database.CreatePostParams, error) {
	publishedAt := sql.NullTime{}
	if t, err := time.Parse(time.RFC1123Z, postData.PubDate); err == nil {
		publishedAt = sql.NullTime{
			Time:  t,
			Valid: true,
		}
	}

	description := sql.NullString{
		String: postData.Description,
		Valid:  postData.Description == "",
	}

	postParams := database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Title:       postData.Title,
		Url:         postData.Link,
		Description: description,
		PublishedAt: publishedAt,
		FeedID:      feed.ID,
	}

	return postParams, nil
}
