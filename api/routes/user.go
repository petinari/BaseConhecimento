package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {

	rg.POST("/", controllers.CadastrarUsuarios)
	rg.GET("/:id", controllers.GetUserById)
	rg.GET("/", controllers.GetAllUsers)

}
