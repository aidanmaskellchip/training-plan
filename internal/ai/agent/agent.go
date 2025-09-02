package agent

import (
	"github.com/pontus-devoteam/agent-sdk-go/pkg/agent"
	"training-plan/internal/ai/client"
	"training-plan/internal/ai/config"
	"training-plan/internal/ai/tools"
)

func NewTrainingPlanAgent(cfg *config.AIConfig) *agent.Agent {
	provider := NewLMStudioProvider(cfg)
	apiClient := client.NewAPIClient(cfg.APIBaseURL)

	assistant := agent.NewAgent("TrainingPlanAssistant")
	assistant.SetModelProvider(provider)
	assistant.WithModel(cfg.LMStudioModel)
	assistant.SetSystemInstructions(cfg.SystemInstructions)

	// Add all tools
	assistant.WithTools(
		tools.NewCreateUserTool(apiClient),
		tools.NewCreateRunningProfileTool(apiClient),
		// Add more tools as needed
	)

	return assistant
}
