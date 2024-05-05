package repository

import (
	"butler/application/domains/services/packing/models"
	"context"
)

type IRepository interface {
	Update(ctx context.Context, obj *models.Packing) (*models.Packing, error)
	UpdateMany(ctx context.Context, objs []*models.Packing) error
	GetById(ctx context.Context, id int64) (*models.Packing, error)
	GetOne(ctx context.Context, params *models.GetRequest) (*models.Packing, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.Packing, error)
}
