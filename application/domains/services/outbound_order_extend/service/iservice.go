package service

import (
	"butler/application/domains/services/outbound_order_extend/models"
	"context"
)

type IService interface {
	GetOne(ctx context.Context, params *models.GetRequest) (*models.OutboundOrderExtend, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.OutboundOrderExtend, error)
	Update(ctx context.Context, obj *models.OutboundOrderExtend) (*models.OutboundOrderExtend, error)
}
