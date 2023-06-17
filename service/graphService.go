package service

import (
	"errors"
	"log"

	"github.com/graduation-fci/service-graph/dependencies"
	"github.com/graduation-fci/service-graph/domain"
	"github.com/graduation-fci/service-graph/proto"
	"github.com/graduation-fci/service-graph/repository"
	"go.mongodb.org/mongo-driver/mongo"
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

func (g *GraphService) MedecinePermutation(medecines []*proto.Medecine, knowledge map[string]domain.InteractionMetadata) []*proto.Permutation {
	var perumations []*proto.Permutation

	// maintaining two index pointers; doing two pinters approach.
	leftPointer, rightPointer := 0, 1
	stopIndex := len(medecines) - 1

	for leftPointer < stopIndex {
		for rightPointer <= stopIndex {
			permutation := domain.MedecineDrugInteractions(
				medecines[leftPointer],
				medecines[rightPointer],
				knowledge,
			)

			if len(permutation.Interactions) > 0 {
				perumations = append(perumations, permutation)
			}

			rightPointer++
		}

		leftPointer++
		rightPointer = leftPointer + 1
	}

	return perumations
}

func (g *GraphService) GetNotification(currentReport []*proto.Permutation, medicationId int64) *proto.Notification {
	if medicationId == 0 {
		return &proto.Notification{}
	}
	report, err := g.drugRepository.LatestReport(medicationId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Printf("[Report] cannot get latest report to build notifications, with error: %s \n", err)
		}

		return &proto.Notification{}
	}

	return report.BuildNotification(currentReport)
}

func (g *GraphService) SaveReport(currentReport []*proto.Permutation, medicationId int64) {
	if medicationId == 0 {
		return
	}

	defer func() {
		err := recover()
		if err != nil {
			log.Printf("[SaveReport] recovered from panic with error %s \n", err)
		}
	}()
	report := domain.Report{}.ToModel(currentReport, medicationId)
	g.drugRepository.SaveReport(report)
}
