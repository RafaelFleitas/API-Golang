package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/controller"
	"github.com/RafaelFleitas/API-Golang/src/controller/routes"
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

	fmt.Println(os.Getenv("TEST"))

	// Inicializa o router, registra as rotas da aplicação e inicia o servidor na porta 8000

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default() //Vai registrar as rotas de aplicação recebida

	routes.InitRoutes(&router.RouterGroup, userController) // Pega a requisição HTTP do router e inicia
	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}

}
