package lfu

import "context"

type WeeklyTrackChartData struct {
	WeeklyTrackChart WeeklyTrackChart `json:"weeklytrackchart"`
}

type WeeklyTrackChart struct {
	Track []WeeklyTrack        `json:"track"`
	Attr  WeeklyTrackChartAttr `json:"@attr"`
}

type WeeklyTrackChartAttr struct {
	From string `json:"from"`
	User string `json:"user"`
	To   string `json:"to"`
}

type WeeklyTrack struct {
	Artist    TrackArtist `json:"artist"`
	Image     []Image     `json:"image"`
	Mbid      string      `json:"mbid"`
	URL       string      `json:"url"`
	Name      string      `json:"name"`
	Attr      TrackAttr   `json:"@attr"`
	PlayCount string      `json:"playcount"`
}

type TrackAttr struct {
	Rank string `json:"rank"`
}

type TrackArtist struct {
	Mbid string `json:"mbid"`
	Text string `json:"#text"`
}

func (c *Client) WeeklyTrackChart(ctx context.Context, page int) (*WeeklyTrackChartData, error) {
	url := buildURL(c.baseURL, "user.getweeklytrackchart", c.userName, c.apiKey, page)
	var result WeeklyTrackChartData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
