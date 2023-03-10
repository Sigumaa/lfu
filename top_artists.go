package lfu

import "context"

type TopArtistsData struct {
	TopArtists TopArtists `json:"topartists"`
}

type TopArtists struct {
	Artist []TopArtist `json:"artist"`
	Attr   Attr        `json:"@attr"`
}

type TopArtist struct {
	Streamable string   `json:"streamable"`
	Image      []Image  `json:"image"`
	Mbid       string   `json:"mbid"`
	URL        string   `json:"url"`
	PlayCount  string   `json:"playcount"`
	Attr       RankAttr `json:"@attr"`
	Name       string   `json:"name"`
}

// TopArtists Get the top artists listened to by a user. You can stipulate a time period. Sends the overall chart by default.
// Supported options:  limit, page, period
func (c *Client) TopArtists(ctx context.Context, opts ...RequestOption) (*TopArtistsData, error) {
	url := buildURL(c.baseURL, "user.gettopartists", c.userName, c.apiKey)
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		url += "&" + params
	}
	var result TopArtistsData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
