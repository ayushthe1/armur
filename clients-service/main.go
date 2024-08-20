package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ayushthe1/armur/clients-service/clients/proto"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen on port 50053: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterClientsServiceServer(grpcServer, &server{})

	fmt.Println("Clients gRPC server is running on port 50053...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
