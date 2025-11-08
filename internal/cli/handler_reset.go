package cli

import (
	"context"
	"fmt"
	"gator/internal/models"
)

// HandlerReset resets the database
func HandlerReset(s *models.State, cmd Command) error {
	_, err := s.Db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error on database reset: %w", err)
	}
	fmt.Printf("database reset successful\n")
	return nil
}
