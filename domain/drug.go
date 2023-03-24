package domain

import "github.com/graduation-fci/service-graph/proto"

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
