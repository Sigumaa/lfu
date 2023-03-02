package lfu

import "context"

type TopTagsData struct {
	TopTags TopTags `json:"toptags"`
}

type TopTags struct {
	Tag  []TopTag    `json:"tag"`
	Attr TopTagsAttr `json:"@attr"`
}

type TopTag struct {
	Name  string `json:"name"`
	Count string `json:"count"`
	URL   string `json:"url"`
}

type TopTagsAttr struct {
	User string `json:"user"`
}

func (c *Client) TopTags(ctx context.Context, page int) (*TopTagsData, error) {
	url := buildURL(c.baseURL, "user.gettoptags", c.userName, c.apiKey, page)
	var result TopTagsData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
