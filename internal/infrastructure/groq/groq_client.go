package groq

import (
	"context"
	"fmt"
	"strings"

	"gin-go-api/internal/domain/service"

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

func (c *GroqClient) Compare(team1, team2 string) (*service.CompareResult, error) {
	messages := buildMessages(team1, team2)

	resp, err := c.llm.GenerateContent(context.Background(), messages)
	if err != nil {
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no result was returned")
	}

	return parseResponse(resp.Choices[0].Content), nil
}

func parseResponse(text string) *service.CompareResult {
	const marker = "===SEPARATOR==="
	idx := strings.Index(text, marker)

	if idx == -1 {
		return &service.CompareResult{Analysis: text}
	}

	score := strings.TrimSpace(text[:idx])
	analysis := strings.TrimSpace(text[idx+len(marker):])

	return &service.CompareResult{Score: score, Analysis: analysis}
}

func buildMessages(team1, team2 string) []llms.MessageContent {
	agentDescription := "You are an expert soccer historian with deep knowledge of clubs, players, coaches, tactics, and competitions worldwide across all eras. You must always answer in Brazilian Portuguese."

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
	return fmt.Sprintf(`Compare os seguintes clubes de futebol (se houver número no nome, refira-se àquela temporada específica, senão fale do clube em geral):

- Clube A: %s
- Clube B: %s

IMPORTANTE: Na PRIMEIRA LINHA da sua resposta, coloque apenas o placar previsto no formato exato "X - Y" (onde X é o número de gols do Clube A e Y do Clube B). Por exemplo: "2 - 1". Não escreva mais nada nessa linha.
Após essa linha, escreva a seguinte comparação:

===SEPARATOR===

## 1. Elenco & Jogadores Chave
Liste os jogadores mais importantes de cada time. Inclua posição, nacionalidade e o que os tornou excepcionais naquela temporada.

## 2. Títulos da temporada
Que títulos, competições ou resultados notáveis cada clube conquistou naquele ano ou campanha?

## 3. Veredito da comparação
Se esses dois elencos jogassem entre si em campo neutro, quem provavelmente venceria e por quê? Considere confronto tático, qualidade individual e profundidade do elenco.

Seja específico. Use nomes reais, estatísticas e fatos históricos reais. Se não tiver certeza sobre algum detalhe, diga em vez de inventar.`,
		team1, team2,
	)
}
