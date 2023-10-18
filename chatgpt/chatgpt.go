package chatgpt

import (
	"github.com/ayush6624/go-chatgpt"
	"github.com/yeslayla/birdbot-chatgpt-plugin/integration"
)

type ChatGPT struct {
	Prompts []integration.Prompt

	client *chatgpt.Client
}

func NewChatGPT(key string, prompts []integration.Prompt) *ChatGPT {

	client, _ := chatgpt.NewClient(key)
	return &ChatGPT{
		client:  client,
		Prompts: prompts,
	}
}
