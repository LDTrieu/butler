package repository

import (
	"butler/application/domains/services/outbound_order_extend/models"
	"context"
)

type IRepository interface {
	Update(ctx context.Context, obj *models.OutboundOrderExtend) (*models.OutboundOrderExtend, error)
	UpdateMany(ctx context.Context, objs []*models.OutboundOrderExtend) error
	GetById(ctx context.Context, id int64) (*models.OutboundOrderExtend, error)
	GetOne(ctx context.Context, params *models.GetRequest) (*models.OutboundOrderExtend, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.OutboundOrderExtend, error)
}
