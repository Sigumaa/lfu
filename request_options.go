package lfu

import (
	"net/url"
	"strconv"
)

type RequestOption func(*requestOptions)

type requestOptions struct {
	urlParams url.Values
}

type prange string

const (
	// PeriodOverall is the default period
	PeriodOverall prange = "overall"

	// Period7Day is the last 7 days
	Period7Day prange = "7day"

	// Period1Month is the last 1 month
	Period1Month prange = "1month"

	// Period3Month is the last 3 months
	Period3Month prange = "3month"

	// Period6Month is the last 6 months
	Period6Month prange = "6month"

	// Period12Month is the last 12 months
	Period12Month prange = "12month"
)

// Period sets the value of the `period` parameter in the request URL to the given value.
func Period(period prange) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("period", string(period))
	}
}

// Page sets the value of the `page` parameter in the request URL to the given value.
func Page(page int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("page", strconv.Itoa(page))
	}
}

// Limit sets the value of the `limit` parameter in the request URL to the given value.
func Limit(limit int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("limit", strconv.Itoa(limit))
	}
}

// From sets the value of the `from` parameter in the request URL to the given value.
func From(from int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("from", strconv.Itoa(from))
	}
}

// To sets the value of the `to` parameter in the request URL to the given value.
func To(to int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("to", strconv.Itoa(to))
	}
}

// Extended sets the value of the `extended` parameter in the request URL to "1" if the given
// boolean is true, and "0" otherwise.
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
