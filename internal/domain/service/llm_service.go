package service

type LLMService interface {
	Compare(team1, team2 string) (*CompareResult, error)
}

type CompareResult struct {
	Score    string
	Analysis string
}
