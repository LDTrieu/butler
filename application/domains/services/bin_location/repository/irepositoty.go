package repository

import (
	"butler/application/domains/services/bin_location/models"
	"context"
)

type IRepository interface {
	Update(ctx context.Context, obj *models.BinLocation) (*models.BinLocation, error)
	Create(ctx context.Context, obj *models.BinLocation) (*models.BinLocation, error)
	UpdateMany(ctx context.Context, objs []*models.BinLocation) error
	GetById(ctx context.Context, id int64) (*models.BinLocation, error)
	GetOne(ctx context.Context, params *models.GetRequest) (*models.BinLocation, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.BinLocation, error)
}
