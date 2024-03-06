package dalle

import "github.com/sashabaranov/go-openai"

type DALLE struct {
	client *openai.Client
}

func New(key string) *DALLE {
	client := openai.NewClient(key)
	return &DALLE{
		client: client,
	}
}
