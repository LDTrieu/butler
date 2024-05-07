package repository

import (
	"butler/application/domains/services/picking/models"
	"context"
)

type IRepository interface {
	Update(ctx context.Context, obj *models.Picking) (*models.Picking, error)
	UpdateMany(ctx context.Context, objs []*models.Picking) error
	GetById(ctx context.Context, id int64) (*models.Picking, error)
	GetOne(ctx context.Context, params *models.GetRequest) (*models.Picking, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.Picking, error)
}
