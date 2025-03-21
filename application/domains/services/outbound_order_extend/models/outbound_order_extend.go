package models

import (
	"time"
)

type OutboundOrderExtend struct {
	OutboundOrderId  int64     `gorm:"primaryKey;column:outbound_order_id" json:"outbound_order_id"`
	PartnerOrderCode string    `gorm:"column:partner_order_code" json:"partner_order_code"`
	CreatedAt        time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	CreatedBy        int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	UpdatedBy        int64     `gorm:"column:updated_by" json:"updated_by"`
	Sla              time.Time `gorm:"column:sla" json:"sla"`
	VoucherType      int64     `gorm:"column:voucher_type;default:(-)" json:"voucher_type"`
}

func (o *OutboundOrderExtend) TableName() string {
	return "wms_outbound_order_extend"
}

type GetRequest struct {
	OutboundOrderId int64
}

type Response struct {
	OutboundOrderId  int64
	PartnerOrderCode string
	CreatedAt        time.Time
	CreatedBy        int64
	UpdatedAt        time.Time
	UpdatedBy        int64
	Sla              time.Time
	VoucherType      int64
}

type ListPaging struct {
	Records []*Response
}
