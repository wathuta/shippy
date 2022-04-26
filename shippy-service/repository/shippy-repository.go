package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/wathuta/shippy-service/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type shippy_repository struct {
	Collection *mongo.Collection
}
type Shippy_repository interface {
	GetConsignments(ctx context.Context, filter interface{}) ([]*models.Consignment, error)
	CreateConsignment(ctx context.Context, _consignment *models.Consignment) (interface{}, error)
}

func New_shippy_repository(client *mongo.Collection) Shippy_repository {
	return &shippy_repository{Collection: client}
}

func (r *shippy_repository) GetConsignments(ctx context.Context, filter interface{}) ([]*models.Consignment, error) {
	cur, err := r.Collection.Find(ctx, filter, nil)
	var consignments []*models.Consignment
	var _consignment *models.Consignment
	for cur.Next(ctx) {
		if err := cur.Decode(&_consignment); err != nil {
			return nil, fmt.Errorf("decoding error* %v", err)
		}
		log.Println(_consignment)
		consignments = append(consignments, _consignment)
	}
	return consignments, err
}

func (r *shippy_repository) CreateConsignment(ctx context.Context, _consignment *models.Consignment) (interface{}, error) {
	result, err := r.Collection.InsertOne(ctx, _consignment)
	if err != nil {
		return nil, err
	}
	result_id := result.InsertedID
	return result_id, nil
}
