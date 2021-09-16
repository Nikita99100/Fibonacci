package api

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Start(grpcS *grpc.Server, grpcServ FibonacciServer, addr string) {
	RegisterFibonacciServer(grpcS, grpcServ)
	grpcLis, err := net.Listen("tcp", fmt.Sprintf(":%v", addr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", grpcLis.Addr())
	if err := grpcS.Serve(grpcLis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
