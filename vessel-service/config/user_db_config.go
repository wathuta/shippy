package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateConn(ctx context.Context, uri string, retries int32) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("DB Error %v", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		if retries >= 3 {
			return nil, err
		}
		retries++
		time.Sleep(time.Second * 3)
		return CreateConn(ctx, uri, retries)
	}
	return client, nil
}