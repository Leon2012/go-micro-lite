package broker

import (
	"context"

	"github.com/Leon2012/go-micro-lite/registry"
)

type Options struct {
	Addrs   []string
	Context context.Context
}

type PublishOptions struct {
	Context context.Context
}

type SubscribeOptions struct {
	AutoAck bool
	Queue   string
	Context context.Context
}

type Option func(*Options)

type PublishOption func(*PublishOptions)

type SubscribeOption func(*SubscribeOptions)

var (
	registryKey = "github.com/Leon2012/go-micro-lite/registry"
)

func NewSubscribeOptions(opts ...SubscribeOption) SubscribeOptions {
	opt := SubscribeOptions{
		AutoAck: true,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

func DisableAutoAck() SubscribeOption {
	return func(o *SubscribeOptions) {
		o.AutoAck = false
	}
}

func Queue(name string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Queue = name
	}
}

func Registry(r registry.Registry) Option {
	return func(o *Options) {
		o.Context = context.WithValue(o.Context, registryKey, r)
	}
}

func SubscribeContext(ctx context.Context) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Context = ctx
	}
}
