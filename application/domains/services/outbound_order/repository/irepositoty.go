package repository

import (
	"butler/application/domains/services/outbound_order/models"
	"context"
)

type IRepository interface {
	Update(ctx context.Context, obj *models.OutboundOrder) (*models.OutboundOrder, error)
	UpdateMany(ctx context.Context, objs []*models.OutboundOrder) error
	GetById(ctx context.Context, id int64) (*models.OutboundOrder, error)
	GetOne(ctx context.Context, params *models.GetRequest) (*models.OutboundOrder, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.OutboundOrder, error)
}
