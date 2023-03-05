package openaiutils

import (
	"context"

	"github.com/EdgeJay/psg-navi-bot/articles-upload/env"
	"github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAIClient() *OpenAIClient {
	apiKey := env.GetOpenAiApiKey()
	client := openai.NewClient(apiKey)
	return &OpenAIClient{
		client: client,
	}
}

func (c *OpenAIClient) PerformTextCompletion(prompt string) (string, error) {
	ctx := context.Background()
	req := openai.CompletionRequest{
		Model:       openai.GPT3TextDavinci003,
		MaxTokens:   1024,
		Prompt:      prompt,
		Temperature: 0,
		User:        "12edc093-e516-48fa-aa88-d2c3b22ef633",
	}

	res, err := c.client.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	return res.Choices[0].Text, nil
}
