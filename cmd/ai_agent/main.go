package main

import (
	"github.com/pontus-devoteam/agent-sdk-go/pkg/runner"
	"training-plan/internal/ai/agent"
	"training-plan/internal/ai/config"
	handlers "training-plan/internal/ai/handler"
)

func main() {
	cfg := config.NewAIConfig()

	assistant := agent.NewTrainingPlanAgent(cfg)

	r := runner.NewRunner()
	r.WithDefaultProvider(agent.NewLMStudioProvider(cfg))

	conversationHandler := handlers.NewConversationHandler(assistant, r)
	conversationHandler.StartInteractiveSession()
}
