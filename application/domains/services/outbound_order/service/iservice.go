package service

import (
	"butler/application/domains/services/outbound_order/models"
	"context"
)

type IService interface {
	GetOne(ctx context.Context, params *models.GetRequest) (*models.OutboundOrder, error)
	GetList(ctx context.Context, params *models.GetRequest) ([]*models.OutboundOrder, error)
	Update(ctx context.Context, obj *models.OutboundOrder) (*models.OutboundOrder, error)
	GetListOutboundItems(ctx context.Context, outboundOrderID int64) ([]*models.OutboundOrderItem, error)
}
