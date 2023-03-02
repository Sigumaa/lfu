package lfu

import "context"

type InfoData struct {
	User UserInfo `json:"user"`
}

type UserInfo struct {
	Name        string         `json:"name"`
	Age         string         `json:"age"`
	Subscriber  string         `json:"subscriber"`
	Realname    string         `json:"realname"`
	Bootstrap   string         `json:"bootstrap"`
	Playcount   string         `json:"playcount"`
	ArtistCount string         `json:"artist_count"`
	Playlists   string         `json:"playlists"`
	TrackCount  string         `json:"track_count"`
	AlbumCount  string         `json:"album_count"`
	Image       []Image        `json:"image"`
	Registered  InfoRegistered `json:"registered"`
	Country     string         `json:"country"`
	Gender      string         `json:"gender"`
	URL         string         `json:"url"`
	Type        string         `json:"type"`
}

type InfoRegistered struct {
	Unixtime string `json:"unixtime"`
	Text     int    `json:"#text"`
}

func (c *Client) Info(ctx context.Context, page int) (*InfoData, error) {
	url := buildURL(c.baseURL, "user.getinfo", c.userName, c.apiKey, page)
	var result InfoData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
