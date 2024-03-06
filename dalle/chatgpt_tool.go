package dalle

import (
	"errors"

	"github.com/sashabaranov/go-openai/jsonschema"
	"github.com/yeslayla/birdbot-chatgpt-plugin/chatgpt"
	"github.com/yeslayla/birdbot-common/common"
)

func RegisterDalleWithChatGPT(chat *chatgpt.ChatGPT, image *DALLE) {
	chat.RegisterToolHandler(chatgpt.ToolerHandlerDefinion{
		Name:        "image_generation",
		Description: "Generates an image from a prompt",
		Parameters: []chatgpt.ToolHandlerParameter{
			{
				Name:        "prompt",
				Description: "The prompt to generate the image from",
				Type:        jsonschema.String,
				Required:    true,
			},
		},
	}, func(user common.User, m map[string]any) (string, error) {
		prompt, ok := m["prompt"].(string)
		if !ok {
			return "", errors.New("prompt is required")
		}

		return image.Ask(user, prompt), nil
	})
}
