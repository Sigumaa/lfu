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
	Artist    Artist   `json:"artist"`
	Image     []Image  `json:"image"`
	Mbid      string   `json:"mbid"`
	URL       string   `json:"url"`
	PlayCount string   `json:"playcount"`
	Attr      RankAttr `json:"@attr"`
	Name      string   `json:"name"`
}

// TopAlbums Get the top albums listened to by a user. You can stipulate a time period. Sends the overall chart by default.
// Supported options:  limit,page,period
func (c *Client) TopAlbums(ctx context.Context, opts ...RequestOption) (*TopAlbumsData, error) {
	url := buildURL(c.baseURL, "user.gettopalbums", c.userName, c.apiKey)
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		url += "&" + params
	}
	var result TopAlbumsData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
