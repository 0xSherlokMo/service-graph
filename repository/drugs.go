package repository

import (
	"log"
	"time"

	"github.com/graduation-fci/service-graph/dependencies"
	"github.com/graduation-fci/service-graph/domain"
	"github.com/graduation-fci/service-graph/utilities/neo4j"
	neo4jDriver "github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"golang.org/x/net/context"
)

const tag = "[DrugRepository]"

type DrugRespository struct {
	dp      *dependencies.DP
	timeout time.Duration
}

func NewDrugRepository(dp *dependencies.DP) *DrugRespository {
	return &DrugRespository{
		dp:      dp,
		timeout: 10 * time.Second,
	}
}

func (d *DrugRespository) Interaction(drugSet []domain.Drug) ([]domain.Interaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()

	session, close := neo4j.NewSession(ctx, d.dp.GetNeo4j(), neo4jDriver.AccessModeRead)
	close()
	interactions, err := neo4jDriver.ExecuteRead(ctx, session, func(tx neo4jDriver.ManagedTransaction) ([]domain.Interaction, error) {
		result, err := tx.Run(ctx,
			`MATCH 
			(node1:Drug)-[interaction:INTERACTS]->(node2:Drug)
			WHERE node1.name IN $drugList AND node2.name IN $drugList
			RETURN 
			node1.name, 
			node2.name, 
			interaction.consumerEffect AS consumerEffect,
			interaction.professionalEffect AS professionalEffect,
			interaction.level AS severity, 
			interaction.hashedName AS hash`,
			map[string]any{
				"drugList": drugSet,
			})
		if err != nil {
			log.Println(tag, "error", err)
			return nil, err
		}

		var interactions []domain.Interaction
		for result.Next(ctx) {
			interactions = append(interactions, domain.NewInteraction(result))
		}
		return interactions, nil
	})
	log.Println(interactions)
	return interactions, err
}
