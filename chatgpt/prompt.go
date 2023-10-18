package chatgpt

import (
	"github.com/ayush6624/go-chatgpt"
	"github.com/yeslayla/birdbot-chatgpt-plugin/integration"
)

// PromptToMessage converts a prompt to a chatgpt.ChatMessage
func PromptToMessage(prompt integration.Prompt) chatgpt.ChatMessage {
	return chatgpt.ChatMessage{
		Role:    chatgpt.ChatGPTModelRole(prompt.Role),
		Content: prompt.Text,
	}
}

// ConvertPrompts converts a slice of integration.Prompt to a slice of chatgpt.ChatMessage
func ConvertPrompts(prompts []integration.Prompt) []chatgpt.ChatMessage {
	messages := make([]chatgpt.ChatMessage, len(prompts))
	for i, prompt := range prompts {
		messages[i] = PromptToMessage(prompt)
	}
	return messages
}
