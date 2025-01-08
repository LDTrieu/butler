package models

import "time"

type Warehouse struct {
	WarehouseId     int64     `gorm:"primaryKey;column:warehouse_id"`
	WarehouseCode   string    `gorm:"column:warehouse_code"`
	WarehouseName   string    `gorm:"column:warehouse_name"`
	Type            string    `gorm:"column:type"`
	Width           int64     `gorm:"column:width"`
	Length          int64     `gorm:"column:length"`
	Height          int64     `gorm:"column:height"`
	Status          int64     `gorm:"column:status"`
	Description     string    `gorm:"column:description"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	CreatedBy       int64     `gorm:"column:created_by"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	UpdatedBy       int64     `gorm:"column:updated_by"`
	MappingId       string    `gorm:"column:mapping_id"`
	LocationId      int64     `gorm:"column:location_id"`
	IsPhysical      int       `gorm:"column:is_physical"`
	IsShare         int       `gorm:"column:is_share"`
	ConfigZone      int64     `gorm:"column:config_zone"`
	IpList          string    `gorm:"column:ip_list"`
	PackedLabelZone string    `gorm:"column:packed_label_zone"`
	Config          int64     `gorm:"column:config"`

	// Custom fields
	EnableLocation bool `gorm:"->"`
}

func (w *Warehouse) TableName() string {
	return "wms_warehouse"
}

type GetRequest struct {
	WarehouseId          int64
	WarehouseName        string
	WarehouseNameSimilar string
}
