package http

import (
	"net/http"

	micro "github.com/Leon2012/go-micro-lite"
	"github.com/Leon2012/go-micro-lite/selector"
)

func NewRoundTripper(opts ...Option) http.RoundTripper {
	options := Options{
		Registry: micro.DefaultRegistry,
	}
	for _, o := range opts {
		o(&options)
	}

	return &roundTripper{
		rt:   http.DefaultTransport,
		st:   selector.Random,
		opts: options,
	}
}
