// Package router provides api service routing
package router

import (
	"net/http"

	"github.com/Leon2012/go-micro-lite/api"
)

// Router is used to determine an endpoint for a request
type Router interface {
	// Returns options
	Options() Options
	// Stop the router
	Close() error
	// Endpoint returns an api.Service endpoint or an error if it does not exist
	Endpoint(r *http.Request) (*api.Service, error)
	// Route returns an api.Service route
	Route(r *http.Request) (*api.Service, error)
}
