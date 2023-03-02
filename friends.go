package lfu

import "context"

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
	Playcount  string           `json:"playcount"`
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

func (c *Client) Friends(ctx context.Context, page int) (*FriendsData, error) {
	url := buildURL(c.baseURL, "user.getfriends", c.userName, c.apiKey, page)
	var result FriendsData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
