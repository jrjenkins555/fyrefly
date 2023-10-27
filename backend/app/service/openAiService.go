package service

import (
	"context"
	"fmt"
	"github.com/otiai10/openaigo"
)

func CallOpenAI() error {
	client := openaigo.NewClient("OPEN_AI_KEY")
	request := openaigo.ChatRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openaigo.Message{
			{Role: "user", Content: "Hello!"},
		},
	}
	ctx := context.Background()
	response, err := client.Chat(ctx, request)
	fmt.Println(response, err)
	return nil
}
