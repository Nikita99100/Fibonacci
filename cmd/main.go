package main

import (
	"flag"
	"github.com/Nikita99100/Fibonacci/config"
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
	defaultConfigPath         = "config/configs.toml"
	restServerShutdownTimeout = 30 * time.Second
)

func main() {
	// Parse flags
	configPath := flag.String("config", defaultConfigPath, "configuration file path")
	flag.Parse()

	cfg, err := config.Parse(*configPath)
	if err != nil {
		log.Fatalf("failed to parse the config file: %v", err)
	}
	restPort := cfg.RestPort
	grpcPort := cfg.GrpcPort
	memcacheAddress := cfg.MemcacheAddress
	flag.Parse()

	//Create memcache client
	mc := memcache.New(memcacheAddress)
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
	go web.Start(router, restPort)
	defer web.Stop(router, restServerShutdownTimeout)

	// Start a grpc server
	grpcS := grpc.NewServer()
	grpcServ := &server.GRPCServer{Handler: &hdlr}
	go api.Start(grpcS, grpcServ, grpcPort)
	defer grpcS.Stop()

	// Wait for program exit
	<-os.NotifyAboutExit()
}
