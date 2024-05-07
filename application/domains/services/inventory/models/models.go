package models

import "time"

type Inventory struct {
	InventoryId             int64     `gorm:"primaryKey;column:inventory_id" json:"inventory_id"`
	UID                     UID       `gorm:"->" json:"uid"`
	StatusId                int64     `gorm:"column:status_id" json:"status_id"`
	ProductId               int64     `gorm:"->" json:"product_id"`
	Sku                     string    `gorm:"->" json:"sku"`
	SerialNumber            string    `gorm:"->" json:"serial_number"`
	Barcode                 string    `gorm:"->" json:"barcode"`
	Uom                     string    `gorm:"->" json:"uom"`
	BatchNumber             string    `gorm:"->" json:"batch_number"`
	ProductStatusId         int64     `gorm:"->" json:"product_status_id"`
	BrandId                 int64     `gorm:"->" json:"brand_id"`
	VendorId                int64     `gorm:"->" json:"vendor_id"`
	VendorCode              string    `gorm:"->" json:"vendor_code"`
	CompanyId               int64     `gorm:"->" json:"company_id"`
	CompanyCode             string    `gorm:"->" json:"company_code"`
	OwnerId                 int64     `gorm:"->" json:"owner_id"`
	OwnerCode               string    `gorm:"->" json:"owner_code"`
	WarehouseId             int64     `gorm:"->" json:"warehouse_id"`
	WarehouseCode           string    `gorm:"->" json:"warehouse_code"`
	FloorId                 int64     `gorm:"column:floor_id" json:"floor_id"`
	AreaId                  int64     `gorm:"column:area_id" json:"area_id"`
	RackId                  int64     `gorm:"column:rack_id" json:"rack_id"`
	AsileId                 int64     `gorm:"column:asile_id" json:"asile_id"`
	ShelfId                 int64     `gorm:"column:shelf_id" json:"shelf_id"`
	BinId                   int64     `gorm:"column:bin_id" json:"bin_id"`
	LocationDescription     string    `gorm:"column:location_description;default:(-)" json:"location_description"`
	OriginWarehouseId       int64     `gorm:"->" json:"origin_warehouse_id"`
	OriginWarehouseCode     string    `gorm:"->" json:"origin_warehouse_code"`
	PreWarehouseId          int64     `gorm:"->" json:"pre_warehouse_id"`
	PreWarehouseCode        string    `gorm:"->" json:"pre_warehouse_code"`
	PreInventoryStatus      int64     `gorm:"->" json:"pre_inventory_status"`
	InboundShmtId           int64     `gorm:"->" json:"inbound_shmt_id"`
	InboundShmtNumber       string    `gorm:"->" json:"inbound_shmt_number"`
	InboundShmtType         string    `gorm:"->" json:"inbound_shmt_type"`
	InboundShmtItemId       int64     `gorm:"->" json:"inbound_shmt_item_id"`
	PurchaseOrderId         string    `gorm:"->" json:"purchase_order_id"`
	PurchaseOrderNumber     string    `gorm:"->" json:"purchase_order_number"`
	PurchaseOrderType       string    `gorm:"->" json:"purchase_order_type"`
	PurchaseOrderItemId     int64     `gorm:"->" json:"purchase_order_item_id"`
	OriginInboundShmtId     int64     `gorm:"->" json:"origin_inbound_shmt_id"`
	OriginInboundShmtNumber string    `gorm:"->" json:"origin_inbound_shmt_number"`
	OriginInboundShmtItemId int64     `gorm:"->" json:"origin_inbound_shmt_item_id"`
	OutboundOrderId         int64     `gorm:"->" json:"outbound_order_id"`
	OutboundOrderNumber     string    `gorm:"->" json:"outbound_order_number"`
	OutboundOrderType       string    `gorm:"->" json:"outbound_order_type"`
	OutboundOrderItemId     int64     `gorm:"->" json:"outbound_order_item_id"`
	SalesOrderId            string    `gorm:"->" json:"sales_order_id"`
	SalesOrderNumber        string    `gorm:"->" json:"sales_order_number"`
	SalesOrderType          string    `gorm:"->" json:"sales_order_type"`
	SalesOrderItemId        string    `gorm:"->" json:"sales_order_item_id"`
	IsReturned              int64     `gorm:"->" json:"is_returned"`
	Description             string    `gorm:"->" json:"description"`
	WhTemp                  int       `gorm:"->" json:"wh_temp"`
	ComboSku                int64     `gorm:"->" json:"combo_sku"`
	CreatedAt               time.Time `gorm:"->" json:"created_at"`
	CreatedBy               int64     `gorm:"->" json:"created_by"`
	UpdatedAt               time.Time `gorm:"->" json:"updated_at"`
	UpdatedBy               int64     `gorm:"->" json:"updated_by"`
}

func (m *Inventory) TableName() string {
	return "wms_inventory"
}

type GetRequest struct {
	InventoryIds []int64
}
