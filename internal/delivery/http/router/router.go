package router

import (
	"gin-go-api/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.RouterGroup,
	comparisonHandler *handler.ComparisonHandler,
	historyHandler *handler.HistoryHandler,
) {
	api := r.Group("/api")
	{
		api.GET("/ia", comparisonHandler.Compare)
		api.GET("/history", historyHandler.GetAll)
		api.GET("/health-check", healthCheck)
	}
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "ok"})
}
