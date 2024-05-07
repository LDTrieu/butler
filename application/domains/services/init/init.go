package init

import (
	cartInit "butler/application/domains/services/cart/init"
	cartSv "butler/application/domains/services/cart/service"
	inventoryInit "butler/application/domains/services/inventory/init"
	inventorySv "butler/application/domains/services/inventory/service"
	outboundOrderInit "butler/application/domains/services/outbound_order/init"
	outboundOrderSv "butler/application/domains/services/outbound_order/service"
	packingInit "butler/application/domains/services/packing/init"
	packingSv "butler/application/domains/services/packing/service"
	pickingInit "butler/application/domains/services/picking/init"
	pickingSv "butler/application/domains/services/picking/service"
	pickingGroupInit "butler/application/domains/services/picking_group/init"
	pickingGroupSv "butler/application/domains/services/picking_group/service"
	pickingItemInit "butler/application/domains/services/picking_item/init"
	pickingItemSv "butler/application/domains/services/picking_item/service"
	promtAiSv "butler/application/domains/services/promt_ai/service"
	"butler/config"

	"github.com/google/generative-ai-go/genai"
	"gorm.io/gorm"
)

type Services struct {
	PromtAiSv            promtAiSv.IService
	CartService          cartSv.IService
	PackingService       packingSv.IService
	PickingGroupService  pickingGroupSv.IService
	OutboundOrderService outboundOrderSv.IService
	PickingService       pickingSv.IService
	PickingItemService   pickingItemSv.IService
	InventoryService     inventorySv.IService
}

func InitService(cfg *config.Config, db *gorm.DB, genaiClient *genai.Client) *Services {
	initPromtAiSv := promtAiSv.InitService(cfg, genaiClient)
	cart := cartInit.NewInit(db, cfg)
	packing := packingInit.NewInit(db, cfg)
	pickingGroup := pickingGroupInit.NewInit(db, cfg)
	outboundOrder := outboundOrderInit.NewInit(db, cfg)
	picking := pickingInit.NewInit(db, cfg)
	pickingItem := pickingItemInit.NewInit(db, cfg)
	inventory := inventoryInit.NewInit(db, cfg)

	return &Services{
		PromtAiSv:            initPromtAiSv,
		CartService:          cart.Service,
		PackingService:       packing.Service,
		PickingGroupService:  pickingGroup.Service,
		OutboundOrderService: outboundOrder.Service,
		PickingService:       picking.Service,
		PickingItemService:   pickingItem.Service,
		InventoryService:     inventory.Service,
	}
}
