package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ollama/ollama/api"
)

type OllamaProvider struct {
	client *api.Client
}

func NewOllama(endpointURL string) (*OllamaProvider, error) {
	url, err := url.Parse(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse endpoint URL: %w", err)
	}
	client := api.NewClient(url, http.DefaultClient)
	return &OllamaProvider{
		client: client,
	}, nil
}

func (p *OllamaProvider) Embed(text string) ([]float64, error) {
	ctx := context.Background()
	embedRequest := api.EmbeddingRequest{
		Model:  "deepseek-r1:14b",
		Prompt: text,
	}

	response, err := p.client.Embeddings(ctx, &embedRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to generate embeddings: %w", err)
	}

	return response.Embedding, nil
}

func main() {
	provider, err := NewOllama("http://localhost:11434")
	if err != nil {
		log.Fatalf("failed to create Ollama provider: %v", err)
	}

	text := "Hello, world!"
	embeddings, err := provider.Embed(text)
	if err != nil {
		log.Fatalf("failed to generate embeddings: %v", err)
	}

	fmt.Printf("Embeddings: %v\n", embeddings)
}
