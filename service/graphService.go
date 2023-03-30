package service

import (
	"github.com/graduation-fci/service-graph/dependencies"
	"github.com/graduation-fci/service-graph/domain"
	"github.com/graduation-fci/service-graph/repository"
)

type GraphService struct {
	dp             *dependencies.DP
	drugRepository *repository.DrugRespository
}

func NewGraphService(dp *dependencies.DP) *GraphService {
	return &GraphService{
		dp:             dp,
		drugRepository: repository.NewDrugRepository(dp),
	}
}

func (g *GraphService) InteractionsMap(drugSet []domain.Drug) (map[domain.Hash]domain.InteractionMetadata, error) {
	interactions, err := g.drugRepository.Interaction(drugSet)
	if err != nil {
		return nil, err
	}

	interactionsMap := domain.InteractionsMap(interactions)
	return interactionsMap, nil
}
