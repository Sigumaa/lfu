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
	Streamable Streamable   `json:"streamable"`
	Mbid       string       `json:"mbid"`
	Name       string       `json:"name"`
	Image      []Image      `json:"image"`
	Artist     Artist       `json:"artist"`
	Url        string       `json:"url"`
	Duration   string       `json:"duration"`
	Attr       TopTrackAttr `json:"@attr"`
	PlayCount  string       `json:"playcount"`
}

type TopTrackAttr struct {
	Rank string `json:"rank"`
}

func (c *Client) TopTracks(ctx context.Context, page int) (*TopTracksData, error) {
	url := buildURL(c.baseURL, "user.gettoptracks", c.userName, c.apiKey, page)
	var result TopTracksData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
