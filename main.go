package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nolanprewit1/employee_directory_api/database"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Global variables
var err error
var Router *gin.Engine

func main() {
	// Initialize Database
	database.InitializeDatabase()

	// Set the router to the default gin router
	Router := gin.Default()

	// Allow requests for all origins
	// In future versions this will be updated to further refine access / security
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE"}

	Router.Use(cors.New(config))

	// Define the Directory api router group
	apiDirectory := Router.Group("/api/directory")
	{
		apiDirectory.POST("/", postContact)
		apiDirectory.GET("/", getContacts)
		apiDirectory.GET("/:id", getContact)
		apiDirectory.PUT("/:id", updateContact)
		apiDirectory.DELETE("/:id", deleteContact)
	}

	// Define a test route
	Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Employee Directory API is working!",
		})
	})

	// Start the web server
	Router.Run(":8080")
}
