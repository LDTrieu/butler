package service

import (
	"butler/application/domains/services/cart/models"
	"context"
)

type IService interface {
	GetOne(ctx context.Context, params *models.GetRequest) (*models.Cart, error)
	GetById(ctx context.Context, id int64) (*models.Cart, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.Cart, error)
	Update(ctx context.Context, obj *models.Cart) (*models.Cart, error)
}
