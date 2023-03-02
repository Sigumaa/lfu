package lfu

import "context"

type LovedTracks struct {
	Track []Track `json:"track"`
	Attr  Attr    `json:"@attr"`
}

type Track struct {
	Artist     Artist     `json:"artist"`
	Date       Date       `json:"date"`
	Mbid       string     `json:"mbid"`
	Url        string     `json:"url"`
	Name       string     `json:"name"`
	Image      []Image    `json:"image"`
	Streamable Streamable `json:"streamable"`
}

type Streamable struct {
	Fulltrack string `json:"fulltrack"`
	Text      string `json:"#text"`
}

type Artist struct {
	Url  string `json:"url"`
	Name string `json:"name"`
	Mbid string `json:"mbid"`
}

type Date struct {
	Uts  string `json:"uts"`
	Text string `json:"#text"`
}

type LovedTracksData struct {
	LovedTracks LovedTracks `json:"lovedtracks"`
}

func (c *Client) LovedTracks(ctx context.Context, page int) (*LovedTracksData, error) {
	url := buildURL(c.baseURL, "user.getlovedtracks", c.userName, c.apiKey, page)
	var result LovedTracksData
	err := c.get(ctx, url, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
