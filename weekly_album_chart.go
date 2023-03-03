package lfu

import "context"

type WeeklyAlbumChartData struct {
	WeeklyAlbumChart WeeklyAlbumChart `json:"weeklyalbumchart"`
}

type WeeklyAlbumChart struct {
	Album []WeeklyAlbum        `json:"album"`
	Attr  WeeklyAlbumChartAttr `json:"@attr"`
}

type WeeklyAlbumChartAttr struct {
	From string `json:"from"`
	User string `json:"user"`
	To   string `json:"to"`
}

type WeeklyAlbum struct {
	Artist    AlbumArtist `json:"artist"`
	Mbid      string      `json:"mbid"`
	URL       string      `json:"url"`
	Name      string      `json:"name"`
	Attr      AlbumAttr   `json:"@attr"`
	PlayCount string      `json:"playcount"`
}

type AlbumAttr struct {
	Rank string `json:"rank"`
}

type AlbumArtist struct {
	Mbid string `json:"mbid"`
	Text string `json:"#text"`
}

func (c *Client) WeeklyAlbumChart(ctx context.Context, page int) (*WeeklyAlbumChartData, error) {
	url := buildURL(c.baseURL, "user.getweeklyalbumchart", c.userName, c.apiKey, page)
	var result WeeklyAlbumChartData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
