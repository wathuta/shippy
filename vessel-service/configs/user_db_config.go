package configs

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host        string
	Port        string
	User        string
	Password    string
	Exchange    string
	Queue       string
	RoutingKey  string
	ConsumerTag string
	// WorkerPoolSize int
}
type ServerConfigs struct {
	Port    string
	Timeout time.Time
	Retries int
}

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
func NewRabbitConfigs() *Config {
	return &Config{
		Host:       os.Getenv("HOST"),
		Port:       os.Getenv("PORT"),
		User:       os.Getenv("USER"),
		Password:   os.Getenv("PASSWORD"),
		Exchange:   os.Getenv("EXCHANGE"),
		Queue:      os.Getenv("QUEUE"),
		RoutingKey: os.Getenv("ROUTINGKEY"),
	}
}
