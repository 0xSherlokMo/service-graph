package domain

import (
	"strings"

	"github.com/graduation-fci/service-graph/proto"
)

func MedecineDrugInteractions(firstMedecine, secondMedecine *proto.Medecine, knowledgeMap map[Hash]InteractionMetadata) *proto.Permutation {
	permutation := PrimitivePermutationObject(firstMedecine, secondMedecine)

	memo := make(map[Hash]Empty)
	for _, drug1 := range firstMedecine.GetDrugs() {
		for _, drug2 := range secondMedecine.GetDrugs() {
			memoHash := drug1 + SpaceDelimiter + drug2
			memoInverseConcatinationHash := drug2 + SpaceDelimiter + drug1
			_, indexExists := memo[memoHash]
			_, inverseIndexExists := memo[memoInverseConcatinationHash]
			if indexExists || inverseIndexExists || Assert(drug1, drug2) {
				continue
			}
			memo[memoHash] = Empty{}
			memo[memoInverseConcatinationHash] = Empty{}

			suffixToHash := strings.Split(drug1, SpaceDelimiter)[0] + SpaceDelimiter + strings.Split(drug2, SpaceDelimiter)[0]
			internalHash := ToInternalHash(suffixToHash)
			interaction, exists := knowledgeMap[internalHash]
			if !exists {
				continue
			}
			permutation.Interactions = append(permutation.Interactions, &proto.Interaction{
				Drugs:              []string{drug1, drug2},
				ProfessionalEffect: interaction.ProfessionalEffect,
				ConsumerEffect:     interaction.ConsumerEffect,
				Severity:           interaction.Serverity,
			})
		}
	}

	return permutation
}

func PrimitivePermutationObject(firstMedecine, secondMedecine *proto.Medecine) *proto.Permutation {
	return &proto.Permutation{
		Medecines: []*proto.I18N{
			{
				NameEn: firstMedecine.GetName().GetNameEn(),
				NameAr: firstMedecine.GetName().GetNameAr(),
			},
			{
				NameEn: secondMedecine.GetName().GetNameEn(),
				NameAr: secondMedecine.GetName().GetNameAr(),
			},
		},
		Interactions: []*proto.Interaction{},
	}
}
