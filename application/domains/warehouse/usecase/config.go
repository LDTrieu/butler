package usecase

import (
	whModels "butler/application/domains/services/warehouse/models"
	"butler/application/domains/warehouse/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func (u *usecase) AddConfigWarehouse(ctx context.Context, params *models.UpdateConfigWarehouseRequest) error {
	if params.WarehouseId == 0 {
		return fmt.Errorf("kho không tồn tại")
	}

	wh, err := u.whSv.GetOne(ctx, &whModels.GetRequest{
		WarehouseId: params.WarehouseId,
	})
	if err != nil {
		return err
	}
	if wh == nil || wh.WarehouseName == "" {
		return fmt.Errorf("kho không tồn tại")
	}

	if err := u.whSv.UpdateWithMap(ctx, wh.WarehouseId, map[string]any{
		"config": gorm.Expr("config | ?", params.Config),
	}); err != nil {
		return err
	}

	return nil
}

func (u *usecase) RemoveConfigWarehouse(ctx context.Context, params *models.UpdateConfigWarehouseRequest) error {
	if params.WarehouseId == 0 {
		return fmt.Errorf("kho không tồn tại")
	}

	wh, err := u.whSv.GetOne(ctx, &whModels.GetRequest{
		WarehouseId: params.WarehouseId,
	})
	if err != nil {
		return err
	}
	if wh == nil || wh.WarehouseName == "" {
		return fmt.Errorf("kho không tồn tại")
	}

	if err := u.whSv.UpdateWithMap(ctx, wh.WarehouseId, map[string]any{
		"config": wh.Config &^ int64(params.Config),
	}); err != nil {
		return err
	}

	return nil
}

func (u *usecase) GetWarehouseById(ctx context.Context, warehouseId int64) (*whModels.Warehouse, error) {
	wh, err := u.whSv.GetOne(ctx, &whModels.GetRequest{
		WarehouseId: warehouseId,
	})
	if err != nil {
		return nil, err
	}
	if wh == nil || wh.WarehouseName == "" {
		return nil, fmt.Errorf("kho không tồn tại")
	}

	return wh, nil
}
