package agent

import (
	"github.com/pontus-devoteam/agent-sdk-go/pkg/model/providers/lmstudio"
	"training-plan/internal/ai/config"
)

func NewLMStudioProvider(cfg *config.AIConfig) *lmstudio.Provider {
	provider := lmstudio.NewProvider()

	provider.SetBaseURL(cfg.LMStudioBaseURL)
	provider.SetDefaultModel(cfg.LMStudioModel)

	return provider
}
