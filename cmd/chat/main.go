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
	fmt.Println("Enter your prompt:")
	scanner.Scan()
	prompt := scanner.Text()

	service := openai_service.New()
	res, err := service.ChatCompletion(prompt)
	if err != nil {
		panic(err)
	}

	fmt.Println("answer:", res.Choices[0].Message.Content)
}
