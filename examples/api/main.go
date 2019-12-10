package main

import (
	micro "github.com/Leon2012/go-micro-lite"
	"github.com/Leon2012/go-micro-lite/api/resolver"
	"github.com/Leon2012/go-micro-lite/api/resolver/grpc"
	"github.com/Leon2012/go-micro-lite/libs/go-log"
	"github.com/gorilla/mux"
	"net/http"
	regRouter "github.com/Leon2012/go-micro-lite/api/router/registry"
	"github.com/Leon2012/go-micro-lite/api/router"
	httpapi "github.com/Leon2012/go-micro-lite/api/server/http"
	"github.com/Leon2012/go-micro-lite/api/server"
)

var (
	Name                  = "go.micro.api"
	Address               = ":8080"
	Handler               = "meta"
	Resolver              = "micro"
	RPCPath               = "/rpc"
	APIPath               = "/"
	ProxyPath             = "/{service:[a-zA-Z0-9]+}"
	Namespace             = "go.micro.api"
	HeaderPrefix          = "X-Micro-"
	EnableRPC             = false
	ACMEProvider          = "autocert"
	ACMEChallengeProvider = "cloudflare"
)

func main() {
	var opts []server.Option

	// create the router
	var h http.Handler
	r := mux.NewRouter()
	h = r

	// initialise service
	service := micro.NewService()
	// resolver options
	ropts := []resolver.Option{
		resolver.WithNamespace(Namespace),
		resolver.WithHandler(Handler),
	}
	rr := grpc.NewResolver(ropts...)
	log.Logf("Registering API Request Handler at %s", APIPath)
	log.Logf("Registering API Default Handler at %s", APIPath)
	rt := regRouter.NewRouter(
		router.WithNamespace(Namespace),
		router.WithResolver(rr),
		router.WithRegistry(service.Options().Registry),
	)
	r.PathPrefix(APIPath).Handler(handler.Meta(service, rt))

	// create the server
	api := httpapi.NewServer(Address)
	api.Init(opts...)
	api.Handle("/", h)

	// Start API
	if err := api.Start(); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	// Stop API
	if err := api.Stop(); err != nil {
		log.Fatal(err)
	}
}
