package repository

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"vessel_service/model"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Vessel_respository interface {
	FindBySpec(ctx context.Context, spec *model.Specification) (*model.Vessel, error)
	Create(ctx context.Context, vessel *model.Vessel) (interface{}, error)
	FindAll(ctx context.Context) ([]*model.Vessel, error)
}
type vessel_Repository struct {
	loger      *logrus.Entry
	collection *mongo.Collection
}

func New_vessel_repository(loger *logrus.Entry, collection *mongo.Collection) Vessel_respository {
	loger.Logger.WithFields(logrus.Fields{"_Level": "service"})
	return &vessel_Repository{collection: collection, loger: loger}
}

func (vr *vessel_Repository) FindBySpec(ctx context.Context, spec *model.Specification) (*model.Vessel, error) {
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}
	vessel := model.Vessel{}
	if err := vr.collection.FindOne(ctx, filter).Decode(&vessel); err != nil {
		vr.loger.Error(err)
		return nil, err
	}
	vr.loger.Info("Found vessel :", vessel)
	return &vessel, nil
}

//Create creates a vessel and stores it in the db
func (vr *vessel_Repository) Create(ctx context.Context, vessel *model.Vessel) (interface{}, error) {
	results, err := vr.collection.InsertOne(ctx, vessel)
	if err != nil {
		vr.loger.Error(err)
		return nil, err
	}
	vr.loger.Info(results.InsertedID)
	return results.InsertedID, nil
}
func (vr *vessel_Repository) FindAll(ctx context.Context) ([]*model.Vessel, error) {
	var ret []*model.Vessel
	cursor, err := vr.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Print("unable to get the document cursor due to err: ", err)
		return nil, err
	}
	for cursor.Next(ctx) {
		var episode bson.M
		if err := cursor.Decode(&episode); err != nil {
			log.Println("unable to read from the document cursor due to err: ", err)
			return nil, err
		}
		fmt.Println(episode)

	}
	return ret, nil
}
func (vr *vessel_Repository) DeleteByCriteria(ctx context.Context, key string, value interface{}) error {
	filter := bson.D{{key, bson.D{{"$eq", value}}}}
	res, err := vr.collection.DeleteMany(ctx, filter)

	if err != nil {
		log.Println("unable to delete documents because of:", err)
		return err
	}
	fmt.Println("DeleteMany() result TYPE:", reflect.TypeOf(res))
	return nil
}
