package service

type LLMService interface {
	Compare(team1, team2 string) (string, error)
}
