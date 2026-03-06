package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	groq_router "gin-go-api/routes/groq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	app := router.Group("/api")

	groq_router.RegisterUserRoutes(app)

	app.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Health Check Successful",
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
