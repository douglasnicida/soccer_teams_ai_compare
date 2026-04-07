package repository

import "gin-go-api/internal/domain/entity"

type HistoryRepository interface {
	Save(history *entity.History) error
	FindAll() ([]entity.History, error)
}
