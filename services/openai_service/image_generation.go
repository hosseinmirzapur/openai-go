package openai_service

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func (oai *openAIService) ImageGeneration(prompt string) (*openai.ImageResponse, error) {
	respUrl, err := client.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:         prompt,
			Size:           openai.CreateImageSize256x256,
			ResponseFormat: openai.CreateImageResponseFormatURL,
			N:              2,
		},
	)

	if err != nil {
		return nil, err
	}

	return &respUrl, nil
}
