package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"log"
	proto "user-client/proto/user"

	"github.com/sirupsen/logrus"
	micro "go-micro.dev/v4"
)

var logger = logrus.New()

var contextLogger = logger.WithFields(logrus.Fields{
	"service": "vessel service",
})

func init() {
	contextLogger.Logger.Out = createFileHelper()
	contextLogger.Logger.SetReportCaller(true)

	contextLogger.Logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
}
func createFileHelper() *os.File {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// logger.Fatal(err)
		contextLogger.Error(err)
		contextLogger.Warn(err)
		os.Exit(1)
	}
	return file
}
func createUser(ctx context.Context, client proto.UserService) error {
	resp, err := client.CreateUser(ctx, &proto.User{Name: "brian Wambug", Company: "Tangazo letu", Email: "mercy@gmail.com", Password: "CChangez13115!@"})
	if err != nil {
		logger.Println(err)
		return err
	}
	fmt.Println(resp.User)
	return nil
}
func getAllUsers(ctx context.Context, client proto.UserService) []*proto.User {
	resp, err := client.GetAll(ctx, &proto.Request{})
	if err != nil {
		logger.Println(err)
		return nil
	}
	return resp.Users
}
func AuthUser(ctx context.Context, client proto.UserService) (*proto.Token, error) {
	resp, err := client.AuthUser(ctx, &proto.User{Email: "wathutabrian@gmail.com", Password: "CChangez13115!@"})
	if err != nil {
		return nil, err
	}
	return resp, err
}
func ValidateToken(ctx context.Context, client proto.UserService, token string) error {
	resp, err := client.ValidateToken(ctx, &proto.Token{Token: token})
	if err != nil {
		return err
	}
	if !resp.Valid {
		return errors.New("invalid user")
	}
	return nil
}

func main() {
	service := micro.NewService(micro.Name("user-service"))
	service.Init()

	client := proto.NewUserService("user-service", service.Client())
	 if err:=createUser(context.Background(), client);err!=nil{
	 	log.Println(err)
	 }
//	 users := getAllUsers(context.Background(), client)
//	 Token, err := AuthUser(context.Background(), client)
//	 if err != nil {
//	 	log.Println(err)
//	 	return
//	 }
//	 fmt.Println(Token.Token)
///	if err := ValidateToken(context.Background(), client, token); err != nil {
		// logger.Fatal(err)
//		contextLogger.Error("DB logs", err)
//		contextLogger.Warn("DB logs", err)
//		return
//	}
	fmt.Println("Valid user")
}

//var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImVtYWlsIjoid2F0aHV0YWJyaWFuQGdtYWlsLmNvbSIsInBhc3N3b3JkIjoiQ0NoYW5nZXoxMzExNSFAIn0sImV4cCI6MTY1MjgwNTU1OCwiaXNzIjoic2hpcHB5X3NlcnZpY2UifQ.ieFznQjCrtTIG5eBJOodrjnlQN0100TmAb_sl40qlU0"
//
//Hey
//I have read your job description and it seems like you are looking for software developers with skills in golang and a background in C++ and some experience in rust for your projects.
//I am a perfect fit for the job. My three years of experience writing software in golang will come a long way in  helping you to develop software solutions for your organisation. And yes my upwork job success rating is not that high because i'm relatively new to upwork and ask that you consider the attached projects as a measure of my ability.

//I have experience working at an entreprise level setting and using agile development methodology. If you were to hire me you would get access to my system design skills, my knowledge about block chain, my experience working as a technical product owner and my golang coding skills

//I would love and will appreciate a chance to work with you.

//Best regards
//Brian Wambugu


//github.com/wathuta

//grpc
//goGin
//goMicro
//gorillla
//testify

//Attached are golang APIs that i've been working on . They  utilize different architectures of software development ranging from microservice to monoliths ,  Event driven architecture to interservice communication.
//Two of the attached api utilize RPC communication protocol for internal communication and the other is a traditional restful api.
