package openai_service

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func (oai *openAIService) ChatCompletion(prompt string) (*openai.ChatCompletionResponse, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: prompt,
				},
			},
			MaxTokens: 100,
		},
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
