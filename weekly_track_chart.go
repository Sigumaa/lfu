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

// Len returns the number of tracks in the chart.
func (w *WeeklyTrackChartData) Len() int {
	return len(w.WeeklyTrackChart.Track)
}

// User returns the name of the user
func (w *WeeklyTrackChartData) User() string {
	return w.WeeklyTrackChart.Attr.User
}

// From returns the start date of the chart.
func (w *WeeklyTrackChartData) From() string {
	return w.WeeklyTrackChart.Attr.From
}

// To returns the end date of the chart.
func (w *WeeklyTrackChartData) To() string {
	return w.WeeklyTrackChart.Attr.To
}

// ArtistName returns the name of the artist for the track at the given index.
// The index must be range [0, w.Len()).
func (w *WeeklyTrackChartData) ArtistName(index int) (string, error) {
	if index < 0 || index >= w.Len() {
		return "", ErrOutOfBounds
	}
	return w.WeeklyTrackChart.Track[index].Artist.Text, nil
}

// ArtistMbid returns the Mbid of the artist for the track at the given index.
// The index must be range [0, w.Len()).
func (w *WeeklyTrackChartData) ArtistMbid(index int) (string, error) {
	if index < 0 || index >= w.Len() {
		return "", ErrOutOfBounds
	}
	return w.WeeklyTrackChart.Track[index].Artist.Mbid, nil
}

// Image returns the image for the track at the given index.
// The index must be range [0, w.Len()).
func (w *WeeklyTrackChartData) Image(index int) ([]Image, error) {
	if index < 0 || index >= w.Len() {
		return nil, ErrOutOfBounds
	}
	return w.WeeklyTrackChart.Track[index].Image, nil
}

// ImageURL returns the image URL for the track at the given index.
// The index must be range [0, w.Len()).
func (w *WeeklyTrackChartData) ImageURL(index int, size imageSize) (string, error) {
	if index < 0 || index >= w.Len() {
		return "", ErrOutOfBounds
	}
	return w.WeeklyTrackChart.Track[index].Image[size].Text, nil
}

// Mbid returns the Mbid of the track at the given index.
// The index must be range [0, w.Len()).
func (w *WeeklyTrackChartData) Mbid(index int) (string, error) {
	if index < 0 || index >= w.Len() {
		return "", ErrOutOfBounds
	}
	return w.WeeklyTrackChart.Track[index].Mbid, nil
}

// URL returns the URL of the track at the given index.
// The index must be range [0, w.Len()).
func (w *WeeklyTrackChartData) URL(index int) (string, error) {
	if index < 0 || index >= w.Len() {
		return "", ErrOutOfBounds
	}
	return w.WeeklyTrackChart.Track[index].URL, nil
}

// Name returns the name of the track at the given index.
// The index must be range [0, w.Len()).
func (w *WeeklyTrackChartData) Name(index int) (string, error) {
	if index < 0 || index >= w.Len() {
		return "", ErrOutOfBounds
	}
	return w.WeeklyTrackChart.Track[index].Name, nil
}

// Rank returns the rank of the track at the given index.
// The index must be range [0, w.Len()).
func (w *WeeklyTrackChartData) Rank(index int) (string, error) {
	if index < 0 || index >= w.Len() {
		return "", ErrOutOfBounds
	}
	return w.WeeklyTrackChart.Track[index].Attr.Rank, nil
}

// PlayCount returns the play count of the track at the given index.
// The index must be range [0, w.Len()).
func (w *WeeklyTrackChartData) PlayCount(index int) (string, error) {
	if index < 0 || index >= w.Len() {
		return "", ErrOutOfBounds
	}
	return w.WeeklyTrackChart.Track[index].PlayCount, nil
}
