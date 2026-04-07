package database

import (
	"gin-go-api/internal/domain/entity"
	"gin-go-api/internal/domain/repository"

	"gorm.io/gorm"
)

type HistoryRepositoryImpl struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) repository.HistoryRepository {
	return &HistoryRepositoryImpl{db: db}
}

func (r *HistoryRepositoryImpl) Save(history *entity.History) error {
	return r.db.Create(history).Error
}

func (r *HistoryRepositoryImpl) FindAll() ([]entity.History, error) {
	var histories []entity.History
	err := r.db.Find(&histories).Error
	return histories, err
}
