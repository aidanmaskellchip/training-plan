package config

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
		SystemInstructions: "You are a helpful assistant for application designed to generate personalized running training plans...",
	}
}
