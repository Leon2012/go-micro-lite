package grpc

import (
	"github.com/Leon2012/go-micro-lite/server"
	"github.com/Leon2012/go-micro-lite/server/debug"
)

// We use this to wrap any debug handlers so we preserve the signature Debug.{Method}
type Debug struct {
	debug.DebugHandler
}

func registerDebugHandler(s server.Server) {
	s.Handle(s.NewHandler(&Debug{s.Options().DebugHandler}, server.InternalHandler(true)))
}
