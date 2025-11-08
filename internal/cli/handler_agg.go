package cli

import (
	"context"
	"encoding/xml"
	"fmt"
	"gator/internal/models"
	"gator/internal/utils"
	"io"
	"net/http"
	"time"
)

// func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error)

func fetchFeed(ctx context.Context, feedURL string) (*models.RSSFeed, error) {
	// 1. Create request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	req.Header.Set("User-Agent", `gator`)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	fmt.Printf("created request: %s", utils.ToJSON(req))

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

func HandlerAgg(s *models.State, cmd Command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	fmt.Printf("rssFeed: %s\n", utils.ToJSON(rssFeed))
	return nil
}
