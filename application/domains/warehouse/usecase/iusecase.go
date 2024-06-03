package usecase

import (
	"butler/application/domains/warehouse/models"
	"context"
)

type IUseCase interface {
	ShowWarehouse(ctx context.Context, params *models.ShowWarehouseRequest) error
	ResetShowWarehouse(ctx context.Context) error
}
