package lfu

import "context"

type TopTracksData struct {
	TopTracks TopTracks `json:"toptracks"`
}

type TopTracks struct {
	Track []TopTrack `json:"track"`
	Attr  Attr       `json:"@attr"`
}

type TopTrack struct {
	Streamable Streamable `json:"streamable"`
	Mbid       string     `json:"mbid"`
	Name       string     `json:"name"`
	Image      []Image    `json:"image"`
	Artist     Artist     `json:"artist"`
	Url        string     `json:"url"`
	Duration   string     `json:"duration"`
	Attr       RankAttr   `json:"@attr"`
	PlayCount  string     `json:"playcount"`
}

// TopTracks Get the top tracks listened to by a user. You can stipulate a time period. Sends the overall chart by default.
// Supported options:  limit, page, period
func (c *Client) TopTracks(ctx context.Context, opts ...RequestOption) (*TopTracksData, error) {
	url := buildURL(c.baseURL, "user.gettoptracks", c.userName, c.apiKey)
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		url += "&" + params
	}
	var result TopTracksData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
