package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ayushthe1/armur/contacts-service/contacts/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedContactsServiceServer
	contacts []*proto.AddContactRequest
}

func (s *server) AddContact(ctx context.Context, req *proto.AddContactRequest) (*proto.AddContactResponse, error) {
	// For now, we'll just log the contact and add it to an in-memory list
	log.Printf("Received contact: %v", req)
	s.contacts = append(s.contacts, req)

	log.Println("All contacts are : ", s.contacts)
	return &proto.AddContactResponse{Message: "Contact added successfully"}, nil
}

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
