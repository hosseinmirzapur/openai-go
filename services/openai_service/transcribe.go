package openai_service

import (
	"context"
	"errors"
	"os"

	"github.com/sashabaranov/go-openai"
)

func (oai *openAIService) VoiceToText(filePath string) (*string, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("please provide an audio file to convert to text")
	}

	if err, isAudio := fileIsAudio(filePath); !isAudio {
		return nil, err
	}

	resp, err := client.CreateTranscription(
		context.Background(),
		openai.AudioRequest{
			Model:    openai.Whisper1,
			FilePath: filePath,
		},
	)
	if err != nil {
		return nil, err
	}

	return &resp.Text, nil
}
