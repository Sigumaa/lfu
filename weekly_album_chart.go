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

func (c *Client) WeeklyAlbumChart(ctx context.Context) (*WeeklyAlbumChartData, error) {
	url := buildURL(c.baseURL, "user.getweeklyalbumchart", c.userName, c.apiKey)
	var result WeeklyAlbumChartData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
