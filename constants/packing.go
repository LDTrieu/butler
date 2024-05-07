package constants

const (
	PACKING_STATUS_OPEN int64 = iota + 1 // get from table wms_status
	PACKING_STATUS_CANCELED
	PACKING_STATUS_PACKING
	PACKING_STATUS_PACKED
	PACKING_STATUS_SHIPPED
)