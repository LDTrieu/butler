package models

type BinLocation struct {
	Id                   int64  `gorm:"primaryKey;column:id"`
	WarehouseId          int64  `gorm:"column:warehouse_id"`
	Zone                 string `gorm:"column:zone"`
	Area                 string `gorm:"column:area"`
	Aisle                string `gorm:"column:aisle"`
	Rack                 string `gorm:"column:rack"`
	Shelf                string `gorm:"column:shelf"`
	Bin                  string `gorm:"column:bin"`
	Code                 string `gorm:"column:code"`
	AllowPicklistedOrder int64  `gorm:"column:allow_picklisted_order"`
	AllowPickOrder       int64  `gorm:"column:allow_pick_order"`
}

func (t *BinLocation) TableName() string {
	return "wms_bin_location"
}

type GetRequest struct {
	WarehouseId      int64
	Code             string
	IsAllowPickOrder string
}
