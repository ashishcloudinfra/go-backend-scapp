package helpers

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"server.simplifycontrol.com/secrets"
)

// PromtResponse sends a chat request to OpenAI with a given model and messages
func GetGPTResponse(model string, messages []openai.ChatCompletionMessage) (string, error) {
	// Create OpenAI client
	client := openai.NewClient(secrets.SecretJSON.OpenAI.Secret)

	// Prepare request
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    model,    // Dynamic model selection
			Messages: messages, // Dynamic message list
		},
	)
	if err != nil {
		return "", fmt.Errorf("API request failed: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}
