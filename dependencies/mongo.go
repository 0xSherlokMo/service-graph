package dependencies

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGO_MAX_CONNECTION_POOL = 100
	MONGO_MIN_CONNECTION_POOL = 2
	IDLE_TIME_MS              = 0
)

func (d *DP) WithMongo() *DP {
	ctx := context.Background()
	mongodbClient, err := mongo.NewClient(
		options.Client().ApplyURI(os.Getenv("MONGO_URL")),
		options.Client().SetMaxPoolSize(MONGO_MAX_CONNECTION_POOL),
		options.Client().SetMinPoolSize(MONGO_MIN_CONNECTION_POOL),
		options.Client().SetMaxConnIdleTime(IDLE_TIME_MS),
	)
	err = mongodbClient.Connect(ctx)
	if err != nil {
		log.Fatalf("Cannot connect to MongoDB due to: %s", err)
	}
	err = mongodbClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("failed to ping mongo: %s", err)
	}
	d.mongoClient = mongodbClient
	return d
}

func (d *DP) GetMongo() *mongo.Client {
	if d.mongoClient == nil {
		d.WithMongo()
	}

	return d.mongoClient
}
