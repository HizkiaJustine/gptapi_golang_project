package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetResponse(client *openai.Client, ctx context.Context, question string) {
	req := openai.ChatCompletionRequest{
        Model: openai.GPT3Babbage002,
        Messages: []openai.ChatCompletionMessage{
            {
                Role:    openai.ChatMessageRoleUser,
                Content: question,
            },
        },
        Stream: true,
    }
	stream, err := client.CreateChatCompletionStream(ctx, req)
    if err != nil {
        fmt.Printf("ChatCompletionStream error: %v\n", err)
        return
    }
    defer stream.Close()

    for {
        response, err := stream.Recv()
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Printf("Stream error: %v\n", err)
            break
        }
        fmt.Print(response.Choices[0].Delta.Content)
    }
    fmt.Print("\n")
}


func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apikey := viper.GetString("API_KEY")
	if apikey == "" {
		panic("missing API key")
	}

	ctx := context.Background()
	client := openai.NewClient(apikey)
	rootCmd := &cobra.Command{
		Use: "chatgpt",
		Short: "Chat with ChatGPT in console.",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Print("Type something ('quit' to end): ")
				if !scanner.Scan() {
					break
				}
				question := scanner.Text()
				switch question {
				case "quit":
					quit = true

				default:
					GetResponse(client, ctx, question)
				}
			}
		},
	}
	rootCmd.Execute()
}