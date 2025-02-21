package service

import (
	"butler/application/domains/services/outbound_order/models"
	"context"
)

func (s *service) GetListOutboundItems(ctx context.Context, outboundOrderID int64) ([]*models.OutboundOrderItem, error) {
	items, err := s.repo.GetListOutboundItems(ctx, outboundOrderID)
	if err != nil {
		return nil, err
	}
	return items, nil
}
