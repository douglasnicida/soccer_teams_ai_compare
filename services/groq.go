package groq_service

import (
	"fmt"
	"net/http"
	"os"

	groq_model "gin-go-api/models/groq"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"

	"github.com/gin-gonic/gin"
)

const groq_url = "https://api.groq.com/openai/v1"

func newGroqLLM() (*openai.LLM, error) {
	return openai.New(
		openai.WithModel("llama-3.3-70b-versatile"),
		openai.WithBaseURL(groq_url),
		openai.WithToken(os.Getenv("GROQ_API_KEY")),
	)
}

func GetPromptComparisonResult(ctx *gin.Context) {
	llm, err := newGroqLLM()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var query groq_model.ComparisonQuery

	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	team1 := query.Team1
	team2 := query.Team2

	if team1 == "" || team2 == "" {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": "Team 1 and Team 2 are required"})
		return
	}

	content := buildMessage(team1, team2)

	result, err := getGroqResponse(llm, ctx, content)

	ctx.JSON(http.StatusOK, gin.H{
		"team1":  team1,
		"team2":  team2,
		"result": result,
	})
}

func getGroqResponse(llm *openai.LLM, ctx *gin.Context, content []llms.MessageContent) (*llms.ContentResponse, error) {
	result, err := llm.GenerateContent(ctx, content)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)

		return nil, err
	}

	if len(result.Choices) == 0 {
		ctx.JSON(
			http.StatusNoContent,
			gin.H{"error": "No result was returned"},
		)
		return nil, err
	}

	return result, nil
}

func buildMessage(team1 string, team2 string) []llms.MessageContent {
	agent_description := "You are an expert soccer historian with deep knowledge of clubs, players, coaches, tactics, and competitions worldwide across all eras."

	content := []llms.MessageContent{
		{
			Role: llms.ChatMessageTypeSystem,
			Parts: []llms.ContentPart{
				llms.TextPart(agent_description),
			},
		},
		{
			Role: llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{
				llms.TextPart(buildPrompt(team1, team2)),
			},
		},
	}
	return content
}

func buildPrompt(team1 string, team2 string) string {
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
