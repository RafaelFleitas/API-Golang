package controller

import (
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

// Controller do Create
func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
