package server

import (
	"context"
	"errors"

	"github.com/graduation-fci/service-graph/dependencies"
	"github.com/graduation-fci/service-graph/domain"
	"github.com/graduation-fci/service-graph/proto"
	"github.com/graduation-fci/service-graph/service"
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
		return nil, errors.New("invalid request")
	}

	set := domain.DrugSet(request.GetMedecines())
	knowledge, err := s.graphService.InteractionsMap(set)
	if err != nil {
		return nil, err
	}

	permutations := s.graphService.MedecinePermutation(request.GetMedecines(), knowledge)

	return &proto.CheckInteractionsResponse{
		Permutations: permutations,
	}, nil
}
