package domain

import (
	"sort"
	"strings"

	"github.com/graduation-fci/service-graph/proto"
	neo4jDriver "github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type (
	Drug = string
	Hash = string
)

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
	severity, _ := options.Record().Get("severity")
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

// TODO:
// Insert hash preprocessed in the graph (pipeline tool); to avoid runtime hashing.
func (i Interaction) HashKey() Hash {
	internalHash := i.Metadata.Hash

	if internalHash == EmptyString {
		return EmptyString
	}

	internalHash = strings.ToLower(internalHash)
	splitHash := strings.Split(internalHash, SpaceDelimiter)

	sort.Strings(splitHash)
	return strings.Join(splitHash, EmptyString)
}

func ToInternalHash(concatinationHash Hash) Hash {
	if concatinationHash == EmptyString {
		return EmptyString
	}

	concatinationHash = strings.ToLower(concatinationHash)
	splitHash := strings.Split(concatinationHash, SpaceDelimiter)

	sort.Strings(splitHash)
	return strings.Join(splitHash, EmptyString)
}

func InteractionsMap(interactions []Interaction) map[Hash]InteractionMetadata {
	hashMap := make(map[Hash]InteractionMetadata)
	for _, interaction := range interactions {
		key := interaction.HashKey()
		hashMap[key] = interaction.Metadata
	}
	return hashMap
}
