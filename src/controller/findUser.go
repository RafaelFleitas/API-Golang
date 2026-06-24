package controller

import "github.com/gin-gonic/gin"

//O gin.Context tem todas as informações da request
//Controller dos Find
func (uc *userControllerInterface) FindUserByID(c *gin.Context)    {}
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {}
