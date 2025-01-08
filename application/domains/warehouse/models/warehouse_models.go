package models

type ShowWarehouseRequest struct {
	WarehouseName string
}

type UpdateConfigWarehouseRequest struct {
	WarehouseId int64
	Config      int
	Operation   string
}
