// Package rpc provides an rpc server
package rpc

import (
	"github.com/Leon2012/go-micro-lite/server"
	"github.com/Leon2012/go-micro-lite/server/grpc"
)

// NewServer returns a micro server interface
func NewServer(opts ...server.Option) server.Server {
	return grpc.NewServer(opts...)
}
