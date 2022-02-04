package routes

import (
	"api/controllers"
	"api/middlewares"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run will start the server
func Run() {
	getRoutes()

	router.Run(":5000")

}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {

	router.Use(middlewares.Authenticate())
	router.POST("/login", controllers.Login)

	v1 := router.Group("/users")
	addUserRoutes(v1)
	v2 := router.Group("/categories")
	addCategoriesRoutes(v2)
	v3 := router.Group("/articles")
	addArticlesRoutes(v3)
}
