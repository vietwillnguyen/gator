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

// HandlerRegister creates a user in the postgres db
func HandlerRegister(s *models.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("register command requires a username argument")
	}
	username := cmd.Args[0]
	fmt.Printf("Register: register user: %s\n", username)
	user, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	fmt.Printf("user object created: %s", utils.ToJSON(user))
	if err != nil {
		return fmt.Errorf("failed to create user: %v, error: %w", username, err)
	}
	// Also perform the login, setting the user in the config
	err = HandlerLogin(s, cmd)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
