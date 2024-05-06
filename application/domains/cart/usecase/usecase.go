package usecase

import (
	"butler/application/domains/cart/models"
	cartModels "butler/application/domains/services/cart/models"
	cartSv "butler/application/domains/services/cart/service"
	initServices "butler/application/domains/services/init"
	packingModels "butler/application/domains/services/packing/models"
	packingSv "butler/application/domains/services/packing/service"
	pgModels "butler/application/domains/services/picking_group/models"
	pgSv "butler/application/domains/services/picking_group/service"
	"butler/constants"
	"context"
	"fmt"
)

type usecase struct {
	cartSv         cartSv.IService
	pickingGroupSv pgSv.IService
	packingSv      packingSv.IService
}

func InitUseCase(
	services *initServices.Services,
) IUseCase {
	return &usecase{
		cartSv:         services.CartService,
		pickingGroupSv: services.PickingGroupService,
		packingSv:      services.PackingService,
	}
}

func (u *usecase) ResetCart(ctx context.Context, params *models.ResetCartRequest) error {
	cart, err := u.cartSv.GetOne(ctx, &cartModels.GetRequest{CartCode: params.CartCode})
	if err != nil {
		return err
	}
	if cart == nil || cart.CartId == 0 {
		return fmt.Errorf("cart with code [%s] not found", params.CartCode)
	}

	pickingGroups, err := u.pickingGroupSv.GetList(ctx, &pgModels.GetRequest{
		CartCode: cart.CartCode,
		StatusIds: []int64{
			constants.PICKING_GROUP_STATUS_NEW, constants.PICKING_GROUP_STATUS_PICKING,
		}},
	)
	if err != nil {
		return err
	}

	packings, err := u.packingSv.GetList(ctx, &packingModels.GetRequest{
		CartCode: cart.CartCode,
		StatusIds: []int64{
			constants.PACKING_STATUS_OPEN, constants.PACKING_STATUS_PACKING,
		},
	})
	if err != nil {
		return err
	}

	_, err = u.cartSv.Update(ctx, &cartModels.Cart{
		CartId: cart.CartId,
		Status: constants.CART_STATUS_AVAILABLE,
	})
	if err != nil {
		return err
	}
	for _, pg := range pickingGroups {
		_, err := u.pickingGroupSv.Update(ctx, &pgModels.PickingGroup{
			PickingGroupId: pg.PickingGroupId,
			Status:         constants.PICKING_GROUP_STATUS_CANCELED,
		})
		if err != nil {
			return err
		}
	}
	for _, packing := range packings {
		_, err := u.packingSv.Update(ctx, &packingModels.Packing{
			PackingId: packing.PackingId,
			StatusId:  constants.PACKING_STATUS_CANCELED,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
