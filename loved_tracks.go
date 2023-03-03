package lfu

import "context"

type LovedTracksData struct {
	LovedTracks LovedTracks `json:"lovedtracks"`
}

type LovedTracks struct {
	Track []LovedTrack `json:"track"`
	Attr  Attr         `json:"@attr"`
}

type LovedTrack struct {
	Artist     Artist     `json:"artist"`
	Date       Date       `json:"date"`
	Mbid       string     `json:"mbid"`
	URL        string     `json:"url"`
	Name       string     `json:"name"`
	Image      []Image    `json:"image"`
	Streamable Streamable `json:"streamable"`
}

// LovedTracks Get the last 50 tracks loved by a user.
// Supported options:  limit, page
func (c *Client) LovedTracks(ctx context.Context, opts ...RequestOption) (*LovedTracksData, error) {
	url := buildURL(c.baseURL, "user.getlovedtracks", c.userName, c.apiKey)
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		url += "&" + params
	}
	var result LovedTracksData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
