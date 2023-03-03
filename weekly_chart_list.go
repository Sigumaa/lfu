package lfu

import "context"

type WeeklyChartListData struct {
	WeeklyChartList WeeklyChartList `json:"weeklychartlist"`
}

type WeeklyChartList struct {
	Chart []WeeklyChart       `json:"chart"`
	Attr  WeeklyChartListAttr `json:"@attr"`
}

type WeeklyChartListAttr struct {
	User string `json:"user"`
}

type WeeklyChart struct {
	Text string `json:"#text"`
	From string `json:"from"`
	To   string `json:"to"`
}

// WeeklyChartList Get a list of available charts for this user, expressed as date ranges which can be sent to the chart services.
func (c *Client) WeeklyChartList(ctx context.Context) (*WeeklyChartListData, error) {
	url := buildURL(c.baseURL, "user.getweeklychartlist", c.userName, c.apiKey)
	var result WeeklyChartListData
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
