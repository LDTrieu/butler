package repository

import (
	"butler/application/domains/services/cart/models"
	"context"
)

type IRepository interface {
	Update(ctx context.Context, obj *models.Cart) (*models.Cart, error)
	UpdateMany(ctx context.Context, objs []*models.Cart) error
	GetById(ctx context.Context, id int64) (*models.Cart, error)
	GetOne(ctx context.Context, params *models.GetRequest) (*models.Cart, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.Cart, error)
}
