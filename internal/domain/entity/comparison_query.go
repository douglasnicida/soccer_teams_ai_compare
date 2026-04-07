package entity

type ComparisonQuery struct {
	Team1 string `form:"team1" binding:"required"`
	Team2 string `form:"team2" binding:"required"`
}
