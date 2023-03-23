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
	os.Setenv("neo4j", *neo4j)
	os.Setenv("neo4jusername", *neo4jUserName)
	os.Setenv("neo4jpassword", *neo4jPassword)

	return &DP{}
}

func (d *DP) Shutdown() {
	ctx := context.Background()
	if d.neo4jDriver != nil {
		d.neo4jDriver.Close(ctx)
	}
}
