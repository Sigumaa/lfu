package lfu

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Image struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}

type Client struct {
	baseURL  string
	userName string
	apiKey   string
}

func New(userName, apiKey string) *Client {
	return &Client{
		baseURL:  "https://ws.audioscrobbler.com/",
		userName: userName,
		apiKey:   apiKey,
	}
}

func buildURL(url, method, userName, apiKey string, page int) string {
	return fmt.Sprintf("%s2.0/?method=%s&user=%s&api_key=%s&page=%d&format=json", url, method, userName, apiKey, page)
}

func (c *Client) get(ctx context.Context, url string, result any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}
	return nil
}
