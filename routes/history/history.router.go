package history_router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHistoryRoutes(router *gin.RouterGroup, db *gorm.DB) {
	app := router.Group("history")

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "History",
		})
	})
}
