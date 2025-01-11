package ollama

import (
	"context"
	"github.com/ollama/ollama/api"
)

// https://github.com/ollama/ollmakeama/blob/main/api/client.go

type Client interface {
	Chat(ctx context.Context, req *api.ChatRequest, fn api.ChatResponseFunc) error
	Generate(ctx context.Context, req *api.GenerateRequest, fn api.GenerateResponseFunc) error
}

func NewClient() Client {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		panic(err)
	}
	return client
}
