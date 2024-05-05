package models

import "time"

type PickingGroup struct {
	PickingGroupId     int64     `gorm:"primaryKey;column:picking_group_id" json:"picking_group_id"`
	PickingGroupNumber string    `gorm:"column:picking_group_number" json:"picking_group_number"`
	PickingGroupType   int64     `gorm:"column:picking_group_type" json:"picking_group_type"`
	ShippingUnitId     int64     `gorm:"column:shipping_unit_id" json:"shipping_unit_id"`
	WarehouseId        int64     `gorm:"column:warehouse_id" json:"warehouse_id"`
	CartCode           string    `gorm:"column:cart_code" json:"cart_code"`
	Status             int64     `gorm:"column:status" json:"status"`
	Description        string    `gorm:"column:description" json:"description"`
	NumGroup           int64     `gorm:"column:num_group" json:"num_group"`
	NumOrder           int64     `gorm:"column:num_order" json:"num_order"`
	GroupId            int64     `gorm:"column:group_id" json:"group_id"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy          int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy          int64     `gorm:"column:updated_by" json:"updated_by"`
}

func (p *PickingGroup) TableName() string {
	return "wms_picking_group"
}

type GetRequest struct {
	CartCode    string
	WarehouseId int64
	StatusId    int64
	StatusIds   []int64
}
