package cli

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/internal/models"
)

func HandlerUnfollow(s *models.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("command usage: unfollow <url>")
	}

	url := cmd.Args[0]

	feed, err := s.Db.DeleteFeedFollowsForUser(context.Background(), database.DeleteFeedFollowsForUserParams{
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error deleting feed follows for user: %w", err)
	}

	return nil
}
