package entity

import "gorm.io/gorm"

type History struct {
	gorm.Model
	Team1  string `gorm:"column:team1"`
	Team2  string `gorm:"column:team2"`
	Result string `gorm:"column:result"`
}
