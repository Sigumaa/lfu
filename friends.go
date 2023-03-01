package lfu

import "context"

type FriendsData struct {
	Friends Friends `json:"friends"`
}

type Friends struct {
	Attr Attr   `json:"@attr"`
	User []User `json:"user"`
}

type Attr struct {
	User       string `json:"user"`
	TotalPages string `json:"totalPages"`
	Page       string `json:"page"`
	Total      string `json:"total"`
	PerPage    string `json:"perPage"`
}

type Image struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}

type Registered struct {
	Unixtime string `json:"unixtime"`
	Text     string `json:"#text"`
}

type User struct {
	Name       string     `json:"name"`
	Url        string     `json:"url"`
	Country    string     `json:"country"`
	Playlists  string     `json:"playlists"`
	Playcount  string     `json:"playcount"`
	Image      []Image    `json:"image"`
	Registered Registered `json:"registered"`
	Realname   string     `json:"realname"`
	Subscriber string     `json:"subscriber"`
	Bootstrap  string     `json:"bootstrap"`
	Type       string     `json:"type"`
}

func (c *Client) Friends(ctx context.Context) (*FriendsData, error) {
	url := buildURL(c.baseURL, "user.getfriends", c.userName, c.apiKey)
	var result FriendsData
	err := c.get(ctx, url, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
