package groq

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

type GroqClient struct {
	llm *openai.LLM
}

func NewGroqClient(apiKey string) *GroqClient {
	llm, err := openai.New(
		openai.WithToken(apiKey),
		openai.WithModel("llama-3.3-70b-versatile"),
		openai.WithBaseURL("https://api.groq.com/openai/v1"),
	)
	if err != nil {
		panic(err)
	}
	return &GroqClient{llm: llm}
}

func (c *GroqClient) Compare(team1, team2 string) (string, error) {
	messages := buildMessages(team1, team2)

	resp, err := c.llm.GenerateContent(context.Background(), messages)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no result was returned")
	}

	return resp.Choices[0].Content, nil
}

func buildMessages(team1, team2 string) []llms.MessageContent {
	agentDescription := "You are an expert soccer historian with deep knowledge of clubs, players, coaches, tactics, and competitions worldwide across all eras."

	return []llms.MessageContent{
		{
			Role: llms.ChatMessageTypeSystem,
			Parts: []llms.ContentPart{
				llms.TextPart(agentDescription),
			},
		},
		{
			Role: llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{
				llms.TextPart(buildPrompt(team1, team2)),
			},
		},
	}
}

func buildPrompt(team1, team2 string) string {
	return fmt.Sprintf(`
		Compare the following two soccer clubs from their specific historical seasons (if it has the number, if not, take in general):

		- Club A: %s
		- Club B: %s

		Provide a structured comparison with the following sections:

		## 1. Elenco & Jogadores Chave
		List the most important players for each club that year. Include their position, nationality, and what made them exceptional in that season.

		## 2. Títulos da temporada
		What titles, competitions, or notable results did each club achieve in that year or campaign? Include domestic league, cups, and continental competitions.

		## 3. Veredito da comparação
		If these two squads played each other in a neutral venue, who would likely win and why? Consider the tactical matchup, individual quality, and squad depth. Give a predicted score.

		Be specific. Use real names, real statistics, and real historical facts. If you are uncertain about a specific detail, say so rather than inventing it.
		Return the answer in brazilian portuguese`,
		team1, team2,
	)
}
