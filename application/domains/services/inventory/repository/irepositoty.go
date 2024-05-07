package repository

import (
	"butler/application/domains/services/inventory/models"
	"context"
)

type IRepository interface {
	Update(ctx context.Context, warehouseId int64, obj *models.Inventory) (*models.Inventory, error)
	UpdateMany(ctx context.Context, warehouseId int64, objs []*models.Inventory) error
	GetById(ctx context.Context, warehouseId int64, id int64) (*models.Inventory, error)
	GetOne(ctx context.Context, warehouseId int64, params *models.GetRequest) (*models.Inventory, error)
	GetList(ctx context.Context, warehouseId int64, params *models.GetRequest) ([]*models.Inventory, error)
}
