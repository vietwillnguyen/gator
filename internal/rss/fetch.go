package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"gator/internal/models"
	"io"
	"net/http"
	"time"
)

// FetchFeed retrieves and parses an RSS feed from the given URL
func FetchFeed(ctx context.Context, feedURL string) (*models.RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	dataBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var rssFeed models.RSSFeed
	if err := xml.Unmarshal(dataBytes, &rssFeed); err != nil {
		return nil, fmt.Errorf("error parsing XML: %w", err)
	}

	return &rssFeed, nil
}
