package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/heisenxxd/crud-go/crud-go/src/configuration/logger"
	"github.com/heisenxxd/crud-go/crud-go/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting CRUD application")
	err := godotenv.Load()
	if err != nil {
		log.Println("erro ao carregar as variaveis de ambiente", err)
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}