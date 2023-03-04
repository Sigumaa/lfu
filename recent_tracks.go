package lfu

import (
	"context"
	"errors"
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

// NowPlayingTrack struct is a subset of RecentTrack struct.
type NowPlayingTrack struct {
	Artist     RecentTrackArtist
	Streamable string
	Image      []Image
	Mbid       string
	Album      RecentTrackAlbum
	Name       string
	URL        string
	Date       Date
}

// NowPlayingTrack returns the currently playing track.
func (c *Client) NowPlayingTrack(ctx context.Context) (*NowPlayingTrack, error) {
	url := buildURL(c.baseURL, "user.getrecenttracks", c.userName, c.apiKey)
	var result RecentTracksData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	var npt NowPlayingTrack
	if result.RecentTracks.Track[0].Attr.NowPlaying == "true" {
		npt.Artist = result.RecentTracks.Track[0].Artist
		npt.Streamable = result.RecentTracks.Track[0].Streamable
		npt.Image = result.RecentTracks.Track[0].Image
		npt.Mbid = result.RecentTracks.Track[0].Mbid
		npt.Album = result.RecentTracks.Track[0].Album
		npt.Name = result.RecentTracks.Track[0].Name
		npt.URL = result.RecentTracks.Track[0].URL
		npt.Date = result.RecentTracks.Track[0].Date
		return &npt, nil
	}
	return nil, errors.New("no track is currently playing")
}

// NowPlayingTrack returns the currently playing track.
func (r *RecentTracksData) NowPlayingTrack() (*NowPlayingTrack, error) {
	var npt NowPlayingTrack
	if r.RecentTracks.Track[0].Attr.NowPlaying == "true" {
		npt.Artist = r.RecentTracks.Track[0].Artist
		npt.Streamable = r.RecentTracks.Track[0].Streamable
		npt.Image = r.RecentTracks.Track[0].Image
		npt.Mbid = r.RecentTracks.Track[0].Mbid
		npt.Album = r.RecentTracks.Track[0].Album
		npt.Name = r.RecentTracks.Track[0].Name
		npt.URL = r.RecentTracks.Track[0].URL
		npt.Date = r.RecentTracks.Track[0].Date
		return &npt, nil
	}
	return nil, errors.New("no track is currently playing")
}

// ArtistName returns the name of the artist.
func (n *NowPlayingTrack) ArtistName() string {
	return n.Artist.Text
}

// ArtistMbid returns the Mbid of the artist.
func (n *NowPlayingTrack) ArtistMbid() string {
	return n.Artist.Mbid
}

// GetImageURL returns the URL of the image. small, medium, large, extralarge.
func (n *NowPlayingTrack) GetImageURL(size imageSize) string {
	return n.Image[size].Text
}

// AlbumName returns the name of the album.
func (n *NowPlayingTrack) AlbumName() string {
	return n.Album.Text
}

// AlbumMbid returns the Mbid of the album.
func (n *NowPlayingTrack) AlbumMbid() string {
	return n.Album.Mbid
}

// GetDateText returns the date of the track.
func (n *NowPlayingTrack) GetDateText() string {
	return n.Date.Text
}

// GetDateUTS returns the UTS of the track.
func (n *NowPlayingTrack) GetDateUTS() string {
	return n.Date.UTS
}
