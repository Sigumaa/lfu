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
	Streamable string        `json:"streamable"`
	Image      []Image       `json:"image"`
	Mbid       string        `json:"mbid"`
	URL        string        `json:"url"`
	PlayCount  string        `json:"playcount"`
	Attr       TopArtistAttr `json:"@attr"`
	Name       string        `json:"name"`
}

type TopArtistAttr struct {
	Rank string `json:"rank"`
}

func (c *Client) TopArtists(ctx context.Context, page int) (*TopArtistsData, error) {
	url := buildURL(c.baseURL, "user.gettopartists", c.userName, c.apiKey, page)
	var result TopArtistsData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
