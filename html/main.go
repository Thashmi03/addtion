package main

import (
	"context"
	"fmt"
	"add/config"
	"add/constants"
	"add/controllers"
	"add/routes"
	"add/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initRoutes() {
	routes.Default(server)
}

func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	profileCollection := mongoClient.Database("operations").Collection("add")
	profileService := services.InitService(profileCollection, ctx)
	profileController := controllers.InitController(profileService)
	routes.Add(server,profileController)
}
func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDatabase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}