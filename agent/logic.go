package agent

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func ArchetypeFromBalance(balance float64) (string, string) {
	switch {
	case balance < 0.5:
		return "Wanderer", "A lonely soul wandering the digital desert, barely surviving."
	case balance < 5:
		return "Seeker", "Searching for meaning through scarce tokens, their journey has just begun."
	case balance < 20:
		return "Collector", "With careful hands, they accumulate value—not too much, not too little."
	case balance < 100:
		return "Alchemist", "Transmuting SOL into meaning, their path is intentional."
	default:
		return "Whale", "A mighty force in the ocean of tokens, shaping the currents with their will."
	}
}

func GenerateReflectionWithOpenAI(archetype string, balance float64) string {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "🤖 No reflection available (missing OpenAI key)"
	}

	client := openai.NewClient(apiKey)
	prompt := fmt.Sprintf("Write a poetic reflection in 2–3 lines for a Solana wallet archetype '%s' holding %.2f SOL", archetype, balance)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: "You are a mystical poet of the blockchain."},
				{Role: "user", Content: prompt},
			},
		},
	)

	if err != nil || len(resp.Choices) == 0 {
		return "🤖 Failed to generate poetic reflection."
	}

	fmt.Println("RESP, GPT", resp);
	return resp.Choices[0].Message.Content
}
