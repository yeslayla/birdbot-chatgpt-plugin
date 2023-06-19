package chatgpt

import "github.com/ayush6624/go-chatgpt"

type ChatGPT struct {
	Prompt string

	client *chatgpt.Client
}

func NewChatGPT(key string, prompt string) *ChatGPT {

	client, _ := chatgpt.NewClient(key)
	return &ChatGPT{
		client: client,
		Prompt: prompt,
	}
}
