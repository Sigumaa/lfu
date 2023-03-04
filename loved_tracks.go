package lfu

import "context"

type LovedTracksData struct {
	LovedTracks LovedTracks `json:"lovedtracks"`
}

type LovedTracks struct {
	Track []LovedTrack `json:"track"`
	Attr  Attr         `json:"@attr"`
}

type LovedTrack struct {
	Artist     Artist     `json:"artist"`
	Date       Date       `json:"date"`
	Mbid       string     `json:"mbid"`
	URL        string     `json:"url"`
	Name       string     `json:"name"`
	Image      []Image    `json:"image"`
	Streamable Streamable `json:"streamable"`
}

// LovedTracks Get the last 50 tracks loved by a user.
// Supported options:  limit, page
func (c *Client) LovedTracks(ctx context.Context, opts ...RequestOption) (*LovedTracksData, error) {
	url := buildURL(c.baseURL, "user.getlovedtracks", c.userName, c.apiKey)
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		url += "&" + params
	}
	var result LovedTracksData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// User returns the name of the user who loved the tracks.
func (l *LovedTracksData) User() string {
	return l.LovedTracks.Attr.User
}

// TotalPages returns the total number of pages of loved tracks.
func (l *LovedTracksData) TotalPages() string {
	return l.LovedTracks.Attr.TotalPages
}

// Page returns the current page of loved tracks.
func (l *LovedTracksData) Page() string {
	return l.LovedTracks.Attr.Page
}

// PerPage returns the number of loved tracks per page.
func (l *LovedTracksData) PerPage() string {
	return l.LovedTracks.Attr.PerPage
}

// Total returns the total number of loved tracks.
func (l *LovedTracksData) Total() string {
	return l.LovedTracks.Attr.Total
}

// Tracks returns the list of loved tracks.
func (l *LovedTracksData) Tracks() []LovedTrack {
	return l.LovedTracks.Track
}

// Len returns the number of loved tracks.
func (l *LovedTracksData) Len() int {
	return len(l.LovedTracks.Track)
}

// ArtistName returns the name of the artist who created the track.
// The index must be range [0, l.Len()).
func (l *LovedTracksData) ArtistName(index int) (string, error) {
	if index < 0 || index >= l.Len() {
		return "", ErrOutOfBounds
	}
	return l.LovedTracks.Track[index].Artist.Name, nil
}

// ArtistMbid returns the Mbid of the artist who created the track.
// The index must be range [0, l.Len()).
func (l *LovedTracksData) ArtistMbid(index int) (string, error) {
	if index < 0 || index >= l.Len() {
		return "", ErrOutOfBounds
	}
	return l.LovedTracks.Track[index].Artist.Mbid, nil
}

// ArtistURL returns the URL of the artist who created the track.
// The index must be range [0, l.Len()).
func (l *LovedTracksData) ArtistURL(index int) (string, error) {
	if index < 0 || index >= l.Len() {
		return "", ErrOutOfBounds
	}
	return l.LovedTracks.Track[index].Artist.URL, nil
}

// DateUTS returns the date the track was loved in Unix timestamp format.
// The index must be range [0, l.Len()).
func (l *LovedTracksData) DateUTS(index int) (string, error) {
	if index < 0 || index >= l.Len() {
		return "", ErrOutOfBounds
	}
	return l.LovedTracks.Track[index].Date.UTS, nil
}

// DateText returns the date the track was loved in text format.
// The index must be range [0, l.Len()).
func (l *LovedTracksData) DateText(index int) (string, error) {
	if index < 0 || index >= l.Len() {
		return "", ErrOutOfBounds
	}
	return l.LovedTracks.Track[index].Date.Text, nil
}

// Mbid returns the Mbid of the track.
// The index must be range [0, l.Len()).
func (l *LovedTracksData) Mbid(index int) (string, error) {
	if index < 0 || index >= l.Len() {
		return "", ErrOutOfBounds
	}
	return l.LovedTracks.Track[index].Mbid, nil
}

// URL returns the URL of the track.
// The index must be range [0, l.Len()).
func (l *LovedTracksData) URL(index int) (string, error) {
	if index < 0 || index >= l.Len() {
		return "", ErrOutOfBounds
	}
	return l.LovedTracks.Track[index].URL, nil
}

// Name returns the name of the track.
// The index must be range [0, l.Len()).
func (l *LovedTracksData) Name(index int) (string, error) {
	if index < 0 || index >= l.Len() {
		return "", ErrOutOfBounds
	}
	return l.LovedTracks.Track[index].Name, nil
}

// Image returns the list of images of the track.
// The index must be range [0, l.Len()).
func (l *LovedTracksData) Image(index int) ([]Image, error) {
	if index < 0 || index >= l.Len() {
		return nil, ErrOutOfBounds
	}
	return l.LovedTracks.Track[index].Image, nil
}

// ImageURL returns the URL of the image of the track.
// The index must be range [0, 4). small, medium, large, extralarge.
func (l *LovedTracksData) ImageURL(index int, size imageSize) (string, error) {
	if index < 0 || index >= l.Len() {
		return "", ErrOutOfBounds
	}
	return l.LovedTracks.Track[index].Image[size].Text, nil
}

// Streamable returns the streamable status of the track.
// The index must be range [0, l.Len()).
func (l *LovedTracksData) Streamable(index int) (Streamable, error) {
	if index < 0 || index >= l.Len() {
		return Streamable{}, ErrOutOfBounds
	}
	return l.LovedTracks.Track[index].Streamable, nil
}
