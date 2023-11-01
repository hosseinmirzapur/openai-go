package config

import (
	"os"

	"github.com/hosseinmirzapur/oai/types"
)

func GetOpenAIConfig() *types.OpenAIConfig {
	return &types.OpenAIConfig{
		Api_key: os.Getenv("OPENAI_API_KEY"),
	}
}
