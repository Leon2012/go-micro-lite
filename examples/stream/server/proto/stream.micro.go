// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/examples/stream/server/proto/stream.proto

/*
Package stream is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/examples/stream/server/proto/stream.proto

It has these top-level messages:
	Request
	Response
*/
package stream

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"

	client "github.com/Leon2012/go-micro-lite/client"
	server "github.com/Leon2012/go-micro-lite/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Streamer service

type StreamerService interface {
	Stream(ctx context.Context, opts ...client.CallOption) (Streamer_StreamService, error)
	ServerStream(ctx context.Context, in *Request, opts ...client.CallOption) (Streamer_ServerStreamService, error)
}

type streamerService struct {
	c           client.Client
	serviceName string
}

func NewStreamerService(serviceName string, c client.Client) StreamerService {
	if len(serviceName) == 0 {
		serviceName = "streamer"
	}
	return &streamerService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *streamerService) Stream(ctx context.Context, opts ...client.CallOption) (Streamer_StreamService, error) {
	req := c.c.NewRequest(c.serviceName, "Streamer.Stream", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &streamerStreamService{stream}, nil
}

type Streamer_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Request) error
	Recv() (*Response, error)
}

type streamerStreamService struct {
	stream client.Stream
}

func (x *streamerStreamService) Close() error {
	return x.stream.Close()
}

func (x *streamerStreamService) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerStreamService) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerStreamService) Send(m *Request) error {
	return x.stream.Send(m)
}

func (x *streamerStreamService) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamerService) ServerStream(ctx context.Context, in *Request, opts ...client.CallOption) (Streamer_ServerStreamService, error) {
	req := c.c.NewRequest(c.serviceName, "Streamer.ServerStream", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &streamerServerStreamService{stream}, nil
}

type Streamer_ServerStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Response, error)
}

type streamerServerStreamService struct {
	stream client.Stream
}

func (x *streamerServerStreamService) Close() error {
	return x.stream.Close()
}

func (x *streamerServerStreamService) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerServerStreamService) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerServerStreamService) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Streamer service

type StreamerHandler interface {
	Stream(context.Context, Streamer_StreamStream) error
	ServerStream(context.Context, *Request, Streamer_ServerStreamStream) error
}

func RegisterStreamerHandler(s server.Server, hdlr StreamerHandler, opts ...server.HandlerOption) {
	type streamer interface {
		Stream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
	}
	type Streamer struct {
		streamer
	}
	h := &streamerHandler{hdlr}
	s.Handle(s.NewHandler(&Streamer{h}, opts...))
}

type streamerHandler struct {
	StreamerHandler
}

func (h *streamerHandler) Stream(ctx context.Context, stream server.Stream) error {
	return h.StreamerHandler.Stream(ctx, &streamerStreamStream{stream})
}

type Streamer_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
	Recv() (*Request, error)
}

type streamerStreamStream struct {
	stream server.Stream
}

func (x *streamerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *streamerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerStreamStream) Send(m *Response) error {
	return x.stream.Send(m)
}

func (x *streamerStreamStream) Recv() (*Request, error) {
	m := new(Request)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *streamerHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(Request)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.StreamerHandler.ServerStream(ctx, m, &streamerServerStreamStream{stream})
}

type Streamer_ServerStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
}

type streamerServerStreamStream struct {
	stream server.Stream
}

func (x *streamerServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *streamerServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamerServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamerServerStreamStream) Send(m *Response) error {
	return x.stream.Send(m)
}
