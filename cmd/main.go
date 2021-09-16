package main

import (
	"flag"
	"fmt"
	"github.com/Nikita99100/Fibonacci/handler"
	"github.com/Nikita99100/Fibonacci/server"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"log"
)

func main() {
	// Parse flags
	restPort := flag.String("server", "80", "REST port")
	grpcPort := flag.String("rpc", "8080", "gRPC port")
	memcacheAddress := flag.String("cache", "127.0.0.1:11211", "memcache address")
	flag.Parse()
	println("server port:", *restPort)
	println("grpc port:", *grpcPort)

	//Create memcache client
	mc := memcache.New(*memcacheAddress)
	if err := mc.Ping(); err != nil {
		log.Fatal(errors.Wrap(err, "Failed to connect to memcache server"))
	}

	// Create a handler object
	hdlr := handler.Handler{
		Cache: mc,
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
