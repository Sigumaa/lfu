package lfu

import "context"

type RecentTracksData struct {
	Recenttracks RecentTracks `json:"recenttracks"`
}

type RecentTracks struct {
	Track []RecentTrack `json:"track"`
	Attr  Attr          `json:"@attr"`
}

type RecentTrack struct {
	Artist     RecentTrackArtist `json:"artist"`
	Streamable string            `json:"streamable"`
	Image      []Image           `json:"image"`
	Mbid       string            `json:"mbid"`
	Album      RecentTrackAlbum  `json:"album"`
	Name       string            `json:"name"`
	Attr       RecentTrackAttr   `json:"@attr,omitempty"`
	URL        string            `json:"url"`
	Date       Date              `json:"date"`
}

type RecentTrackArtist struct {
	Mbid string `json:"mbid"`
	Text string `json:"#text"`
}

type RecentTrackAlbum struct {
	Mbid string `json:"mbid"`
	Text string `json:"#text"`
}

type RecentTrackAttr struct {
	Nowplaying string `json:"nowplaying"`
}

func (c *Client) RecentTracks(ctx context.Context, page int) (*RecentTracksData, error) {
	url := buildURL(c.baseURL, "user.getrecenttracks", c.userName, c.apiKey, page)
	var result RecentTracksData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
