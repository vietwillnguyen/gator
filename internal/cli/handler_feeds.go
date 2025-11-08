package cli

import (
	"context"
	"fmt"
	"gator/internal/models"
)

func HandlerFeeds(s *models.State, cmd Command) error {
	feeds, err := s.Db.GetFeedsWithUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found")
		return nil
	}

	fmt.Printf("Found %d feed(s):\n\n", len(feeds))

	for i, feed := range feeds {
		fmt.Printf("%d) Name:  %s\n", i+1, feed.Name)
		fmt.Printf("   URL: %s\n", feed.Url)
		fmt.Printf("   Created by: %s\n", feed.UserName)
		fmt.Println()
	}

	return nil
}
