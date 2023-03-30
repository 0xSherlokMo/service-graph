package domain

import (
	"context"
	"log"
)

type Empty = struct{}

type resource string

const (
	EmptyString    = ""
	SpaceDelimiter = " "
	ResourceNeo4j  = "[Neo4j]"
)

type ResourceWithClosure interface {
	Close(ctx context.Context) error
}

type CloserFunc = func()

func HandleClosure(ctx context.Context, connection ResourceWithClosure, tag resource) {
	err := connection.Close(ctx)
	if err != nil {
		log.Println(tag, "Error while closing resource connection")
	}
}
