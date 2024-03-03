package chatgpt

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ayush6624/go-chatgpt"
	"github.com/yeslayla/birdbot-common/common"
)

func (chat *ChatGPT) Ask(user common.User, message string, historyContext string) string {

	// Get chat history
	chatHistory := chat.chatHistory[historyContext]
	if chatHistory == nil {
		chatHistory = make([]chatgpt.ChatMessage, 0)
	}

	messages := ConvertPrompts(chat.Prompts)
	newMessage := chatgpt.ChatMessage{
		Role:    chatgpt.ChatGPTModelRoleUser,
		Content: fmt.Sprintf("%s: %s", user.DisplayName, message),
	}
	messages = append(messages, newMessage)

	// Add messages to chat history
	chatHistory = append(chatHistory, newMessage)

	ctx := context.Background()
	res, err := chat.client.Send(ctx, &chatgpt.ChatCompletionRequest{
		Model:    chatgpt.GPT4,
		User:     user.ID,
		Messages: append(messages, chatHistory...),
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

	// Add response to history
	chatHistory = append(chatHistory, res.Choices[0].Message)

	// Clear history older than maxHistoryLength
	if len(chatHistory) > chat.maxHistoryLength {
		chatHistory = chatHistory[len(chatHistory)-chat.maxHistoryLength:]
	}

	// Update chat history
	chat.chatHistory[historyContext] = chatHistory

	return content
}
