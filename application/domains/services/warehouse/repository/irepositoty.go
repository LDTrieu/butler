package repository

import (
	"butler/application/domains/services/warehouse/models"
	"context"
)

type IRepository interface {
	Update(ctx context.Context, obj *models.Warehouse) (*models.Warehouse, error)
	GetById(ctx context.Context, id int64) (*models.Warehouse, error)
	GetOne(ctx context.Context, params *models.GetRequest) (*models.Warehouse, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.Warehouse, error)
}
