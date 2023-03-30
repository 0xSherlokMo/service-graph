package domain

import (
	"context"
	"log"
)

type Empty = struct{}

type resource string

type Comparable interface {
	~int | ~float32 | ~float64 | ~string
}

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

func Assert[C Comparable](needle, haystick C) bool {
	return needle == haystick
}
