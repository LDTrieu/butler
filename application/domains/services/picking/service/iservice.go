package service

import (
	"butler/application/domains/services/picking/models"
	"context"
)

type IService interface {
	GetOne(ctx context.Context, params *models.GetRequest) (*models.Picking, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.Picking, error)
	Update(ctx context.Context, obj *models.Picking) (*models.Picking, error)
}
