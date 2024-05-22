package service

import (
	"butler/application/domains/services/bin_location/models"
	"context"
)

type IService interface {
	GetOne(ctx context.Context, params *models.GetRequest) (*models.BinLocation, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.BinLocation, error)
	Update(ctx context.Context, obj *models.BinLocation) (*models.BinLocation, error)
	Create(ctx context.Context, obj *models.BinLocation) (*models.BinLocation, error)
}
