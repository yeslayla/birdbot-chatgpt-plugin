package dalle

import "github.com/sashabaranov/go-openai"

type DALLE struct {
	client *openai.Client
	model  string
}

func New(key string, model string) *DALLE {
	if model == "" {
		model = openai.CreateImageModelDallE3
	}

	client := openai.NewClient(key)
	return &DALLE{
		client: client,
		model:  model,
	}
}
