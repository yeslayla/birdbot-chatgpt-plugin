package chatgpt

import (
	openai "github.com/sashabaranov/go-openai"
	"github.com/yeslayla/birdbot-common/common"
)

type ChatGPT struct {
	Prompts []Prompt

	client *openai.Client
	key    string

	tools        map[string]openai.Tool
	toolHandlers map[string]func(common.User, map[string]any) (string, error)

	maxHistoryLength int
	chatHistory      map[string][]openai.ChatCompletionMessage
}

// New creates a new ChatGPT instance
func New(key string, maxHistoryLength int, prompts []Prompt) *ChatGPT {
	if maxHistoryLength == 0 {
		maxHistoryLength = 5
	}

	client := openai.NewClient(key)
	return &ChatGPT{
		client:           client,
		Prompts:          prompts,
		maxHistoryLength: maxHistoryLength,
		chatHistory:      make(map[string][]openai.ChatCompletionMessage),
		tools:            make(map[string]openai.Tool),
		toolHandlers:     make(map[string]func(common.User, map[string]any) (string, error)),
	}
}
