package models

import "time"

type Picking struct {
	PickingId           int64     `gorm:"primaryKey;column:picking_id" json:"picking_id"`
	PickingNumber       string    `gorm:"column:picking_number" json:"picking_number"`
	PickingType         string    `gorm:"column:picking_type" json:"picking_type"`
	StatusId            int64     `gorm:"column:status_id" json:"status_id"`
	OutboundOrderId     int64     `gorm:"column:outbound_order_id" json:"outbound_order_id"`
	OutboundOrderNumber string    `gorm:"column:outbound_order_number" json:"outbound_order_number"`
	PickupDate          time.Time `gorm:"column:pickup_date;default:(-)" json:"pickup_date"`
	CompanyId           int64     `gorm:"column:company_id" json:"company_id"`
	CompanyCode         string    `gorm:"column:company_code" json:"company_code"`
	OwnerId             int64     `gorm:"column:owner_id" json:"owner_id"`
	OwnerCode           string    `gorm:"column:owner_code" json:"owner_code"`
	WarehouseId         int64     `gorm:"column:warehouse_id" json:"warehouse_id"`
	WarehouseCode       string    `gorm:"column:warehouse_code" json:"warehouse_code"`
	PickerId            int64     `gorm:"column:picker_id" json:"picker_id"`
	PickerCode          string    `gorm:"column:picker_code;default:(-)" json:"picker_code"`
	PickerName          string    `gorm:"column:picker_name;default:(-)" json:"picker_name"`
	PickedDate          time.Time `gorm:"column:picked_date;default:(-)" json:"picked_date"`
	PrintedDate         time.Time `gorm:"column:printed_date;default:(-)" json:"printed_date"`
	Description         string    `gorm:"column:description;default:(-)" json:"description"`
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy           int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at"`
	UpdatedBy           int64     `gorm:"column:updated_by;default:(-)" json:"updated_by"`
}

func (t *Picking) TableName() string {
	return "wms_picking"
}

type GetRequest struct {
	OutboundOrderId int64
	WarehouseId     int64
	StatusId        int64
	StatusIds       []int64
}
