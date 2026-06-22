package routes

import "github.com/gin-gonic/gin"

//Inicialização das rotas
func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId")
	r.GET("/getUserByEmail/:userEmail")
	r.POST("/createaUser")
	r.PUT("/updateUser/:userId")
	r.DELETE("/deleteUser/:userId")

}
