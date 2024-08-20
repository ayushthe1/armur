package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ayushthe1/armur/leads-service/leads/proto"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterLeadsServiceServer(grpcServer, &server{})

	fmt.Println("Leads gRPC server is running on port 50052...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
