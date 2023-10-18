package integration

import "strings"

type Config struct {
	OpenAIKey      string   `yaml:"open_ai_key" env:"OPENAI_KEY"`
	EnableCommand  Feature  `yaml:"enable_command"`
	EnableChat     Feature  `yaml:"enable_chat"`
	Prompts        []Prompt `yaml:"prompts"`
	ResponseChance float64  `yaml:"response_chance"`
}

// Prompt is a struct that represents a ChatGPT message
type Prompt struct {
	Role string `yaml:"role"`
	Text string `yaml:"text"`
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
