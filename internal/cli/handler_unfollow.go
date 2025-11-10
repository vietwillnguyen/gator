package cli

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/internal/models"
)

func HandlerUnfollow(s *models.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("command usage: unfollow <url>")
	}

	// get feed id by URL
	url := cmd.Args[0]
	feed, err := s.Db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed by url: %w", err)
	}

	_, err = s.Db.DeleteFeedFollowsForUser(context.Background(), database.DeleteFeedFollowsForUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error deleting feed follows for user: %w", err)
	}
	fmt.Printf("Unfollowed %s\n", url)

	return nil
}
