package lfu

import "context"

type WeeklyAlbumChartData struct {
	WeeklyAlbumChart WeeklyAlbumChart `json:"weeklyalbumchart"`
}

type WeeklyAlbumChart struct {
	Album []WeeklyAlbum `json:"album"`
	Attr  ChartAttr     `json:"@attr"`
}

type WeeklyAlbum struct {
	Artist    ChartArtist `json:"artist"`
	Mbid      string      `json:"mbid"`
	URL       string      `json:"url"`
	Name      string      `json:"name"`
	Attr      RankAttr    `json:"@attr"`
	PlayCount string      `json:"playcount"`
}

// WeeklyAlbumChart Get an album chart for a user profile, for a given date range. If no date range is supplied, it will return the most recent album chart for this user.
// Supported options:  from, to
func (c *Client) WeeklyAlbumChart(ctx context.Context, opts ...RequestOption) (*WeeklyAlbumChartData, error) {
	url := buildURL(c.baseURL, "user.getweeklyalbumchart", c.userName, c.apiKey)
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		url += "&" + params
	}
	var result WeeklyAlbumChartData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
