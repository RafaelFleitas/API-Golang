package main

import (
	"log"

	oraclesql "github.com/RafaelFleitas/API-Golang/src/configuration/database/oracleSQL"
	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/controller"
	"github.com/RafaelFleitas/API-Golang/src/controller/routes"
	"github.com/RafaelFleitas/API-Golang/src/model/repository"
	"github.com/RafaelFleitas/API-Golang/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logger.Info("About to start the application")
	//Carrega as variaveis de ambiente (env) para main
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//Inicializa o banco de dados ORACLE
	db, err := oraclesql.NewOracleConnection()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	// Inicializa o router, registra as rotas da aplicação e inicia o servidor na porta 8000
	userRepository := repository.NewUserRepository(db)
	service := service.NewUserDomainService(userRepository)

	userController := controller.NewUserControllerInterface(service)

	router := gin.Default() //Vai registrar as rotas de aplicação recebida

	routes.InitRoutes(&router.RouterGroup, userController) // Pega a requisição HTTP do router e inicia
	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}

}
