package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hosseinmirzapur/oai/internal"
	"github.com/hosseinmirzapur/oai/services/openai_service"
)

func main() {
	// load .env file
	err := internal.LoadEnv()
	if err != nil {
		panic(err)
	}

	// get user prompt
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the file path:")
	scanner.Scan()
	filePath := scanner.Text()

	// generate text from audio file
	service := openai_service.New()
	fmt.Println("Generating text from audio file...")
	res, err := service.VoiceToText(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("file content:")
	fmt.Println(*res)
	fmt.Println("Done!")

}
