package youtube

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Client struct {
	apiKeys         []string
	currentKeyIndex int
	mu              sync.Mutex
	httpClient      *http.Client
}

type SearchResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title       string    `json:"title"`
			Description string    `json:"description"`
			PublishedAt time.Time `json:"publishedAt"`
			Thumbnails  struct {
				Default struct {
					URL string `json:"url"`
				} `json:"default"`
				Medium struct {
					URL string `json:"url"`
				} `json:"medium"`
				High struct {
					URL string `json:"url"`
				} `json:"high"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

func NewClient(apiKeys []string) *Client {
	return &Client{
		apiKeys: apiKeys,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) getNextAPIKey() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.currentKeyIndex = (c.currentKeyIndex + 1) % len(c.apiKeys)
	return c.apiKeys[c.currentKeyIndex]
}

func (c *Client) FetchVideos(query string, publishedAfter time.Time) (*SearchResponse, error) {
	baseURL := "https://youtube.googleapis.com/youtube/v3/search"

	for i := 0; i < len(c.apiKeys); i++ {
		apiKey := c.apiKeys[c.currentKeyIndex]
		encodedPublishedAfter := url.QueryEscape(publishedAfter.Format(time.RFC3339))

		url := fmt.Sprintf("%s?part=snippet&q=%s&type=video&order=date&maxResults=25&key=%s&publishedAfter=%s",
			baseURL, query, apiKey, encodedPublishedAfter)
		resp, err := c.httpClient.Get(url)
		if err != nil {
			c.getNextAPIKey()
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusForbidden {
			c.getNextAPIKey()
			continue
		}

		var searchResp SearchResponse
		if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
			return nil, err
		}
		return &searchResp, nil
	}

	return nil, fmt.Errorf("all API keys exhausted")
}
