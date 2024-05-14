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

	openaiModel string

	maxHistoryLength int
	chatHistory      map[string][]openai.ChatCompletionMessage
}

// NewChatGPTOptions are the options for creating a new ChatGPT instance
type NewChatGPTOptions struct {
	OpenAIKey        string
	MaxHistoryLength int
	OpenAIModel      string
	Prompts          []Prompt
}

// New creates a new ChatGPT instance
func New(options *NewChatGPTOptions) *ChatGPT {
	if options.MaxHistoryLength == 0 {
		options.MaxHistoryLength = 5
	}
	if options.OpenAIModel == "" {
		options.OpenAIModel = openai.GPT3Dot5Turbo
	}

	client := openai.NewClient(options.OpenAIKey)
	return &ChatGPT{
		client:           client,
		Prompts:          options.Prompts,
		maxHistoryLength: options.MaxHistoryLength,
		chatHistory:      make(map[string][]openai.ChatCompletionMessage),
		tools:            make(map[string]openai.Tool),
		toolHandlers:     make(map[string]func(common.User, map[string]any) (string, error)),
		openaiModel:      options.OpenAIModel,
	}
}
