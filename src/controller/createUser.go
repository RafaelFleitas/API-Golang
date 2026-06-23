package controller

import (
	"log"

	"github.com/RafaelFleitas/API-Golang/src/configuration/validation"
	"github.com/RafaelFleitas/API-Golang/src/model/request"
	"github.com/gin-gonic/gin"
)

// Controller do Create
func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest // Declara uma variável do tipo UserRequest para armazenar os dados da requisição

	// ShouldBindJson garante que os dados recebidos estão no formato correto
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Println("Error trying to marshal object", err.Error()) //Passa que foi impossivel transformar o objeto em json
		restErr := validation.ValidateUserError(err)               //Chama a função de validação de erros em validation

		c.JSON(restErr.Code, restErr) // Retorna uma resposta JSON com o status HTTP do erro e o objeto de erro.
		return
	}

}
