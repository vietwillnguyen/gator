package cli

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/internal/models"
	"strconv"
)

func HandlerBrowse(s *models.State, cmd Command, user database.User) error {
	limit := 2
	switch len(cmd.Args) {
	case 0:
		// use default
	case 1:
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("limit argument must be an integer")
		}
	default:
		return fmt.Errorf("usage: browse [limit] (default: 2)")
	}

	// func (q *Queries) GetPostsForUser(ctx context.Context, userID uuid.UUID) ([]Post, error) {
	posts, err := s.Db.GetPostsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}
	for i := 0; i < limit; i++ {
		post := posts[i]
		fmt.Printf("%d) Title:  %s\n", i+1, post.Title)
		fmt.Printf("   Date:  %v\n", post.PublishedAt)
		fmt.Printf("   Description:  %v\n", post.Description)
		fmt.Println()
	}

	return nil
}
