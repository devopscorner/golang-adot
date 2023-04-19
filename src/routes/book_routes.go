package routes

import (
	"github.com/devopscorner/golang-adot/src/controller"
	"github.com/devopscorner/golang-adot/src/middleware"
	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	// Login route to create basic auth JWT token
	router.POST("/v1/login", controller.LoginUser)

	api := router.Group("/v1", middleware.AuthMiddleware())
	{
		// Book routes
		api.GET("/books", controller.GetAllBooks)
		api.GET("/books/:id", controller.GetBookByID)
		api.POST("/books", controller.CreateBook)
		api.PUT("/books/:id", controller.UpdateBook)
		api.DELETE("/books/:id", controller.DeleteBook)
	}
}
