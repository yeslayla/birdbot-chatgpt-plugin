package integration

import "strings"

type Config struct {
	OpenAIKey      string  `yaml:"open_ai_key" env:"OPENAI_KEY"`
	EnableCommand  Feature `yaml:"enable_command"`
	EnableChat     Feature `yaml:"enable_chat"`
	Prompt         string  `yaml:"prompt"`
	ResponseChance float64 `yaml:"response_chance"`
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
