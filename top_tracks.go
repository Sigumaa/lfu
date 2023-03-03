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

func (c *Client) TopTracks(ctx context.Context) (*TopTracksData, error) {
	url := buildURL(c.baseURL, "user.gettoptracks", c.userName, c.apiKey)
	var result TopTracksData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
