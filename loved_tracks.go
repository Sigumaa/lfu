package lfu

import "context"

type LovedTracksData struct {
	LovedTracks LovedTracks `json:"lovedtracks"`
}

type LovedTracks struct {
	Track []Track `json:"track"`
	Attr  Attr    `json:"@attr"`
}

type Track struct {
	Artist     Artist     `json:"artist"`
	Date       Date       `json:"date"`
	Mbid       string     `json:"mbid"`
	URL        string     `json:"url"`
	Name       string     `json:"name"`
	Image      []Image    `json:"image"`
	Streamable Streamable `json:"streamable"`
}

type Streamable struct {
	Fulltrack string `json:"fulltrack"`
	Text      string `json:"#text"`
}

type Artist struct {
	URL  string `json:"url"`
	Name string `json:"name"`
	Mbid string `json:"mbid"`
}

type Date struct {
	UTS  string `json:"uts"`
	Text string `json:"#text"`
}

func (c *Client) LovedTracks(ctx context.Context, page int) (*LovedTracksData, error) {
	url := buildURL(c.baseURL, "user.getlovedtracks", c.userName, c.apiKey, page)
	var result LovedTracksData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
