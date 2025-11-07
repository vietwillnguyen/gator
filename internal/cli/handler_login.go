package cli

import (
	"fmt"
	"gator/internal/models"
)

// handlerLogin sets the current user in the config
func HandlerLogin(s *models.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login command requires a username argument")
	}

	username := cmd.Args[0]

	// Update the config with the new username
	s.Config.CurrentUserName = username

	// Save the updated config
	err := s.Config.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user: %v, error: %w", username, err)
	}

	fmt.Printf("User has been set to: %s\n", username)
	return nil
}
