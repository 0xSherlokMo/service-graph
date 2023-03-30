package domain

import (
	"github.com/graduation-fci/service-graph/proto"
	neo4jDriver "github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Drug = string

func DrugSet(medecines []*proto.Medecine) []Drug {
	drugsMap := make(map[Drug]Empty)
	for _, medecine := range medecines {
		for _, drug := range medecine.Drugs {
			drugsMap[drug] = Empty{}
		}
	}

	var drugSet []Drug
	for drug := range drugsMap {
		drugSet = append(drugSet, drug)
	}

	return drugSet
}

type Interaction struct {
	Node1    Drug
	Node2    Drug
	Metadata InteractionMetadata
}

type InteractionMetadata struct {
	Hash               string
	Serverity          string
	ProfessionalEffect string
	ConsumerEffect     string
}

func NewInteraction(options neo4jDriver.ResultWithContext) Interaction {
	node1, _ := options.Record().Get("node1.name")
	node2, _ := options.Record().Get("node2.name")
	consumerEffect, _ := options.Record().Get("consumerEffect")
	professionalEffect, _ := options.Record().Get("professionalEffect")
	severity, _ := options.Record().Get("consumerEffect")
	hash, _ := options.Record().Get("hash")
	return Interaction{
		Node1: node1.(Drug),
		Node2: node2.(Drug),
		Metadata: InteractionMetadata{
			Hash:               hash.(string),
			Serverity:          severity.(string),
			ConsumerEffect:     consumerEffect.(string),
			ProfessionalEffect: professionalEffect.(string),
		},
	}
}
