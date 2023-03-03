package lfu

import "context"

type WeeklyArtistChartData struct {
	WeeklyArtistChart WeeklyArtistChart `json:"weeklyartistchart"`
}

type WeeklyArtistChart struct {
	Artist []WeeklyArtist        `json:"artist"`
	Attr   WeeklyArtistChartAttr `json:"@attr"`
}

type WeeklyArtistChartAttr struct {
	From string `json:"from"`
	User string `json:"user"`
	To   string `json:"to"`
}

type WeeklyArtist struct {
	Mbid      string     `json:"mbid"`
	Url       string     `json:"url"`
	Name      string     `json:"name"`
	Attr      ArtistAttr `json:"@attr"`
	PlayCount string     `json:"playcount"`
}

type ArtistAttr struct {
	Rank string `json:"rank"`
}

func (c *Client) WeeklyArtistChart(ctx context.Context, page int) (*WeeklyArtistChartData, error) {
	url := buildURL(c.baseURL, "user.getweeklyartistchart", c.userName, c.apiKey, page)
	var result WeeklyArtistChartData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
