package lfu

import (
	"net/url"
	"strconv"
)

type RequestOption func(*requestOptions)

type requestOptions struct {
	urlParams url.Values
}

type Range string

const (
	PeriodOverall Range = "overall"
	Period7Day    Range = "7day"
	Period1Month  Range = "1month"
	Period3Month  Range = "3month"
	Period6Month  Range = "6month"
	Period12Month Range = "12month"
)

func Period(period Range) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("period", string(period))
	}
}

func Page(page int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("page", strconv.Itoa(page))
	}
}

func Limit(limit int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("limit", strconv.Itoa(limit))
	}
}

func From(from int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("from", strconv.Itoa(from))
	}
}

func To(to int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("to", strconv.Itoa(to))
	}
}

func Extended(extended bool) RequestOption {
	e := "0"
	if extended {
		e = "1"
	}
	return func(o *requestOptions) {
		o.urlParams.Set("extended", e)
	}
}

func processOptions(opts ...RequestOption) requestOptions {
	o := requestOptions{
		urlParams: url.Values{},
	}
	for _, opt := range opts {
		opt(&o)
	}

	return o
}
