package ai

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

const ModelYiCoder = "yi-coder"

type OllamaClient struct {
	client  *ollama.LLM
	history []string
}

type AnalyzeRequest struct {
	Input string `json:"input"`
}

type AnalyzeResponse struct {
	Result string `json:"result"`
}

func NewOllamaClient() (*OllamaClient, error) {
	client, err := ollama.New(ollama.WithModel(ModelYiCoder))
	if err != nil {
		return nil, err
	}

	return &OllamaClient{client: client, history: []string{}}, nil
}

func (c *OllamaClient) Analyze(input string) (string, error) {
	query := "very briefly, tell me the difference between a comet and a meteor"

	ctx := context.Background()

	// memory := memory.NewConversationBuffer()
	// llmChain := chains.NewConversation(c.client, memory)

	completion, err := llms.GenerateFromSinglePrompt(ctx, c.client, query)
	if err != nil {
		return "", err
	}

	fmt.Printf("Completion: %s\n", completion)

	return completion, nil
}
