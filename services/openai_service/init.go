package openai_service

import (
	"github.com/hosseinmirzapur/oai/config"
	"github.com/sashabaranov/go-openai"
)

var client *openai.Client

func New() *openAIService {
	client = openai.NewClient(config.GetOpenAIConfig().Api_key)
	return &openAIService{}
}
