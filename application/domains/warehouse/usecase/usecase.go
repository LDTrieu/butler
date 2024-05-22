package usecase

import (
	initServices "butler/application/domains/services/init"
	whModels "butler/application/domains/services/warehouse/models"
	whSv "butler/application/domains/services/warehouse/service"
	"butler/application/domains/warehouse/models"
	"context"
	"fmt"
	"strings"
)

type usecase struct {
	whSv whSv.IService
}

func InitUseCase(
	services *initServices.Services,
) IUseCase {
	return &usecase{
		whSv: services.WarehouseService,
	}
}

const LOCATION_ID_555 = 12

func (u *usecase) ShowWarehouse(ctx context.Context, params *models.ShowWarehouseRequest) error {
	if params.WarehouseName == "" {
		return fmt.Errorf("warehouse name is required")
	}
	warehouseNameKeyword := strings.ToUpper(params.WarehouseName)

	suggestedWarehouses, err := u.whSv.GetList(ctx, &whModels.GetRequest{
		WarehouseNameSimilar: warehouseNameKeyword,
	})
	if err != nil {
		return err
	}
	if len(suggestedWarehouses) == 0 {
		return fmt.Errorf("không có kho nào có tên giống [%s]", warehouseNameKeyword)
	}
	if len(suggestedWarehouses) > 1 {
		var warehouseNames []string
		for _, wh := range suggestedWarehouses {
			warehouseNames = append(warehouseNames, wh.WarehouseName)
		}
		return fmt.Errorf("có nhiều kho có tên giống [%s], vui lòng nhập đúng tên: \n - %s", params.WarehouseName, strings.Join(warehouseNames, "\n- "))
	}
	warehouse := suggestedWarehouses[0]
	if warehouse.LocationId == LOCATION_ID_555 {
		return fmt.Errorf("kho [%s] đã có thể đi pick ở vị trí kho 555", warehouse.WarehouseName)
	}

	if _, err := u.whSv.Update(ctx, &whModels.Warehouse{
		WarehouseId: warehouse.WarehouseId,
		LocationId:  LOCATION_ID_555,
		Description: fmt.Sprintf("location-%d", warehouse.LocationId),
	}); err != nil {
		return err
	}

	return nil
}
