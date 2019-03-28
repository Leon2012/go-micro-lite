package server

import (
	"context"
)

// HandlerFunc represents a single method of a handler. It's used primarily
// for the wrappers. What's handed to the actual method is the concrete
// request and response types.
type HandlerFunc func(ctx context.Context, req Request, rsp interface{}) error

// HandlerWrapper wraps the HandlerFunc and returns the equivalent
type HandlerWrapper func(HandlerFunc) HandlerFunc

// StreamWrapper wraps a Stream interface and returns the equivalent.
// Because streams exist for the lifetime of a method invocation this
// is a convenient way to wrap a Stream as its in use for trace, monitoring,
// metrics, etc.
type StreamWrapper func(Stream) Stream
