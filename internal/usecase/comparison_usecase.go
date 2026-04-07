package usecase

import (
	"gin-go-api/internal/domain/entity"
	"gin-go-api/internal/domain/service"
)

type ComparisonUsecase struct {
	llm service.LLMService
}

func NewComparisonUsecase(llm service.LLMService) *ComparisonUsecase {
	return &ComparisonUsecase{llm: llm}
}

func (u *ComparisonUsecase) Execute(query entity.ComparisonQuery) (string, error) {
	return u.llm.Compare(query.Team1, query.Team2)
}
