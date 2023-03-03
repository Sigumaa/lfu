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
