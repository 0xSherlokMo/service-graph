package server

import (
	"context"
	"os"

	"github.com/graduation-fci/service-graph/dependencies"
	"github.com/graduation-fci/service-graph/domain"
	"github.com/graduation-fci/service-graph/proto"
	"github.com/graduation-fci/service-graph/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DefaultGraphServer struct {
	proto.UnimplementedGraphServiceServer
	dp           *dependencies.DP
	graphService *service.GraphService
}

func NewGraphServer(dp *dependencies.DP) *DefaultGraphServer {
	return &DefaultGraphServer{
		dp:           dp,
		graphService: service.NewGraphService(dp),
	}
}

func (s *DefaultGraphServer) CheckInteractions(ctx context.Context, request *proto.CheckInteractionsRequest) (*proto.CheckInteractionsResponse, error) {
	if len(request.GetMedecines()) < 2 {
		return nil, status.Error(codes.InvalidArgument, "invalid_data")
	}

	set := domain.DrugSet(request.GetMedecines())
	knowledge, err := s.graphService.InteractionsMap(set)
	if err != nil {
		return nil, err
	}

	permutations := s.graphService.MedecinePermutation(request.GetMedecines(), knowledge)
	response := &proto.CheckInteractionsResponse{Permutations: permutations}

	if os.Getenv("DISABLE_NOTIFICATION") == "TRUE" {
		return response, nil
	}

	response.Notification = s.graphService.GetNotification(permutations, request.MedicationId)

	go s.graphService.SaveReport(permutations, request.MedicationId)

	return response, nil
}
