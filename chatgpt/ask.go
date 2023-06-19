package chatgpt

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ayush6624/go-chatgpt"
	"github.com/yeslayla/birdbot-common/common"
)

func (chat *ChatGPT) Ask(user common.User, message string) string {

	ctx := context.Background()
	res, err := chat.client.Send(ctx, &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT35Turbo0301,
		User:  user.ID,
		Messages: []chatgpt.ChatMessage{
			{
				Role:    chatgpt.ChatGPTModelRoleSystem,
				Content: chat.Prompt,
			},
			{
				Role:    chatgpt.ChatGPTModelRoleUser,
				Content: fmt.Sprintf("%s: %s", user.DisplayName, message),
			},
		},
	})
	if err != nil {
		log.Printf("Failed simple send: %s", err)
		return ""
	}

	if len(res.Choices) == 0 {
		log.Printf("OpenAPI returned with no choices")
		return ""
	}

	// Format response
	prefix := "birdbot:"
	content := res.Choices[0].Message.Content
	if strings.HasPrefix(strings.ToLower(content), strings.ToLower(prefix)) {
		content = strings.TrimPrefix(content, prefix)
	}
	content = strings.TrimSpace(content)

	return content
}
