package lfu

import (
	"context"
)

type RecentTracksData struct {
	RecentTracks RecentTracks `json:"recenttracks"`
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
	Attr       RecentTrackAttr   `json:"@attr"`
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
	NowPlaying string `json:"nowplaying"`
}

// RecentTracks Get a list of the recent tracks listened to by this user. Also includes the currently playing track with the nowplaying="true" attribute if the user is currently listening.
// Supported options:  limit, page, from, extended, to
func (c *Client) RecentTracks(ctx context.Context, opts ...RequestOption) (*RecentTracksData, error) {
	url := buildURL(c.baseURL, "user.getrecenttracks", c.userName, c.apiKey)
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		url += "&" + params
	}
	var result RecentTracksData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
