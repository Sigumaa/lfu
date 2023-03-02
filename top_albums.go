package lfu

import "context"

type TopAlbumsData struct {
	TopAlbums Topalbums `json:"topalbums"`
}

type Topalbums struct {
	Album []TopAlbum `json:"album"`
	Attr  Attr       `json:"@attr"`
}

type TopAlbum struct {
	Artist    Artist       `json:"artist"`
	Image     []Image      `json:"image"`
	Mbid      string       `json:"mbid"`
	URL       string       `json:"url"`
	PlayCount string       `json:"playcount"`
	Attr      TopAlbumAttr `json:"@attr"`
	Name      string       `json:"name"`
}

type TopAlbumAttr struct {
	Rank string `json:"rank"`
}

func (c *Client) TopAlbums(ctx context.Context, page int) (*TopAlbumsData, error) {
	url := buildURL(c.baseURL, "user.gettopalbums", c.userName, c.apiKey, page)
	var result TopAlbumsData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
