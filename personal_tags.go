package lfu

import "context"

type PersonalTagsData struct {
	Taggings Taggings `json:"taggings"`
}

type Taggings struct {
	Artists TagArtists `json:"artists"`
	Attr    TagAttr    `json:"@attr"`
}

type TagArtists struct {
	Artist []TagArtist `json:"artist"`
}

type TagArtist struct {
	Name       string  `json:"name"`
	Mbid       string  `json:"mbid"`
	URL        string  `json:"url"`
	Streamable string  `json:"streamable"`
	Image      []Image `json:"image"`
}

type TagAttr struct {
	User       string `json:"user"`
	Tag        string `json:"tag"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

// PersonalTags Get the user's personal tags.
// Supported options:  limit, page
func (c *Client) PersonalTags(ctx context.Context, tag, taggingtype string, opts ...RequestOption) (*PersonalTagsData, error) {
	url := buildURL(c.baseURL, "user.getpersonaltags", c.userName, c.apiKey) + "&tag=" + tag + "&taggingtype=" + taggingtype
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		url += "&" + params
	}
	var result PersonalTagsData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
