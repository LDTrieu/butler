package models

import "time"

type OutboundOrderItem struct {
	OutboundOrderItemId int64     `gorm:"primaryKey;column:outbound_order_item_id" json:"outbound_order_item_id"`
	OutboundOrderId     int64     `gorm:"column:outbound_order_id" json:"outbound_order_id"`
	StatusId            int64     `gorm:"column:status_id" json:"status_id"`
	ProductId           int64     `gorm:"column:product_id" json:"product_id"`
	ProductName         string    `gorm:"column:product_name" json:"product_name"`
	Sku                 string    `gorm:"column:sku" json:"sku"`
	Quantity            int64     `gorm:"column:quantity" json:"quantity"`
	InboundShmtItemId   int64     `gorm:"column:inbound_shmt_item_id" json:"inbound_shmt_item_id"`
	ReasonCode          string    `gorm:"column:reason_code" json:"reason_code"`
	ReasonName          string    `gorm:"column:reason_name" json:"reason_name"`
	Description         string    `gorm:"column:description" json:"description"`
	CreatedAt           time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy           int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt           time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy           int64     `gorm:"column:updated_by" json:"updated_by"`
	MappingId           string    `gorm:"column:mapping_id" json:"mapping_id"`
	Price               float64   `gorm:"column:price" json:"price"`
	Amount              float64   `gorm:"column:amount" json:"amount"`
	IsCombo             bool      `gorm:"column:is_combo" json:"is_combo"`
	ComboSku            string    `gorm:"column:combo_sku" json:"combo_sku"`
	ComboQty            int64     `gorm:"column:combo_qty" json:"combo_qty"`
	Config              int       `gorm:"column:config;default:(-)" json:"config"`
}

func (m *OutboundOrderItem) TableName() string {
	return "wms_outbound_order_item"
}
