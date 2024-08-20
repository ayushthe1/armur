package main

import (
	"context"
	"log"

	"github.com/ayushthe1/armur/contacts-service/contacts/proto"
	leads "github.com/ayushthe1/armur/leads-service/leads/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedContactsServiceServer
	contacts []*proto.AddContactRequest
}

func (s *server) AddContact(ctx context.Context, req *proto.AddContactRequest) (*proto.AddContactResponse, error) {
	log.Printf("Received contact: %v", req)
	s.contacts = append(s.contacts, req)
	return &proto.AddContactResponse{Message: "Contact added successfully"}, nil
}

func (s *server) UpdateContactStatus(ctx context.Context, req *proto.UpdateContactStatusRequest) (*proto.UpdateContactStatusResponse, error) {
	for _, contact := range s.contacts {
		if contact.Email == req.Email {
			contact.Status = req.NewStatus
			log.Printf("Updated contact status: %v", contact)

			// If status is "contacted", notify Leads service
			if req.NewStatus == "contacted" {
				err := s.notifyLeadsService(contact)
				if err != nil {
					return &proto.UpdateContactStatusResponse{Message: "Failed to update lead status"}, err
				}
			}
			return &proto.UpdateContactStatusResponse{Message: "Contact status updated successfully"}, nil
		}
	}
	return &proto.UpdateContactStatusResponse{Message: "Contact not found"}, nil
}

func (s *server) notifyLeadsService(contact *proto.AddContactRequest) error {
	// Connect to Leads service
	conn, err := grpc.Dial("leads-service:50052", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	leadsClient := leads.NewLeadsServiceClient(conn)
	_, err = leadsClient.AddLead(context.Background(), &leads.AddLeadRequest{
		Name:   contact.Name,
		Email:  contact.Email,
		Phone:  contact.Phone,
		Status: "contacted",
	})

	return err
}
