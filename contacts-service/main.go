package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ayushthe1/armur/contacts-service/contacts/proto"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterContactsServiceServer(grpcServer, &server{})

	fmt.Println("Contacts gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
