package lfu

import "context"

type TopTagsData struct {
	TopTags TopTags `json:"toptags"`
}

type TopTags struct {
	Tag  []TopTag    `json:"tag"`
	Attr TopTagsAttr `json:"@attr"`
}

type TopTagsAttr struct {
	User string `json:"user"`
}

type TopTag struct {
	Name  string `json:"name"`
	Count string `json:"count"`
	URL   string `json:"url"`
}

// TopTags Get the top tags used by this user.
// Supported options:  limit
func (c *Client) TopTags(ctx context.Context, opts ...RequestOption) (*TopTagsData, error) {
	url := buildURL(c.baseURL, "user.gettoptags", c.userName, c.apiKey)
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		url += "&" + params
	}
	var result TopTagsData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
