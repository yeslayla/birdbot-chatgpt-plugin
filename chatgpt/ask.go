package chatgpt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/sashabaranov/go-openai"
	"github.com/yeslayla/birdbot-common/common"
)

// Ask sends a message to the OpenAI API and returns the response
func (chat *ChatGPT) Ask(user common.User, message string, historyContext string) string {

	// Get chat history
	chatHistory := chat.chatHistory[historyContext]
	if chatHistory == nil {
		chatHistory = make([]openai.ChatCompletionMessage, 0)
	}

	messages := ConvertPrompts(chat.Prompts)
	newMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: fmt.Sprintf("%s: %s", user.DisplayName, message),
	}
	messages = append(messages, newMessage)

	// Add messages to chat history
	chatHistory = append(chatHistory, messages...)

	ctx := context.Background()
	res, err := chat.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:     chat.openaiModel,
		Messages:  chatHistory,
		MaxTokens: 500,
		User:      user.ID,

		Tools: chat.GetTools(),
	})
	if err != nil {
		log.Printf("Failed to send: %s", err)
		return ""
	}

	if len(res.Choices) == 0 {
		log.Printf("OpenAPI returned with no choices")
		return ""
	}

	// Format response
	prefix := "birdbot:"
	choice := res.Choices[0]
	content := choice.Message.Content
	if strings.HasPrefix(strings.ToLower(content), strings.ToLower(prefix)) {
		content = strings.TrimPrefix(content, prefix)
	}
	content = strings.TrimSpace(content)

	// Call tool handler
	if choice.Message.ToolCalls != nil {
		for _, toolCall := range choice.Message.ToolCalls {
			if toolCall.Type != openai.ToolTypeFunction {
				continue
			}

			params := make(map[string]any)
			_ = json.Unmarshal([]byte(toolCall.Function.Arguments), &params)

			result, err := chat.callHandler(user, toolCall.Function.Name, params)
			if err != nil {
				log.Printf("Failed to call tool handler: %s", err)
			}

			if result != "" {
				content += "\n" + result
			}
		}
	}

	// Trim content
	content = strings.TrimSpace(content)

	// Add response to history
	if choice.Message.ToolCalls == nil {
		chatHistory = append(chatHistory, choice.Message)
	}

	// Clear history older than maxHistoryLength
	if len(chatHistory) > chat.maxHistoryLength {
		chatHistory = chatHistory[len(chatHistory)-chat.maxHistoryLength:]
	}

	// Update chat history
	chat.chatHistory[historyContext] = chatHistory

	return content
}
