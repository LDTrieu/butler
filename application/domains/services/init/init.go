package init

import (
	promtAiSv "butler/application/domains/services/promt_ai/service"
	"butler/config"

	"github.com/google/generative-ai-go/genai"
)

type Services struct {
	PromtAiSv promtAiSv.IService
}

func InitService(cfg *config.Config, genaiClient *genai.Client) *Services {
	initPromtAiSv := promtAiSv.InitService(cfg, genaiClient)
	return &Services{
		PromtAiSv: initPromtAiSv,
	}
}
