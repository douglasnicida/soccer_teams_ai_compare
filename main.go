package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	database "gin-go-api/database"
	groq_router "gin-go-api/routes/groq"
	history_router "gin-go-api/routes/history"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, dbErr := database.DBConnection()

	if dbErr != nil {
		log.Fatal("Error connecting to database.")
	}

	log.Println("Database connected successfully!")

	router := gin.Default()
	app := router.Group("/api")

	groq_router.RegisterGroqRoutes(app)
	history_router.RegisterHistoryRoutes(app, db)

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
