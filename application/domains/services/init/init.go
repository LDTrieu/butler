package init

import (
	promtAiSv "butler/application/domains/services/promt_ai/service"
	"butler/config"

	"github.com/google/generative-ai-go/genai"
)

type services struct {
	promtAiSv promtAiSv.IService
}

func InitService(cfg *config.Config, genaiClient *genai.Client) *services {
	initPromtAiSv := promtAiSv.InitService(cfg, genaiClient)
	return &services{
		promtAiSv: initPromtAiSv,
	}
}
