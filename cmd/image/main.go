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

	// generate image
	service := openai_service.New()
	res, err := service.ImageGeneration(prompt)
	fmt.Println("Generating image...")
	if err != nil {
		panic(err)
	}

	fmt.Println("your image url:", res.Data[0].URL)
}
