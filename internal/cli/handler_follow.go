package cli

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/internal/models"
	"gator/internal/utils"
	"time"

	"github.com/google/uuid"
)

// Have the current user follow a feed by url
// additionally create the feed follow record.
func HandlerFollow(s *models.State, cmd Command, user database.User) error {
	// Handle usage.
	if len(cmd.Args) != 1 {
		return fmt.Errorf("command usage: follow <url>")
	}

	url := cmd.Args[0]

	// Look up feed
	feed, err := s.Db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error get feed: %w", err)
	}

	// Create record for feed and follow
	feedFollowRows, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error create feed follow: %w", err)
	}
	// Print name of feed and current user once record.
	fmt.Printf("feed follow created: %s", utils.ToJSON(feedFollowRows))
	return nil
}
