package neo4j

import (
	"context"

	"github.com/graduation-fci/service-graph/domain"
	neo4jDriver "github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func NewSession(ctx context.Context, driver neo4jDriver.DriverWithContext, AccessMode neo4jDriver.AccessMode) (neo4jDriver.SessionWithContext, domain.CloserFunc) {
	session := driver.NewSession(ctx, neo4jDriver.SessionConfig{
		AccessMode: AccessMode,
	})

	return session, func() {
		domain.HandleClosure(ctx, session, domain.ResourceNeo4j)
	}
}
