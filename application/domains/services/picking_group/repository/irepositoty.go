package repository

import (
	"butler/application/domains/services/picking_group/models"
	"context"
)

type IRepository interface {
	Update(ctx context.Context, obj *models.PickingGroup) (*models.PickingGroup, error)
	UpdateMany(ctx context.Context, objs []*models.PickingGroup) error
	GetById(ctx context.Context, id int64) (*models.PickingGroup, error)
	GetOne(ctx context.Context, params *models.GetRequest) (*models.PickingGroup, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.PickingGroup, error)
}
