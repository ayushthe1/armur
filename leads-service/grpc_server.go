package main

import (
	"context"
	"log"

	"github.com/ayushthe1/armur/leads-service/leads/proto"
)

type server struct {
	proto.UnimplementedLeadsServiceServer
	leads []*proto.AddLeadRequest
}

func (s *server) AddLead(ctx context.Context, req *proto.AddLeadRequest) (*proto.AddLeadResponse, error) {
	// Log the lead and add it to the in-memory list
	log.Printf("Received lead: %v", req)
	s.leads = append(s.leads, req)
	return &proto.AddLeadResponse{Message: "Lead added successfully"}, nil
}

func (s *server) UpdateLeadStatus(ctx context.Context, req *proto.UpdateLeadStatusRequest) (*proto.UpdateLeadStatusResponse, error) {
	for _, lead := range s.leads {
		if lead.Email == req.Email {
			lead.Status = req.NewStatus
			log.Printf("Updated lead status: %v", lead)
			return &proto.UpdateLeadStatusResponse{Message: "Lead status updated successfully"}, nil
		}
	}
	return &proto.UpdateLeadStatusResponse{Message: "Lead not found"}, nil
}
