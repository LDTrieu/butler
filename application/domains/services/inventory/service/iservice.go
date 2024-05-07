package service

import (
	"butler/application/domains/services/inventory/models"
	"context"
)

type IService interface {
	GetOne(ctx context.Context, warehouseId int64, params *models.GetRequest) (*models.Inventory, error)
	GetList(ctx context.Context, warehouseId int64, params *models.GetRequest) ([]*models.Inventory, error)
	Update(ctx context.Context, warehouseId int64, obj *models.Inventory) (*models.Inventory, error)
	UpdateMany(ctx context.Context, warehouseId int64, obj []*models.Inventory) error
}
