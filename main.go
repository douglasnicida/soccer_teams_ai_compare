package main

import (
	"log"

	"gin-go-api/internal/config"
	"gin-go-api/internal/delivery/http/handler"
	httprouter "gin-go-api/internal/delivery/http/router"
	"gin-go-api/internal/domain/entity"
	"gin-go-api/internal/infrastructure/database"
	"gin-go-api/internal/infrastructure/groq"
	"gin-go-api/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cfg := config.Load()

	db := database.Connect("localhost", cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresDB)

	db.AutoMigrate(&entity.History{})

	groqClient := groq.NewGroqClient(cfg.GroqAPIKey)

	historyRepo := database.NewHistoryRepository(db)

	comparisonUsecase := usecase.NewComparisonUsecase(groqClient)
	historyUsecase := usecase.NewHistoryUsecase(historyRepo)

	comparisonHandler := handler.NewComparisonHandler(comparisonUsecase)
	historyHandler := handler.NewHistoryHandler(historyUsecase)

	r := gin.Default()

	apiGroup := r.Group("")
	httprouter.RegisterRoutes(apiGroup, comparisonHandler, historyHandler)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
