package init

import (
	cartInit "butler/application/domains/services/cart/init"
	cartSv "butler/application/domains/services/cart/service"
	packingInit "butler/application/domains/services/packing/init"
	packingSv "butler/application/domains/services/packing/service"
	pickingGroupInit "butler/application/domains/services/picking_group/init"
	pickingGroupSv "butler/application/domains/services/picking_group/service"
	promtAiSv "butler/application/domains/services/promt_ai/service"
	"butler/config"

	"github.com/google/generative-ai-go/genai"
	"gorm.io/gorm"
)

type Services struct {
	PromtAiSv           promtAiSv.IService
	CartService         cartSv.IService
	PackingService      packingSv.IService
	PickingGroupService pickingGroupSv.IService
}

func InitService(cfg *config.Config, db *gorm.DB, genaiClient *genai.Client) *Services {
	initPromtAiSv := promtAiSv.InitService(cfg, genaiClient)
	cart := cartInit.NewInit(db, cfg)
	packing := packingInit.NewInit(db, cfg)
	pickingGroup := pickingGroupInit.NewInit(db, cfg)

	return &Services{
		PromtAiSv:           initPromtAiSv,
		CartService:         cart.Service,
		PackingService:      packing.Service,
		PickingGroupService: pickingGroup.Service,
	}
}
