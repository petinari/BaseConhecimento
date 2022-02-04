package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func addArticlesRoutes(rg *gin.RouterGroup) {

	rg.POST("/", controllers.SaveArticles)
	rg.GET("/", controllers.GetAllArticles)
	rg.GET("/:id", controllers.GetArticlesById)
	rg.GET("/bycategory/:cat", controllers.GetArticlesByCategory)
}
