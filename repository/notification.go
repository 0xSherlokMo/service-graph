package repository

import (
	"context"

	"github.com/graduation-fci/service-graph/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DESC_ORDER         = -1
	REPORTS_COLLECTION = "reports"
)

func (repository *DrugRespository) LatestReport(medicationId int64) (domain.Report, error) {
	ctx := context.Background()
	query := bson.M{
		"medicationId": medicationId,
	}

	orderOpt := options.FindOne().SetSort(bson.M{"_id": DESC_ORDER})

	var report domain.Report
	err := repository.
		dp.
		GetMongo().
		Collection(REPORTS_COLLECTION).
		FindOne(ctx, query, orderOpt).
		Decode(&report)
	if err != nil {
		return report, err
	}

	return report, nil
}

func (repository *DrugRespository) SaveReport(report domain.Report) error {
	ctx := context.Background()
	_, err := repository.
		dp.
		GetMongo().
		Collection(REPORTS_COLLECTION).
		InsertOne(ctx, report)
	return err
}
