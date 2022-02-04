package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func addCategoriesRoutes(rg *gin.RouterGroup) {

	rg.POST("/", controllers.SaveCategories)
	rg.GET("/", controllers.GetAllCategories)
}
