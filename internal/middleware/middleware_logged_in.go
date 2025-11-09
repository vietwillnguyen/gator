package middleware

import (
	"context"
	"fmt"
	"gator/internal/cli"
	"gator/internal/database"
	"gator/internal/models"

	// import the driver, but you don't use it directly anywhere in your
	// code. The underscore tells Go that you're importing it for
	// its side effects, not because you need to use it.
	_ "github.com/lib/pq"
)

// middlewareLoggedIn checks if a user is logged in and passes the user to the handler
func LoggedIn(handler func(s *models.State, cmd cli.Command, user database.User) error) func(*models.State, cli.Command) error {
	return func(s *models.State, cmd cli.Command) error {
		// Get the current user from config
		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("user not logged in or user doesn't exist: %w", err)
		}

		// Call the original handler with the user
		return handler(s, cmd, user)
	}
}
