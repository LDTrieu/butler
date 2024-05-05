package service

import (
	"butler/application/domains/services/packing/models"
	"context"
)

type IService interface {
	GetOne(ctx context.Context, params *models.GetRequest) (*models.Packing, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.Packing, error)
	Update(ctx context.Context, obj *models.Packing) (*models.Packing, error)
}
