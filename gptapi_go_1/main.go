package main

import (
	"context"
	"fmt"
	"log"
	"os"
	openai "github.com/sashabaranov/go-openai"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		log.Fatalln("missing API key")
	}

	ctx := context.Background()
	client := openai.NewClient(apikey)

	resp, err := client.CreateCompletion(ctx, openai.CompletionRequest{
		Model:     openai.GPT3Babbage002,
		MaxTokens: 30,
		Prompt:    "What you should know about golang is",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text)
}