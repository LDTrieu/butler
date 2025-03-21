package usecase

import (
	"butler/application/domains/pick_pack/models"
	outboundModel "butler/application/domains/services/outbound_order/models"
	outboundOrderExtendModel "butler/application/domains/services/outbound_order_extend/models"
	"context"
	"errors"
)

func (u *usecase) SetOutboundOrderVoucherType(ctx context.Context, params *models.SetOutboundOrderVoucherTypeRequest) error {
	outbound, err := u.outboundOrderSv.GetOne(ctx, &outboundModel.GetRequest{
		SalesOrderNumber: params.SalesOrderNumber,
	})
	if err != nil {
		return err
	}
	if outbound == nil {
		return errors.New("Không tìm thấy đơn hàng : " + params.SalesOrderNumber)
	}

	outboundExtend, err := u.outboundOrderExtendSv.GetOne(ctx, &outboundOrderExtendModel.GetRequest{
		OutboundOrderId: outbound.OutboundOrderId,
	})
	if err != nil {
		return err
	}
	if outboundExtend == nil {
		return errors.New("Không tìm thấy OutboundOrder_Extend của: " + params.SalesOrderNumber)
	}

	_, err = u.outboundOrderExtendSv.Update(ctx, &outboundOrderExtendModel.OutboundOrderExtend{
		OutboundOrderId: outbound.OutboundOrderId,
		VoucherType:     params.VoucherType,
	})
	if err != nil {
		return err
	}

	return nil
}
