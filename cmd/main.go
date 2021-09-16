package main

import (
	"flag"
	"fmt"
	"github.com/Nikita99100/Fibonacci/handler"
	"github.com/Nikita99100/Fibonacci/pkg/api"
	"github.com/Nikita99100/Fibonacci/server"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Parse flags
	restPort := flag.String("server", "80", "REST port")
	grpcPort := flag.String("rpc", "8080", "gRPC port")
	memcacheAddress := flag.String("cache", "127.0.0.1:11211", "memcache address")
	flag.Parse()
	println("server port:", *restPort)
	println("api port:", *grpcPort)

	//Create memcache client
	mc := memcache.New(*memcacheAddress)
	if err := mc.Ping(); err != nil {
		log.Fatal(errors.Wrap(err, "Failed to connect to memcache server"))
	}

	// Create a handler object
	hdlr := handler.Handler{
		Cache: mc,
	}

	grpcS := grpc.NewServer()
	grpcServ := &server.GRPCServer{Handler: &hdlr}
	api.RegisterFibonacciServer(grpcS, grpcServ)
	grpcLis, err := net.Listen("tcp", fmt.Sprintf(":%v", *grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", grpcLis.Addr())
	if err := grpcS.Serve(grpcLis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Create a rest server gateway and handle http requests
	router := echo.New()
	rest := server.Rest{
		Router:  router,
		Handler: &hdlr,
	}
	rest.Route()

	//start rest server
	if err := router.Start(fmt.Sprintf(":%v", *restPort)); err != nil {
		log.Fatal(errors.Wrap(err, "Failed to start rest server"))
	}
}
