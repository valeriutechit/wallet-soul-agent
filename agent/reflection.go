package agent

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"wallet-soul-agent/utils"
)

func GenerateReflection(tokens []utils.Token, profile string) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	content := fmt.Sprintf(
		"This is a Solana wallet profile: %s. It contains tokens: %v. Please write a short poetic reflection about the soul behind this wallet.",
		profile,
		extractSymbols(tokens),
	)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: "You are a mystical philosopher of the blockchain."},
				{Role: "user", Content: content},
			},
		},
	)

	if err != nil || len(resp.Choices) == 0 {
		return "🌀 The soul remains silent in the face of digital uncertainty."
	}

	return resp.Choices[0].Message.Content
}

func extractSymbols(tokens []utils.Token) string {
	symbols := ""
	for _, t := range tokens {
		if t.UiAmount > 0 {
			symbols += t.Symbol + ", "
		}
	}
	return symbols
}
