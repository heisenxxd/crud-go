package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/heisenxxd/crud-go/crud-go/src/configuration/database/mongodb"
	"github.com/heisenxxd/crud-go/crud-go/src/configuration/logger"
	"github.com/heisenxxd/crud-go/crud-go/src/controller"
	"github.com/heisenxxd/crud-go/crud-go/src/controller/routes"
	"github.com/heisenxxd/crud-go/crud-go/src/model/service"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting CRUD application")
	err := godotenv.Load()
	if err != nil {
		log.Println("erro ao carregar as variaveis de ambiente", err)
	}

	mongodb.InitConnection()

	// init dependecies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}