package usecase

import (
	"butler/application/domains/pick/models"
	outboundModel "butler/application/domains/services/outbound_order/models"
	"butler/constants"
	"context"
	"fmt"
)

func (u *usecase) PreparePick(ctx context.Context) error {
	return nil
}

func (u *usecase) Pick(ctx context.Context, params *models.PickRequest) error {
	if err := u.lib.Validator.Struct(params); err != nil {
		return err
	}

	outbound, err := u.outboundOrderSv.GetOne(ctx, &outboundModel.GetRequest{SalesOrderNumber: params.SalesOrderNumber})
	if err != nil {
		return err
	}
	if outbound.StatusId != constants.OUTBOUND_ORDER_STATUS_PICKLISTED {
		return fmt.Errorf("outbound order [%s] không ở trạng thái picklisted", params.SalesOrderNumber)
	}
	if outbound.Config != 0 {
		if outbound.Config == 1 {
			return fmt.Errorf("outbound order [%s] không đủ hàng đi pick", params.SalesOrderNumber)
		}
		return fmt.Errorf("outbound order [%s] không đủ điều kiện pick", params.SalesOrderNumber)
	}

	if err := u.ReadyPickOutbound(ctx, &models.ReadyPickOutboundRequest{SalesOrderNumber: params.SalesOrderNumber}); err != nil {
		return err
	}

	if outbound.OutboundOrderType == constants.OUTBOUND_ORDER_TYPE_ORDER {
		return u.PickOrder(ctx, outbound)
	}
	if outbound.OutboundOrderType == constants.OUTBOUND_ORDER_TYPE_INTERNAL_TRANSFER {
		return u.PickIt(ctx, outbound)
	}

	return nil
}

func (u *usecase) PickOrder(ctx context.Context, outbound *outboundModel.OutboundOrder) error {
	return nil
}

func (u *usecase) PickIt(ctx context.Context, outbound *outboundModel.OutboundOrder) error {
	return nil
}
