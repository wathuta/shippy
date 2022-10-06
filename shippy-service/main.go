package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wathuta/shippy-service/proto/consignment"
	"github.com/wathuta/shippy-service/proto/user"
	"github.com/wathuta/shippy-service/proto/vessel"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/wathuta/shippy-service/configs"
	"github.com/wathuta/shippy-service/repository"
	"github.com/wathuta/shippy-service/services"
	micro "go-micro.dev/v4"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
)

var (
	client  *mongo.Client
	service micro.Service
)

func init() {
	if err := godotenv.Load("vars.env"); err != nil {
		log.Printf("Loading env file error %v", err)
		log.Println("searching for environment variables")
	}
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		log.Fatal("Empty uri field")
	}
	var err error
	client, err = configs.CreateConn(context.Background(), uri, 0)
	if err != nil {
		log.Fatalf("Error unable to connect to mongo DB: %v ", err)
	}
}

func main() {
	service = micro.NewService(micro.Name("shippy.consignment"), micro.Version("latest"), micro.WrapHandler(AuthWrapper))
	service.Init()

	//to do add a collection and connect to mongodb

	collection := client.Database("shippy").Collection("consignments")
	defer client.Disconnect(context.Background())
	//
	repo := repository.New_shippy_repository(collection)
	vess := vessel.NewVesselService("shippy.vessel", service.Client())
	if vess == nil {
		log.Fatal(vess)
	}
	shippy_service := services.New_shippy_service(repo, vess)

	if err := consignment.RegisterShippingServiceHandler(service.Server(), shippy_service); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		token := meta["Token"]
		log.Println("Authenticating with token", token)

		authclient := user.NewUserService("shippy.user", service.Client())
		_, err := authclient.ValidateToken(context.Background(), &user.Token{Token: token})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
