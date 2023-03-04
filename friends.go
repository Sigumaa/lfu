package lfu

import (
	"context"
)

type FriendsData struct {
	Friends Friends `json:"friends"`
}

type Friends struct {
	Attr Attr   `json:"@attr"`
	User []User `json:"user"`
}

type User struct {
	Name       string           `json:"name"`
	URL        string           `json:"url"`
	Country    string           `json:"country"`
	Playlists  string           `json:"playlists"`
	PlayCount  string           `json:"playcount"`
	Image      []Image          `json:"image"`
	Registered FriendRegistered `json:"registered"`
	RealName   string           `json:"realname"`
	Subscriber string           `json:"subscriber"`
	Bootstrap  string           `json:"bootstrap"`
	Type       string           `json:"type"`
}

type FriendRegistered struct {
	UnixTime string `json:"unixtime"`
	Text     string `json:"#text"`
}

// Friends Get a list of the user's friends on Last.fm.
// Supported options:  limit, page
func (c *Client) Friends(ctx context.Context, opts ...RequestOption) (*FriendsData, error) {
	url := buildURL(c.baseURL, "user.getfriends", c.userName, c.apiKey)
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		url += "&" + params
	}
	var result FriendsData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// User returns the name of the user whose friends are being returned.
func (f *FriendsData) User() string {
	return f.Friends.Attr.User
}

// TotalPages returns the total number of pages of friends.
func (f *FriendsData) TotalPages() string {
	return f.Friends.Attr.TotalPages
}

// Page returns the current page of friends.
func (f *FriendsData) Page() string {
	return f.Friends.Attr.Page
}

// PerPage returns the number of friends per page.
func (f *FriendsData) PerPage() string {
	return f.Friends.Attr.PerPage
}

// Total returns the total number of friends.
func (f *FriendsData) Total() string {
	return f.Friends.Attr.Total
}

// Len returns the number of friends.
func (f *FriendsData) Len() int {
	return len(f.Friends.User)
}

// Name returns the name of the friend at the given index.
// The index must be in the range [0, f.Len()).
func (f *FriendsData) Name(index int) (string, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return "", ErrOutOfBounds
	}
	return f.Friends.User[c].Name, nil
}

// URL returns the URL of the friend at the given index.
// The index must be in the range [0, f.Len()).
func (f *FriendsData) URL(index int) (string, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return "", ErrOutOfBounds
	}
	return f.Friends.User[c].URL, nil
}

// Country returns the country of the friend at the given index.
// The index must be in the range [0, f.Len()).
func (f *FriendsData) Country(index int) (string, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return "", ErrOutOfBounds
	}
	return f.Friends.User[c].Country, nil
}

// Playlists returns the number of playlists of the friend at the given index.
// The index must be in the range [0, f.Len()).
func (f *FriendsData) Playlists(index int) (string, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return "", ErrOutOfBounds
	}
	return f.Friends.User[c].Playlists, nil
}

// PlayCount returns the number of plays of the friend at the given index.
// The index must be in the range [0, f.Len()).
func (f *FriendsData) PlayCount(index int) (string, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return "", ErrOutOfBounds
	}
	return f.Friends.User[c].PlayCount, nil
}

// Image returns the []Image of the friend at the given index.
// The index must be in the range [0, f.Len()).
func (f *FriendsData) Image(index int) ([]Image, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return nil, ErrOutOfBounds
	}
	return f.Friends.User[c].Image, nil
}

// ImageURL returns the URL of the image of the friend at the given index.
// The index must be in the range [0, 4). small, medium, large, extralarge
func (f *FriendsData) ImageURL(index int, size imageSize) (string, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return "", ErrOutOfBounds
	}
	return f.Friends.User[c].Image[size].Text, nil
}

// Registered returns the registration date of the friend at the given index.
// The index must be in the range [0, f.Len()).
func (f *FriendsData) Registered(index int) (FriendRegistered, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return FriendRegistered{}, ErrOutOfBounds
	}
	return f.Friends.User[c].Registered, nil
}

// RegisteredUnixTime returns the unix time of the friend registration
// The index must be in the range [0, f.Len()).
func (f *FriendsData) RegisteredUnixTime(index int) (string, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return "", ErrOutOfBounds
	}
	return f.Friends.User[c].Registered.UnixTime, nil
}

// RegisteredText returns the text of the friend registration
// The index must be in the range [0, f.Len()).
func (f *FriendsData) RegisteredText(index int) (string, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return "", ErrOutOfBounds
	}
	return f.Friends.User[c].Registered.Text, nil
}

// GetUnixTime returns the unix time of the friend registration
func (r *FriendRegistered) GetUnixTime() string {
	return r.UnixTime
}

// GetText returns the text of the friend registration
// Example: 2020-12-27 05:08
func (r *FriendRegistered) GetText() string {
	return r.Text
}

// RealName returns the real name of the friend at the given index.
// The index must be in the range [0, f.Len()).
func (f *FriendsData) RealName(index int) (string, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return "", ErrOutOfBounds
	}
	return f.Friends.User[c].RealName, nil
}

// Subscriber returns the subscriber status of the friend at the given index.
// The index must be in the range [0, f.Len()).
func (f *FriendsData) Subscriber(index int) (string, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return "", ErrOutOfBounds
	}
	return f.Friends.User[c].Subscriber, nil
}

// Bootstrap returns the bootstrap status of the friend at the given index.
// The index must be in the range [0, f.Len()).
func (f *FriendsData) Bootstrap(index int) (string, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return "", ErrOutOfBounds
	}
	return f.Friends.User[c].Bootstrap, nil
}

// Type returns the type of the friend at the given index.
// The index must be in the range [0, f.Len()).
func (f *FriendsData) Type(index int) (string, error) {
	c := index
	if c < 0 || c >= f.Len() {
		return "", ErrOutOfBounds
	}
	return f.Friends.User[c].Type, nil
}
