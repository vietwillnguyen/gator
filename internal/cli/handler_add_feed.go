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

// HandlerAddFeed creates a feed in the postgres db
func HandlerAddFeed(s *models.State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("addFeed command usage: addFeed <name> <url>")
	}

	user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	fmt.Printf("feed object created: %s", utils.ToJSON(feed))
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	return nil
}
