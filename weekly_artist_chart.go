package lfu

import "context"

type WeeklyArtistChartData struct {
	WeeklyArtistChart WeeklyArtistChart `json:"weeklyartistchart"`
}

type WeeklyArtistChart struct {
	Artist []WeeklyArtist `json:"artist"`
	Attr   ChartAttr      `json:"@attr"`
}

type WeeklyArtist struct {
	Mbid      string   `json:"mbid"`
	Url       string   `json:"url"`
	Name      string   `json:"name"`
	Attr      RankAttr `json:"@attr"`
	PlayCount string   `json:"playcount"`
}

// WeeklyArtistChart Get an artist chart for a user profile, for a given date range. If no date range is supplied, it will return the most recent artist chart for this user.
// Supported options:  from, to
func (c *Client) WeeklyArtistChart(ctx context.Context, opts ...RequestOption) (*WeeklyArtistChartData, error) {
	url := buildURL(c.baseURL, "user.getweeklyartistchart", c.userName, c.apiKey)
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		url += "&" + params
	}
	var result WeeklyArtistChartData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
