package integration

import (
	"strings"

	"github.com/yeslayla/birdbot-chatgpt-plugin/chatgpt"
)

type Config struct {
	OpenAIKey             string           `yaml:"open_ai_key" env:"OPENAI_KEY"`
	EnableCommand         Feature          `yaml:"enable_command"`
	EnableChat            Feature          `yaml:"enable_chat"`
	EnableImageGeneration Feature          `yaml:"enable_image_generation"`
	Prompts               []chatgpt.Prompt `yaml:"prompts"`
	ChatGPTModel          string           `yaml:"chatgpt_model"`
	DalleModel            string           `yaml:"dalle_model"`
	ResponseChance        float64          `yaml:"response_chance"`
	HistoryLength         int              `yaml:"history_length"`
}

// Feature is a boolean string used to toggle functionality
type Feature string

// IsEnabled returns true when a feature is set to be true
func (value Feature) IsEnabled() bool {
	return strings.ToLower(string(value)) == "true"
}

// IsEnabled returns true when a feature is set to be true
// or if the feature flag is not set at all
func (value Feature) IsEnabledByDefault() bool {
	v := strings.ToLower(string(value))
	if v == "" {
		v = "true"
	}
	return Feature(v).IsEnabled()
}
