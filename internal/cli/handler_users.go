package cli

import (
	"context"
	"fmt"
	"gator/internal/models"
)

// HandlerUsers resets the database
func HandlerUsers(s *models.State, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching users: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("No feeds found")
		return nil
	}

	for _, user := range users {
		out := user.Name
		if s.Config.CurrentUserName == user.Name {
			out = out + " (current)"
		}
		fmt.Printf("* %s\n", out)
	}

	return nil
}
