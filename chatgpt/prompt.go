package chatgpt

import (
	openai "github.com/sashabaranov/go-openai"
)

// Prompt is a struct that represents a ChatGPT message
type Prompt struct {
	Role string `yaml:"role"`
	Text string `yaml:"text"`
}

// PromptToMessage converts a prompt to a chatgpt.ChatMessage
func PromptToMessage(prompt Prompt) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    prompt.Role,
		Content: prompt.Text,
	}
}

// ConvertPrompts converts a slice of integration.Prompt to a slice of chatgpt.ChatMessage
func ConvertPrompts(prompts []Prompt) []openai.ChatCompletionMessage {
	messages := make([]openai.ChatCompletionMessage, len(prompts))
	for i, prompt := range prompts {
		messages[i] = PromptToMessage(prompt)
	}
	return messages
}
