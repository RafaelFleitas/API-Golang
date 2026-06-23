package controller

import (
	"fmt"

	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model/request"
	"github.com/gin-gonic/gin"
)

// Controller do Create
func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest // Declara uma variável do tipo UserRequest para armazenar os dados da requisição

	// Tenta mapear o JSON do corpo da requisição para a struct userRequest. O ShouldBindJSON também valida os dados com base nas tags de 'binding' na struct.
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Existem campos incorretos, error %s\n", err.Error()))
		c.JSON(restErr.Code, restErr) // Retorna uma resposta JSON com o status HTTP do erro e o objeto de erro.
		return
	}

	// Imprime os dados recebidos no console (placeholder para a lógica de negócio).
	fmt.Println(userRequest)

}
