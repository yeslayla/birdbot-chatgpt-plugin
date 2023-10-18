package chatgpt

import (
	"github.com/ayush6624/go-chatgpt"
)

// Prompt is a struct that represents a ChatGPT message
type Prompt struct {
	Role string `yaml:"role"`
	Text string `yaml:"text"`
}

// PromptToMessage converts a prompt to a chatgpt.ChatMessage
func PromptToMessage(prompt Prompt) chatgpt.ChatMessage {
	return chatgpt.ChatMessage{
		Role:    chatgpt.ChatGPTModelRole(prompt.Role),
		Content: prompt.Text,
	}
}

// ConvertPrompts converts a slice of integration.Prompt to a slice of chatgpt.ChatMessage
func ConvertPrompts(prompts []Prompt) []chatgpt.ChatMessage {
	messages := make([]chatgpt.ChatMessage, len(prompts))
	for i, prompt := range prompts {
		messages[i] = PromptToMessage(prompt)
	}
	return messages
}
