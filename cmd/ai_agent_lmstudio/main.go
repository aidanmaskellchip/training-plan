package main

import (
	"github.com/pontus-devoteam/agent-sdk-go/pkg/runner"
	"training-plan/internal/ai_lmstudio/agent"
	"training-plan/internal/ai_lmstudio/config"
	handlers "training-plan/internal/ai_lmstudio/handler"
)

func main() {
	cfg := config.NewAIConfig()

	assistant := agent.NewTrainingPlanAgent(cfg)

	r := runner.NewRunner()
	r.WithDefaultProvider(agent.NewLMStudioProvider(cfg))

	conversationHandler := handlers.NewConversationHandler(assistant, r)
	conversationHandler.StartInteractiveSession()
}
