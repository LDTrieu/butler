package service

import (
	"butler/application/domains/services/picking_item/models"
	"context"
)

type IService interface {
	GetOne(ctx context.Context, params *models.GetRequest) (*models.PickingItem, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.PickingItem, error)
	Update(ctx context.Context, obj *models.PickingItem) (*models.PickingItem, error)
}
