package dependencies

import (
	"context"
	"flag"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type DP struct {
	neo4jDriver neo4j.DriverWithContext
}

func NewDependencyInjection() *DP {
	neo4j := flag.String("neo4j", "neo4j://localhost:7687", "Used to define neo4j uri")
	neo4jUserName := flag.String("neo4jUsername", "neo4j", "username for neo4j")
	neo4jPassword := flag.String("neo4jPassword", "neo4j", "password for neo4j")
	flag.Parse()

	if os.Getenv("NEO4J_URI") == "" {
		os.Setenv("NEO4J_URI", *neo4j)
	}

	if os.Getenv("NEO4J_USERNAME") == "" {
		os.Setenv("NEO4J_USERNAME", *neo4jUserName)
	}

	if os.Getenv("NEO4J_PASSWORD") == "" {
		os.Setenv("NEO4J_PASSWORD", *neo4jPassword)
	}

	return &DP{}
}

func (d *DP) Shutdown() {
	ctx := context.Background()
	if d.neo4jDriver != nil {
		d.neo4jDriver.Close(ctx)
	}
}
