package models

import "time"

type PickingItem struct {
	PickingItemId       int64     `gorm:"primaryKey;column:picking_item_id" json:"picking_item_id"`
	PickingId           int64     `gorm:"column:picking_id" json:"picking_id"`
	StatusId            int64     `gorm:"column:status_id" json:"status_id"`
	OutboundOrderItemId int64     `gorm:"column:outbound_order_item_id" json:"outbound_order_item_id"`
	ComboId             int64     `gorm:"column:combo_id;default:(-)" json:"combo_id"`
	ProductId           int64     `gorm:"column:product_id" json:"product_id"`
	Sku                 string    `gorm:"column:sku" json:"sku"`
	InventoryId         int64     `gorm:"column:inventory_id" json:"inventory_id"`
	Uid                 string    `gorm:"column:uid" json:"uid"`
	PreferredLocation   int64     `gorm:"column:preferred_location;default:(-)" json:"preferred_location"`
	BinId               int64     `gorm:"column:bin_id;default:(-)" json:"bin_id"`
	LocationDescription string    `gorm:"column:location_description;default:(-)" json:"location_description"`
	ExpirationDate      time.Time `gorm:"column:expiration_date;default:(-)" json:"expiration_date"`
	Description         string    `gorm:"column:description;default:(-)" json:"description"`
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy           int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at"`
	UpdatedBy           int64     `gorm:"column:updated_by;default:(-)" json:"updated_by"`
}

func (t *PickingItem) TableName() string {
	return "wms_picking_item"
}

type GetRequest struct {
	PickingId int64
}
