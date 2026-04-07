package usecase

import (
	"gin-go-api/internal/domain/entity"
	"gin-go-api/internal/domain/repository"
)

type HistoryUsecase struct {
	repo repository.HistoryRepository
}

func NewHistoryUsecase(repo repository.HistoryRepository) *HistoryUsecase {
	return &HistoryUsecase{repo: repo}
}

func (u *HistoryUsecase) GetAll() ([]entity.History, error) {
	return u.repo.FindAll()
}
