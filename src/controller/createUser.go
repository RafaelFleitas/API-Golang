package controller

import (
	"net/http"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/validation"
	"github.com/RafaelFleitas/API-Golang/src/model/request"
	"github.com/RafaelFleitas/API-Golang/src/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Controller do Create
func CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest // Declara uma variável do tipo UserRequest para armazenar os dados da requisição

	// ShouldBindJson garante que os dados recebidos estão no formato correto
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Init CreateUser controller", err,
			zap.String("journey", "createUser"),
		)
		restErr := validation.ValidateUserError(err) //Chama a função de validação de erros em validation

		c.JSON(restErr.Code, restErr) // Retorna uma resposta JSON com o status HTTP do erro e o objeto de erro.
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, response)

}
