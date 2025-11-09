package cli

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/internal/models"
)

// Print all the names of the feeds the current user is following.
func HandlerFollowing(s *models.State, cmd Command, user database.User) error {
	// Get feed follows for user
	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching following: %w", err)
	}
	// Handle none found
	if len(feedFollows) == 0 {
		fmt.Println("No feedFollows found")
		return nil
	}

	fmt.Printf("User following %d feeds(s):\n\n", len(feedFollows))
	for i, feed := range feedFollows {
		fmt.Printf("%d) %s\n", i+1, feed.FeedName)
	}
	fmt.Println()

	return nil
}
