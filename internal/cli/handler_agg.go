package cli

import (
	"context"
	"encoding/xml"
	"fmt"
	"gator/internal/database"
	"gator/internal/models"
	"io"
	"net/http"
	"os"
	"time"
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

	// Run once immediately
	scrapeFeeds(s)

	for range ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*models.RSSFeed, error) {
	// 1. Create request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	req.Header.Set("User-Agent", `gator`)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// 2. Create client
	client := &http.Client{
		Timeout: 10 * time.Second, // Good practice to set a timeout
	}

	// 3. Make request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	// 4. Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// 5. Read all data from response body
	dataBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// 6. Unmarshal XML into RSSFeed struct
	var rssFeed models.RSSFeed
	err = xml.Unmarshal(dataBytes, &rssFeed)
	if err != nil {
		return nil, fmt.Errorf("error parse XML: %w", err)
	}

	return &rssFeed, nil
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

	rssFeed, err := fetchFeed(context.Background(), feedToFetch.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetching %s: %v\n", feedToFetch.Url, err)
		return
	}

	fmt.Printf("Fetched: %s at %v\n", rssFeed.Channel.Title, time.Now())
}
