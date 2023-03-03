package lfu

import "context"

type WeeklyTrackChartData struct {
	WeeklyTrackChart WeeklyTrackChart `json:"weeklytrackchart"`
}

type WeeklyTrackChart struct {
	Track []WeeklyTrack `json:"track"`
	Attr  ChartAttr     `json:"@attr"`
}

type WeeklyTrack struct {
	Artist    ChartArtist `json:"artist"`
	Image     []Image     `json:"image"`
	Mbid      string      `json:"mbid"`
	URL       string      `json:"url"`
	Name      string      `json:"name"`
	Attr      RankAttr    `json:"@attr"`
	PlayCount string      `json:"playcount"`
}

// WeeklyTrackChart Get a track chart for a user profile, for a given date range. If no date range is supplied, it will return the most recent track chart for this user.
// Supported options:  from, to
func (c *Client) WeeklyTrackChart(ctx context.Context, opts ...RequestOption) (*WeeklyTrackChartData, error) {
	url := buildURL(c.baseURL, "user.getweeklytrackchart", c.userName, c.apiKey)
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		url += "&" + params
	}
	var result WeeklyTrackChartData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
