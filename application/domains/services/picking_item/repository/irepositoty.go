package repository

import (
	"butler/application/domains/services/picking_item/models"
	"context"
)

type IRepository interface {
	Update(ctx context.Context, obj *models.PickingItem) (*models.PickingItem, error)
	UpdateMany(ctx context.Context, objs []*models.PickingItem) error
	GetById(ctx context.Context, id int64) (*models.PickingItem, error)
	GetOne(ctx context.Context, params *models.GetRequest) (*models.PickingItem, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.PickingItem, error)
}
