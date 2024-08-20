package main

import (
	"context"
	"log"

	clients "github.com/ayushthe1/armur/clients-service/clients/proto" // Import the generated Clients service code
	"github.com/ayushthe1/armur/leads-service/leads/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedLeadsServiceServer
	leads []*proto.AddLeadRequest
}

func (s *server) AddLead(ctx context.Context, req *proto.AddLeadRequest) (*proto.AddLeadResponse, error) {
	log.Printf("Received lead: %v", req)
	s.leads = append(s.leads, req)
	return &proto.AddLeadResponse{Message: "Lead added successfully"}, nil
}

func (s *server) UpdateLeadStatus(ctx context.Context, req *proto.UpdateLeadStatusRequest) (*proto.UpdateLeadStatusResponse, error) {
	for _, lead := range s.leads {
		if lead.Email == req.Email {
			lead.Status = req.NewStatus
			log.Printf("Updated lead status: %v", lead)

			// If status is "payment done", notify Clients service
			if req.NewStatus == "payment done" {
				err := s.notifyClientsService(lead)
				if err != nil {
					return &proto.UpdateLeadStatusResponse{Message: "Failed to update client status"}, err
				}
			}
			return &proto.UpdateLeadStatusResponse{Message: "Lead status updated successfully"}, nil
		}
	}
	return &proto.UpdateLeadStatusResponse{Message: "Lead not found"}, nil
}

func (s *server) notifyClientsService(lead *proto.AddLeadRequest) error {
	// Connect to Clients service
	conn, err := grpc.Dial("clients-service:50053", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	clientsClient := clients.NewClientsServiceClient(conn)
	_, err = clientsClient.AddClient(context.Background(), &clients.AddClientRequest{
		Name:   lead.Name,
		Email:  lead.Email,
		Phone:  lead.Phone,
		Status: "payment done",
	})

	return err
}
