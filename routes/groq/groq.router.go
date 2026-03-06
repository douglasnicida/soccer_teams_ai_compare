package groq_router

import (
	groq_service "gin-go-api/services"

	"github.com/gin-gonic/gin"
)

// http://localhost:8080/api/ia/?team1=Internacional 2015&team2=Bahia 2015
func RegisterUserRoutes(router *gin.RouterGroup) {
	app := router.Group("ia")

	app.GET("/", groq_service.GetPromptComparisonResult)
}
