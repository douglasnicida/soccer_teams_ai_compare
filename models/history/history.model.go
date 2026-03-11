package history_model

import (
	"gorm.io/gorm"
)

type HistoryModel struct {
	gorm.Model
	team1  string
	team2  string
	result string
}
