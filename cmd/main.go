package main

import (
	"flag"
	"fmt"
	"github.com/Nikita99100/Fibonacci/handler"
	"github.com/Nikita99100/Fibonacci/server"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"log"
)

func main() {
	// Parse flags
	restPort := flag.String("server", "80", "REST port")
	grpcPort := flag.String("rpc", "8080", "gRPC port")
	flag.Parse()
	println("server port:", *restPort)
	println("grpc port:", *grpcPort)

	// Create a handler object
	hdlr := handler.Handler{}

	// Create a server gateway and handle http requests
	router := echo.New()
	rest := server.Rest{
		Router:  router,
		Handler: &hdlr,
	}
	rest.Route()

	//start server server
	if err := router.Start(fmt.Sprintf(":%v", *restPort)); err != nil {
		log.Fatal(errors.Wrap(err, "Failed to start server server"))
	}
}
