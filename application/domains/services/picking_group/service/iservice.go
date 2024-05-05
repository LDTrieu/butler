package service

import (
	"butler/application/domains/services/picking_group/models"
	"context"
)

type IService interface {
	GetOne(ctx context.Context, params *models.GetRequest) (*models.PickingGroup, error)
	GetById(ctx context.Context, id int64) (*models.PickingGroup, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.PickingGroup, error)
	Update(ctx context.Context, obj *models.PickingGroup) (*models.PickingGroup, error)
}
