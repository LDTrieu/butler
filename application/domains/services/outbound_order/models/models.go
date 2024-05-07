package models

import "time"

type OutboundOrder struct {
	OutboundOrderId     int64     `gorm:"primaryKey;column:outbound_order_id" json:"outbound_order_id"`
	OutboundOrderNumber string    `gorm:"column:outbound_order_number" json:"outbound_order_number"`
	StatusId            int64     `gorm:"column:status_id" json:"status_id"`
	OutboundOrderType   string    `gorm:"column:outbound_order_type" json:"outbound_order_type"`
	SalesOrderId        string    `gorm:"column:sales_order_id" json:"sales_order_id"`
	SalesOrderNumber    string    `gorm:"column:sales_order_number" json:"sales_order_number"`
	SalesOrderType      string    `gorm:"column:sales_order_type" json:"sales_order_type"`
	PriorityId          int64     `gorm:"column:priority_id" json:"priority_id"`
	RequireVat          int64     `gorm:"column:require_vat;default:(-)" json:"require_vat"`
	ConfirmedDate       time.Time `gorm:"column:confirmed_date;default:(-)" json:"confirmed_date"`
	ConfirmedBy         string    `gorm:"column:confirmed_by;default:(-)" json:"confirmed_by"`
	DeliveryDate        time.Time `gorm:"column:delivery_date;default:(-)" json:"delivery_date"`
	CompanyId           int64     `gorm:"column:company_id;default:(-)" json:"company_id"`
	CompanyCode         string    `gorm:"column:company_code;default:(-)" json:"company_code"`
	OwnerId             int64     `gorm:"column:owner_id;default:(-)" json:"owner_id"`
	OwnerCode           string    `gorm:"column:owner_code;default:(-)" json:"owner_code"`
	CustomerId          string    `gorm:"column:customer_id;default:(-)" json:"customer_id"`
	CustomerCode        string    `gorm:"column:customer_code;default:(-)" json:"customer_code"`
	CustomerName        string    `gorm:"column:customer_name;default:(-)" json:"customer_name"`
	CustomerPhone       string    `gorm:"column:customer_phone;default:(-)" json:"customer_phone"`
	DeliveryAddress     string    `gorm:"column:delivery_address;default:(-)" json:"delivery_address"`
	DeliveryWard        string    `gorm:"column:delivery_ward;default:(-)" json:"delivery_ward"`
	DeliveryDistrict    string    `gorm:"column:delivery_district;default:(-)" json:"delivery_district"`
	DeliveryProvince    string    `gorm:"column:delivery_province;default:(-)" json:"delivery_province"`
	WarehouseId         int64     `gorm:"column:warehouse_id" json:"warehouse_id"`
	WarehouseCode       string    `gorm:"column:warehouse_code" json:"warehouse_code"`
	InboundShmtId       int64     `gorm:"column:inbound_shmt_id;default:(-)" json:"inbound_shmt_id"`
	InboundShmtNumber   string    `gorm:"column:inbound_shmt_number;default:(-)" json:"inbound_shmt_number"`
	PurchaseOrderId     string    `gorm:"column:purchase_order_id;default:(-)" json:"purchase_order_id"`
	PurchaseOrderNumber string    `gorm:"column:purchase_order_number;default:(-)" json:"purchase_order_number"`
	ReasonCode          string    `gorm:"column:reason_code;default:(-)" json:"reason_code"`
	ReasonName          string    `gorm:"column:reason_name;default:(-)" json:"reason_name"`
	Description         string    `gorm:"column:description;default:(-)" json:"description"`
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy           int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at"`
	UpdatedBy           int64     `gorm:"column:updated_by;default:(-)" json:"updated_by"`
	MappingId           string    `gorm:"column:mapping_id;default:(-)" json:"mapping_id"`
	Amount              float64   `gorm:"column:amount;default:(-)" json:"amount"`
	ShippingFee         float64   `gorm:"column:shipping_fee;default:(-)" json:"shipping_fee"`
	Config              int       `gorm:"column:config;default:(-)" json:"config"`
	ShippingUnitId      int64     `gorm:"column:shipping_unit_id;default:(-)" json:"shipping_unit_id"`
	BoxType             string    `gorm:"column:box_type;default:(-)" json:"box_type"`
	Ctime               int64     `gorm:"column:ctime" json:"ctime"`
	Utime               int64     `gorm:"column:utime" json:"utime"`
}

func (m *OutboundOrder) TableName() string {
	return "wms_outbound_order"
}

type GetRequest struct {
	SalesOrderNumber string
	WarehouseId      int64
	StatusId         int64
	StatusIds        []int64
}
