package main

import (
	"flag"
	"github.com/Nikita99100/Fibonacci/handler"
	"github.com/Nikita99100/Fibonacci/pkg/api"
	"github.com/Nikita99100/Fibonacci/pkg/os"
	"github.com/Nikita99100/Fibonacci/pkg/web"
	"github.com/Nikita99100/Fibonacci/server"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	restServerShutdownTimeout = 30 * time.Second
)

func main() {
	// Parse flags
	restPort := flag.String("server", "80", "REST port")
	grpcPort := flag.String("rpc", "8080", "gRPC port")
	memcacheAddress := flag.String("cache", "127.0.0.1:11211", "memcache address")
	flag.Parse()

	//Create memcache client
	mc := memcache.New(*memcacheAddress)
	if err := mc.Ping(); err != nil {
		log.Fatal(errors.Wrap(err, "Failed to connect to memcache server"))
	}

	// Create a handler object
	hdlr := handler.Handler{Cache: mc}

	// Create a rest server gateway and handle http requests
	router := echo.New()
	rest := server.Rest{
		Router:  router,
		Handler: &hdlr,
	}
	rest.Route()
	// Start an http server and remember to shut it down
	go web.Start(router, *restPort)
	defer web.Stop(router, restServerShutdownTimeout)

	// Start a grpc server
	grpcS := grpc.NewServer()
	grpcServ := &server.GRPCServer{Handler: &hdlr}
	go api.Start(grpcS, grpcServ, *grpcPort)
	defer grpcS.Stop()

	// Wait for program exit
	<-os.NotifyAboutExit()
}
