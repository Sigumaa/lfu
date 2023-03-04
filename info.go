package lfu

import "context"

type InfoData struct {
	User UserInfo `json:"user"`
}

type UserInfo struct {
	Name        string         `json:"name"`
	Age         string         `json:"age"`
	Subscriber  string         `json:"subscriber"`
	RealName    string         `json:"realname"`
	Bootstrap   string         `json:"bootstrap"`
	PlayCount   string         `json:"playcount"`
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
	UnixTime string `json:"unixtime"`
	Text     int    `json:"#text"`
}

// Info Get information about a user profile.
// The last fm api documentation says user (Optional), but it is a lie!
func (c *Client) Info(ctx context.Context) (*InfoData, error) {
	url := buildURL(c.baseURL, "user.getinfo", c.userName, c.apiKey)
	var result InfoData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Name returns the name of the user.
func (i *InfoData) Name() string {
	return i.User.Name
}

// Age returns the age of the user.
func (i *InfoData) Age() string {
	return i.User.Age
}

// Subscriber returns the subscriber status of the user.
func (i *InfoData) Subscriber() string {
	return i.User.Subscriber
}

// RealName returns the real name of the user.
func (i *InfoData) RealName() string {
	return i.User.RealName
}

// Bootstrap returns the bootstrap status of the user.
func (i *InfoData) Bootstrap() string {
	return i.User.Bootstrap
}

// PlayCount returns the play count of the user.
func (i *InfoData) PlayCount() string {
	return i.User.PlayCount
}

// ArtistCount returns the artist count of the user.
func (i *InfoData) ArtistCount() string {
	return i.User.ArtistCount
}

// Playlists returns the playlists of the user.
func (i *InfoData) Playlists() string {
	return i.User.Playlists
}

// TrackCount returns the track count of the user.
func (i *InfoData) TrackCount() string {
	return i.User.TrackCount
}

// AlbumCount returns the album count of the user.
func (i *InfoData) AlbumCount() string {
	return i.User.AlbumCount
}

// Image returns the image of the user.
func (i *InfoData) Image() []Image {
	return i.User.Image
}

// ImageURL returns the image url of the user.
func (i *InfoData) ImageURL(size imageSize) string {
	return i.User.Image[size].Text
}

// Registered returns the registered info of the user.
func (i *InfoData) Registered() InfoRegistered {
	return i.User.Registered
}

// RegisteredUnixTime returns the registered unix time of the user.
func (i *InfoData) RegisteredUnixTime() string {
	return i.User.Registered.UnixTime
}

// RegisteredText returns the registered text of the user.
func (i *InfoData) RegisteredText() int {
	return i.User.Registered.Text
}

// GetUnixTime returns the unix time of the user.
func (r *InfoRegistered) GetUnixTime() string {
	return r.UnixTime
}

// GetText returns the text of the user.
func (r *InfoRegistered) GetText() int {
	return r.Text
}

// Country returns the country of the user.
func (i *InfoData) Country() string {
	return i.User.Country
}

// Gender returns the gender of the user.
func (i *InfoData) Gender() string {
	return i.User.Gender
}

// URL returns the url of the user.
func (i *InfoData) URL() string {
	return i.User.URL
}

// Type returns the type of the user.
func (i *InfoData) Type() string {
	return i.User.Type
}
