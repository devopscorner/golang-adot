package routes

import (
	"fmt"
	"net/http"

	"github.com/devopscorner/golang-adot/src/config"
	"github.com/devopscorner/golang-adot/src/driver"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Load Config
	config.LoadConfig()

	// Initialize Logs
	config.InitLogger()

	// Connect to database
	driver.ConnectDatabase()

	// Routes Healthcheck
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Routes Welcome
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Simple API Bookstore with ADOT (AWS Distro fro OpenTelemetry)!")
	})

	// Telemetry Routes
	TelemetryRoutes(router)

	// Book Routes
	BookRoutes(router)

	// Run the server
	port := fmt.Sprintf(":%v", config.AppPort())
	router.Run(port)
}
