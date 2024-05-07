package usecase

import (
	"butler/application/domains/pick/models"
	initServices "butler/application/domains/services/init"
	invenoryModel "butler/application/domains/services/inventory/models"
	invenorySv "butler/application/domains/services/inventory/service"
	outboundModel "butler/application/domains/services/outbound_order/models"
	outboundOrderSv "butler/application/domains/services/outbound_order/service"
	pickingModel "butler/application/domains/services/picking/models"
	pickingSv "butler/application/domains/services/picking/service"
	pickingItemModel "butler/application/domains/services/picking_item/models"
	pickingItemSv "butler/application/domains/services/picking_item/service"
	"context"
	"fmt"
	"strings"
	"time"
)

type usecase struct {
	outboundOrderSv outboundOrderSv.IService
	pickingSv       pickingSv.IService
	pickingItemSv   pickingItemSv.IService
	invenorySv      invenorySv.IService
}

func InitUseCase(
	services *initServices.Services,
) IUseCase {
	return &usecase{
		pickingSv:       services.PickingService,
		pickingItemSv:   services.PickingItemService,
		invenorySv:      services.InventoryService,
		outboundOrderSv: services.OutboundOrderService,
	}
}

const DEFAULT_LOCATION = "F0-AP-00-00-00-01"

func (u *usecase) ReadyPickOutbound(ctx context.Context, params *models.ReadyPickOutboundRequest) error {
	outbound, err := u.outboundOrderSv.GetOne(ctx, &outboundModel.GetRequest{SalesOrderNumber: params.SalesOrderNumber})
	if err != nil {
		return err
	}
	if outbound == nil || outbound.OutboundOrderId == 0 {
		return fmt.Errorf("không tìm thấy mã outbound với source code [%s]", params.SalesOrderNumber)
	}

	if outbound.StatusId != 1 {
		return fmt.Errorf("outbound order [%s] không ở trạng thái picklisted", params.SalesOrderNumber)
	}
	if outbound.Config != 0 {
		if outbound.Config == 1 {
			return fmt.Errorf("outbound order [%s] không đủ hàng đi pick", params.SalesOrderNumber)
		}
		return fmt.Errorf("outbound order [%s] không đủ điều kiện pick", params.SalesOrderNumber)
	}

	picking, err := u.pickingSv.GetOne(ctx, &pickingModel.GetRequest{OutboundOrderId: outbound.OutboundOrderId})
	if err != nil {
		return err
	}
	if picking != nil && picking.PickingId != 0 {
		return fmt.Errorf("outbound order [%s] không có picking", params.SalesOrderNumber)
	}
	if picking.StatusId != 1 {
		return fmt.Errorf("picking [%s] không ở trạng thái được pick", picking.PickingNumber)
	}

	pickingItems, err := u.pickingItemSv.GetList(ctx, &pickingItemModel.GetRequest{PickingId: picking.PickingId})
	if err != nil {
		return err
	}
	if len(pickingItems) == 0 {
		return fmt.Errorf("outbound [%s] không có hàng để pick", picking.PickingNumber)
	}

	updateLocationPickingItemIds := make([]int64, 0)
	inventoryIds := make([]int64, 0)
	for _, item := range pickingItems {
		if outbound.OutboundOrderType == "INTERNAL_TRANSFER" {
			if !checkSuitableLocationForIt(item.LocationDescription) {
				updateLocationPickingItemIds = append(updateLocationPickingItemIds, item.PickingItemId)
			}
		} else if outbound.OutboundOrderType == "ORDER" {
			if !checkSuitableLocationForOrder(item.LocationDescription) {
				updateLocationPickingItemIds = append(updateLocationPickingItemIds, item.PickingItemId)
			}
		} else {
			return fmt.Errorf("loại outbound order [%s] không hợp lệ: %s ", params.SalesOrderNumber, outbound.OutboundOrderType)
		}
		if item.StatusId != 1 {
			inventoryIds = append(inventoryIds, item.InventoryId)
		}
	}
	if len(inventoryIds) == 0 {
		return fmt.Errorf("outbound [%s] không có item phù hợp để pick", picking.PickingNumber)
	}

	inventories, err := u.invenorySv.GetList(ctx, outbound.WarehouseId, &invenoryModel.GetRequest{InventoryIds: inventoryIds})
	if err != nil {
		return err
	}
	for _, pi := range updateLocationPickingItemIds {
		if _, err := u.pickingItemSv.Update(ctx, &pickingItemModel.PickingItem{
			PickingItemId:       pi,
			LocationDescription: DEFAULT_LOCATION,
		}); err != nil {
			return err
		}
	}

	// update

	if time.Until(outbound.CreatedAt).Abs().Minutes() < 10 {
		if _, err := u.outboundOrderSv.Update(ctx, &outboundModel.OutboundOrder{
			OutboundOrderId: outbound.OutboundOrderId,
			UpdatedAt:       time.Now().Add(-10 * time.Minute),
		}); err != nil {
			return err
		}
	}

	for _, inv := range inventories {
		isUpdate := false
		updatedInventory := &invenoryModel.Inventory{
			InventoryId: inv.InventoryId,
		}
		if outbound.OutboundOrderType == "INTERNAL_TRANSFER" {
			if !checkSuitableLocationForIt(inv.LocationDescription) {
				isUpdate = true
				updatedInventory.LocationDescription = DEFAULT_LOCATION
			}
		} else if outbound.OutboundOrderType == "ORDER" {
			if !checkSuitableLocationForOrder(inv.LocationDescription) {
				isUpdate = true
				updatedInventory.LocationDescription = DEFAULT_LOCATION
			}
		}

		if inv.StatusId != 7 {
			isUpdate = true
			updatedInventory.StatusId = 7
		}

		if inv.AreaId != 22 {
			isUpdate = true
			updatedInventory.AreaId = 22
		}

		if isUpdate {
			if _, err := u.invenorySv.Update(ctx, outbound.WarehouseId, updatedInventory); err != nil {
				return err
			}
		}
	}

	return nil
}

func checkSuitableLocationForIt(location string) bool {
	if strings.Contains(location, "AP") {
		return true
	}
	if strings.Contains(location, "PO") {
		return true
	}
	if strings.Contains(location, "PG") {
		return true
	}
	if strings.Contains(location, "IT") {
		return true
	}
	if strings.Contains(location, "ST") {
		return true
	}
	if strings.Contains(location, "NG") {
		return true
	}

	return false
}
func checkSuitableLocationForOrder(location string) bool {
	return strings.Contains(location, "AP")
}
