package main

import (
	"gateway/configs"
	"gateway/providers"
	"gateway/rabbitmq"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var vesselQueue *rabbitmq.VesselQueue

func init() {
	if err := godotenv.Load("configs.env"); err != nil {
		log.Fatalln(err)
	}
	vesselQueue = rabbitmq.NewRabbitMqConn(configs.NewRabbitConfig())
}

func main() {
	startServer()
}

//startServer Registers handlefuncs and starts a listener on port
func startServer() {
	app := gin.Default()

	vesselhandler := providers.ProvideVesselHandler(Makelogger(), vesselQueue)
	vesselRouter := app.Group("/vessel")
	{
		vesselRouter.POST("/create", vesselhandler.CreateVessel)
		vesselRouter.GET("/getall", vesselhandler.GetAllVessels)
	}
	httpConfigs := configs.NewServerConfigs()
	log.Println("Run HTTP Server in Port: %s", httpConfigs.Port)
	if err := app.Run(":" + httpConfigs.Port); err != nil {
		log.Fatalln("Unable to start server", err)
	}
}
func Makelogger() *log.Logger {
	file, err := os.Create("logs.log")
	checkErr(err)
	return log.New(file, "Gateway", log.LstdFlags)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
