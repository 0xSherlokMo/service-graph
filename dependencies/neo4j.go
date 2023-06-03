package dependencies

import (
	"context"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (d *DP) WithNeo4j() *DP {
	uri := os.Getenv("NEO4J_URI")
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(os.Getenv("NEO4J_USERNAME"), os.Getenv("NEO4J_PASSWORD"), ""))
	if err != nil {
		log.Fatalf("error while connecting to neo4j with error %s", err.Error())
	}

	if err := driver.VerifyConnectivity(context.Background()); err != nil {
		log.Fatalf("Cannot connect to Neo4j with error %s", err.Error())
	}
	d.neo4jDriver = driver
	return d
}

func (d *DP) GetNeo4j() neo4j.DriverWithContext {
	if d.neo4jDriver == nil {
		d.WithNeo4j()
		return d.neo4jDriver
	}

	return d.neo4jDriver
}
