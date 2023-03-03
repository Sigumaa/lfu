package lfu

import (
	"net/url"
	"strconv"
)

type RequestOption func(*requestOptions)

type requestOptions struct {
	urlParams url.Values
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
