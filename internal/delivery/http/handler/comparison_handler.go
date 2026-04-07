package handler

import (
	"net/http"

	"gin-go-api/internal/domain/entity"
	"gin-go-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ComparisonHandler struct {
	usecase *usecase.ComparisonUsecase
}

func NewComparisonHandler(uc *usecase.ComparisonUsecase) *ComparisonHandler {
	return &ComparisonHandler{usecase: uc}
}

func (h *ComparisonHandler) Compare(ctx *gin.Context) {
	var query entity.ComparisonQuery

	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if query.Team1 == "" || query.Team2 == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Team 1 and Team 2 are required",
		})
		return
	}

	result, err := h.usecase.Execute(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"team1":  query.Team1,
		"team2":  query.Team2,
		"score":   result.Score,
		"result":  result.Analysis,
	})
}
