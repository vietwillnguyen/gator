package cli

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/internal/models"
	"gator/internal/rss"
	"time"

	"github.com/google/uuid"
)

// HandlerAddFeed creates a feed in the postgres db
func HandlerAddFeed(s *models.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 { // Only need URL now
		return fmt.Errorf("command usage: addfeed <url>")
	}
	url := cmd.Args[0]

	// Add timeout for feed fetching
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rssFeed, err := rss.FetchFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      rssFeed.Channel.Title, // Use actual feed name
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	_, err = s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Printf("Added feed: %s\n", feed.Name)
	return nil
}
