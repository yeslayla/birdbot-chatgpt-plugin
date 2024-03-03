package chatgpt

import (
	"github.com/ayush6624/go-chatgpt"
)

type ChatGPT struct {
	Prompts []Prompt

	client *chatgpt.Client

	maxHistoryLength int
	chatHistory      map[string][]chatgpt.ChatMessage
}

func NewChatGPT(key string, prompts []Prompt) *ChatGPT {

	client, _ := chatgpt.NewClient(key)
	return &ChatGPT{
		client:           client,
		Prompts:          prompts,
		maxHistoryLength: 50,
		chatHistory:      make(map[string][]chatgpt.ChatMessage),
	}
}
