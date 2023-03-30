package domain

import (
	"context"
	"log"
)

type (
	resource   string
	Empty      = struct{}
	CloserFunc = func()
	Comparable interface {
		~int | ~float32 | ~float64 | ~string
	}
	ResourceWithClosure interface {
		Close(ctx context.Context) error
	}
)

const (
	EmptyString    = ""
	SpaceDelimiter = " "
	ResourceNeo4j  = "[Neo4j]"
)

func HandleClosure(ctx context.Context, connection ResourceWithClosure, tag resource) {
	err := connection.Close(ctx)
	if err != nil {
		log.Println(tag, "Error while closing resource connection")
	}
}

func Assert[C Comparable](needle, haystick C) bool {
	return needle == haystick
}
