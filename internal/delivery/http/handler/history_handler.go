package handler

import (
	"net/http"

	"gin-go-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type HistoryHandler struct {
	usecase *usecase.HistoryUsecase
}

func NewHistoryHandler(uc *usecase.HistoryUsecase) *HistoryHandler {
	return &HistoryHandler{usecase: uc}
}

func (h *HistoryHandler) GetAll(ctx *gin.Context) {
	histories, err := h.usecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": histories,
	})
}
