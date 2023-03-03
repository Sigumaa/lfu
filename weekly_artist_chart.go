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

func (c *Client) WeeklyArtistChart(ctx context.Context, page int) (*WeeklyArtistChartData, error) {
	url := buildURL(c.baseURL, "user.getweeklyartistchart", c.userName, c.apiKey, page)
	var result WeeklyArtistChartData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
