package openai_service

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func fileIsAudio(filePath string) (error, bool) {
	// Extract the file extension
	ext := strings.ToLower(filepath.Ext(filePath))

	// List of audio file extensions
	audioExtensions := []string{".mp3", ".wav", ".flac", ".aac"}

	// Check if the file extension matches the audio extensions
	for _, audioExt := range audioExtensions {
		if ext == audioExt {
			return nil, true
		}
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err, false
	}
	defer file.Close()

	// Read the first few bytes from the file
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return err, false
	}

	// Determine the audio file format based on the file content
	fileType := http.DetectContentType(buffer)

	if !strings.HasPrefix(fileType, "audio/") {
		return errors.New("file is not an audio"), false
	}

	return nil, true
}
