package main

import (
	"context"
	"log"

	"github.com/ayushthe1/armur/clients-service/clients/proto"
)

type server struct {
	proto.UnimplementedClientsServiceServer
	clients []*proto.Client
}

func (s *server) AddClient(ctx context.Context, req *proto.AddClientRequest) (*proto.AddClientResponse, error) {
	// Log the client and add it to the in-memory list
	log.Printf("Received client: %v", req)
	client := &proto.Client{
		Name:   req.Name,
		Email:  req.Email,
		Phone:  req.Phone,
		Status: req.Status,
	}
	s.clients = append(s.clients, client)
	return &proto.AddClientResponse{Message: "Client added successfully"}, nil
}

func (s *server) GetClients(ctx context.Context, req *proto.GetClientsRequest) (*proto.GetClientsResponse, error) {
	return &proto.GetClientsResponse{Clients: s.clients}, nil
}
