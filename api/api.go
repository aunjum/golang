package api

import (
	"GOLANG/api/middleware"
	"GOLANG/api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func StartApi() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	if err != nil {
		return
	}
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	// API-1
	router.GET("/api-1", func(c *gin.Context) {
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	// API-2
	router.GET("/api-2", func(c *gin.Context) {
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	err = router.Run(":" + port)
	if err != nil {
		return
	}
}
