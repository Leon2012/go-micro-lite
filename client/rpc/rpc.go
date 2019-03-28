// Package rpc provides an rpc client
package rpc

import (
	"github.com/Leon2012/go-micro-lite/client"
	"github.com/Leon2012/go-micro-lite/client/grpc"
)

// NewClient returns a new micro client interface
func NewClient(opts ...client.Option) client.Client {
	return grpc.NewClient(opts...)
}
