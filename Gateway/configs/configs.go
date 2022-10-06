package configs

import (
	"os"
	"time"
)

type Config struct {
	Host           string
	Port           string
	User           string
	Password       string
	Exchange       string
	Queue          string
	RoutingKey     string
	ConsumerTag    string
	WorkerPoolSize int
}
type ServerConfigs struct {
	Port    string
	Timeout time.Time
	Retries int
}

func NewRabbitConfig() *Config {
	return &Config{
		Host:       os.Getenv("HOST"),
		Port:       os.Getenv("PORT"),
		User:       os.Getenv("USER"),
		Password:   os.Getenv("PASSWORD"),
		Exchange:   os.Getenv("EXCHANGE"),
		Queue:      os.Getenv("QUEUE"),
		RoutingKey: os.Getenv("ROUTINGKEY")}
}
func NewServerConfigs() *ServerConfigs {
	return &ServerConfigs{
		Port: os.Getenv("SERVER_PORT")}
}
