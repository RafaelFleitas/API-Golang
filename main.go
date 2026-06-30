package main

import (
	"log"

	oraclesql "github.com/RafaelFleitas/API-Golang/src/configuration/database/oracleSQL"
	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logger.Info("About to start the application")
	godotenv.Load()

	//Inicializa o banco de dados ORACLE
	db, err := oraclesql.NewOracleConnection()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	userController := initDependencies(db)

	router := gin.Default()                                //Vai registrar as rotas de aplicação recebida
	routes.InitRoutes(&router.RouterGroup, userController) // Pega a requisição HTTP do router e inicia
	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}

}
