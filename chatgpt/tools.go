package chatgpt

import (
	"log"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"github.com/yeslayla/birdbot-common/common"
)

// ToolerHandlerDefinion is a struct that represents a tool handler definition
type ToolerHandlerDefinion struct {
	Name        string
	Description string
	Parameters  []ToolHandlerParameter
}

// ToolHandlerParameter is a struct that represents a tool handler parameter
type ToolHandlerParameter struct {
	Name        string
	Description string
	Type        jsonschema.DataType
	Required    bool
}

// ChatGPT is a struct that represents a set of tools for ChatGPT
func (chat *ChatGPT) GetTools() []openai.Tool {
	toolList := make([]openai.Tool, 0, len(chat.tools))
	for _, tool := range chat.tools {
		toolList = append(toolList, tool)
	}
	return toolList
}

// callHandler calls a tool handler
func (chat *ChatGPT) callHandler(user common.User, name string, params map[string]any) (string, error) {
	handler, ok := chat.toolHandlers[name]
	if !ok {
		return "", nil
	}
	return handler(user, params)
}

// RegisterToolHandler registers a tool handler
func (chat *ChatGPT) RegisterToolHandler(definition ToolerHandlerDefinion, handler func(common.User, map[string]any) (string, error)) {
	chat.toolHandlers[definition.Name] = handler

	params := jsonschema.Definition{
		Type:       jsonschema.Object,
		Properties: make(map[string]jsonschema.Definition),
		Required:   make([]string, 0),
	}

	for _, parameter := range definition.Parameters {
		params.Properties[parameter.Name] = jsonschema.Definition{
			Type:        jsonschema.DataType(parameter.Type),
			Description: parameter.Description,
		}

		if parameter.Required {
			params.Required = append(params.Required, parameter.Name)
		}
	}

	chat.tools[definition.Name] = openai.Tool{
		Type: openai.ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        definition.Name,
			Description: definition.Description,
			Parameters:  params,
		},
	}

	log.Println("Registered tool handler:", definition.Name)
}
