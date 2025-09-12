package config

import (
	"os"
	"strings"
)

type AIConfig struct {
	LMStudioBaseURL    string
	LMStudioModel      string
	APIBaseURL         string
	SystemInstructions string
}

func NewAIConfig() *AIConfig {
	return &AIConfig{
		LMStudioBaseURL:    "http://127.0.0.1:1234/v1",
		LMStudioModel:      "gpt-oss-20b",
		APIBaseURL:         "http://localhost:4000/v1",
		SystemInstructions: loadSystemInstructions(),
	}
}

func loadSystemInstructions() string {
	content, err := os.ReadFile("internal/ai/config/system_instructions.txt")
	if err != nil {
		// Fallback to default if the file can't be read
		return "You are a helpful assistant for application designed to generate personalized running training plans"
	}
	return strings.TrimSpace(string(content))
}
