package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RafaelFleitas/API-Golang/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//Carrega as variaveis de ambiente (env) para main
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("TEST"))

	// Inicializa o router, registra as rotas da aplicação e inicia o servidor na porta 8000
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}

}
