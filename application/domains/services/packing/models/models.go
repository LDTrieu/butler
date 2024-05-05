package models

import "time"

type Packing struct {
	PackingId           int64     `gorm:"primaryKey;column:packing_id" json:"packing_id"`
	PackingNumber       string    `gorm:"column:packing_number" json:"packing_number"`
	PackingType         string    `gorm:"column:packing_type" json:"packing_type"`
	StatusId            int64     `gorm:"column:status_id" json:"status_id"`
	OutboundOrderId     int64     `gorm:"column:outbound_order_id" json:"outbound_order_id"`
	OutboundOrderNumber string    `gorm:"column:outbound_order_number" json:"outbound_order_number"`
	DeliveryDate        time.Time `gorm:"column:delivery_date;default:(-)" json:"delivery_date"`
	CompanyId           int64     `gorm:"column:company_id" json:"company_id"`
	CompanyCode         string    `gorm:"column:company_code" json:"company_code"`
	OwnerId             int64     `gorm:"column:owner_id" json:"owner_id"`
	OwnerCode           string    `gorm:"column:owner_code" json:"owner_code"`
	WarehouseId         int64     `gorm:"column:warehouse_id" json:"warehouse_id"`
	WarehouseCode       string    `gorm:"column:warehouse_code" json:"warehouse_code"`
	PackerId            int64     `gorm:"column:packer_id" json:"packer_id"`
	PackerCode          string    `gorm:"column:packer_code;default:(-)" json:"packer_code"`
	PackerName          string    `gorm:"column:packer_name;default:(-)" json:"packer_name"`
	PackedDate          time.Time `gorm:"column:packed_date;default:(-)" json:"packed_date"`
	ShippedDate         time.Time `gorm:"column:shipped_date;default:(-)" json:"shipped_date"`
	PrintedDate         time.Time `gorm:"column:printed_date;default:(-)" json:"printed_date"`
	Description         string    `gorm:"column:description;default:(-)" json:"description"`
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy           int64     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at"`
	UpdatedBy           int64     `gorm:"column:updated_by;default:(-)" json:"updated_by"`
	CartId              int64     `gorm:"column:cart_id;default:(-)" json:"cart_id"`
	CartCode            string    `gorm:"column:cart_code;default:(-)" json:"cart_code"`
	PrintedContent      string    `gorm:"column:printed_content;default:(-)" json:"printed_content"`
	PackageStatus       int64     `gorm:"column:package_status;default:(-)" json:"package_status"`
	TrackingNumber      string    `gorm:"column:tracking_number;default:(-)" json:"tracking_number"`
}

func (p *Packing) TableName() string {
	return "wms_packing"
}

type GetRequest struct {
	CartCode    string
	WarehouseId int64
	StatusId    int64
	StatusIds   []int64
}

type GetListResponse struct {
	Records []*PackingResponse
}

type PackingResponse struct {
	CartId      int64
	CartCode    string
	WarehouseId int64
	Status      int64
	CartStt     int64
	CartDesc    string
	CreatedAt   time.Time
	CreatedBy   int64
	UpdatedAt   time.Time
	UpdatedBy   int64
}
