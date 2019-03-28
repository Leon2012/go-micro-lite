package grpc

import (
	"context"
	"testing"

	pb "github.com/Leon2012/go-micro-lite/examples/greeter/proto"
	"github.com/Leon2012/go-micro-lite/registry/consul"
	"github.com/Leon2012/go-micro-lite/server"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type sayServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *sayServer) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func TestGRPCServer(t *testing.T) {
	r := consul.NewRegistry()
	s := NewServer(
		server.Name("foo1"),
		server.Registry(r),
	)

	pb.RegisterSayHandler(s, &sayServer{})

	if err := s.Start(); err != nil {
		t.Fatalf("failed to start: %v", err)
	}

	// check registration
	services, err := r.GetService("foo")
	if err != nil || len(services) == 0 {
		t.Fatalf("failed to get service: %v # %d", err, len(services))
	}

	defer func() {
		if err := s.Stop(); err != nil {
			t.Fatalf("failed to stop: %v", err)
		}
	}()
	t.Logf("address : %s", s.Options().Address)
	cc, err := grpc.Dial(s.Options().Address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	t.Logf("state : %d", cc.GetState())

	testMethods := []string{"/helloworld.Say/Hello", "/greeter.helloworld.Say/Hello"}

	for _, method := range testMethods {
		rsp := pb.Response{}

		if err := cc.Invoke(context.Background(), method, &pb.Request{Name: "John"}, &rsp); err != nil {
			t.Fatalf("error calling server: %v", err)
		}

		if rsp.Msg != "Hello John" {
			t.Fatalf("Got unexpected response %v", rsp.Msg)
		} else {
			t.Logf("rsp : %s", rsp.Msg)
		}
	}
	cc.Close()
}
