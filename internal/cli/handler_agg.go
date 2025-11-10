package cli

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"gator/internal/models"
	"gator/internal/rss"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

func HandlerAgg(s *models.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <time_between_reqs>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration %q: %w", cmd.Args[0], err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	// Scrape feeds

	// Run once immediately
	scrapeFeeds(s)

	for range ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *models.State) {
	feedToFetch, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting feed: %v\n", err)
		return
	}

	if err := s.Db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feedToFetch.ID,
		UpdatedAt: time.Now(),
	}); err != nil {
		fmt.Fprintf(os.Stderr, "error marking feed: %v\n", err)
		return
	}

	RSSFeed, err := rss.FetchFeed(context.Background(), feedToFetch.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetching %s: %v\n", feedToFetch.Url, err)
		return
	}

	fmt.Printf("Fetched: %s at %v\n", RSSFeed.Channel.Title, time.Now())
	fmt.Printf("Found %d posts(s)\n", len(RSSFeed.Channel.Items))
	numPostsIgnored := 0
	numPostsAdded := 0

	for _, item := range RSSFeed.Channel.Items {

		// if post already exists, ignore.
		_, err := s.Db.GetPostByURL(context.Background(), item.Link)
		if err != sql.ErrNoRows {
			// fmt.Printf("Post '%s' already exists, ignoring\n", item.Title)
			numPostsIgnored += 1
			continue
		}

		// attempt parse on publish date
		parsedPublishDate, err := parsePubDate(item.PublishedAt)
		if err != nil {
			// log and skip or fallback to time.Now()
			fmt.Printf("Unable to parse publish date for post titled '%s', ignoring\n", item.Title)
			continue
		}

		// save post to database
		_, err = s.Db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "", // store description if it is not empty
			},
			PublishedAt: sql.NullTime{
				Time:  parsedPublishDate,
				Valid: true, // always store publish date
			},
			FeedID: feedToFetch.ID,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating post: %v\n", err)
			continue
		}
		numPostsAdded += 1
	}
	fmt.Printf("New Posts Added: %d, Posts Ignored: %d\n", numPostsAdded, numPostsIgnored)
}

var timeLayouts = []string{
	time.RFC1123Z,                         // "Mon, 02 Jan 2006 15:04:05 -0700"
	time.RFC1123,                          // "Mon, 02 Jan 2006 15:04:05 MST"
	time.RFC822Z,                          // "02 Jan 06 15:04 -0700"
	time.RFC822,                           // "02 Jan 06 15:04 MST"
	time.RFC850,                           // "Monday, 02-Jan-06 15:04:05 MST"
	time.RFC3339,                          // ISO 8601
	"Mon, 02 Jan 2006 15:04:05 -0700 MST", // some feeds include both
}

func parsePubDate(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	for _, layout := range timeLayouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	// optional: attempt common cleanup (remove commas, collapse spaces, etc.) then retry
	return time.Time{}, fmt.Errorf("unrecognized pubDate format: %q", s)
}
