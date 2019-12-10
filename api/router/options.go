package router

import (
	mmicro "github.com/Leon2012/go-micro-lite"
	"github.com/Leon2012/go-micro-lite/api/resolver"
	"github.com/Leon2012/go-micro-lite/registry"
	"github.com/micro/go-micro/api/resolver/micro"
)

type Options struct {
	Namespace string
	Handler   string
	Registry  registry.Registry
	Resolver  resolver.Resolver
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Handler:  "meta",
		Registry: mmicro.DefaultRegistry,
	}

	for _, o := range opts {
		o(&options)
	}

	if options.Resolver == nil {
		options.Resolver = micro.NewResolver(
			resolver.WithHandler(options.Handler),
			resolver.WithNamespace(options.Namespace),
		)
	}

	return options
}

func WithHandler(h string) Option {
	return func(o *Options) {
		o.Handler = h
	}
}

func WithNamespace(ns string) Option {
	return func(o *Options) {
		o.Namespace = ns
	}
}

func WithRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

func WithResolver(r resolver.Resolver) Option {
	return func(o *Options) {
		o.Resolver = r
	}
}
