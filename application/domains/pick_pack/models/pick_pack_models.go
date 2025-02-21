package models

import "time"

type AutoPickPackRequest struct {
	LoginRequest
	SalesOrderNumber string
	ShippingUnitId   int64
}

type LoginWmsRequest struct {
	EmailWms    string `json:"email"`
	PasswordWms string `json:"password"`
}

type LoginWmsResponse struct {
	Token string `json:"token"`
	User  struct {
		UserId   int    `json:"user_id"`
		LastName string `json:"last_name"`
		Email    string `json:"email"`
		Status   string `json:"status"`
	} `json:"user"`
	Message string `json:"message"`
}

type LoginDiscordRequest struct {
	LoginDiscord    string      `json:"login"`
	PasswordDiscord string      `json:"password"`
	Undelete        bool        `json:"undelete"`
	LoginSource     interface{} `json:"login_source"`
	GiftCodeSkuId   interface{} `json:"gift_code_sku_id"`
}

type LoginDiscordResponse struct {
	UserId  string `json:"user_id"`
	Token   string `json:"token"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Errors  struct {
		Login struct {
			Errors []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"_errors"`
		} `json:"login"`
	} `json:"errors"`
}

type LoginRequest struct {
	LoginWmsRequest     LoginWmsRequest
	LoginDiscordRequest LoginDiscordRequest
}

type LoginResponse struct {
	LoginWmsResponse     LoginWmsResponse
	LoginDiscordResponse LoginDiscordResponse
}

type WmsOrderPayload struct {
	OrderID        int64   `json:"order_id"`
	OrderNumber    string  `json:"order_number"`
	OrderStatus    string  `json:"order_status"`
	BoxType        string  `json:"box_type"`
	StockID        int     `json:"stock_id"`
	ShipmentId     int64   `json:"shipment_id"`
	ShipmentIds    []int64 `json:"shipment_ids"`
	OrderType      string  `json:"order_type"`
	PackedCode     string  `json:"packed_code"`
	ShippingUnitId int64   `json:"shipping_unit_id"`

	Items []*WmsOrderPayloadItem `json:"items"`

	PickerID int       `json:"picker_id"`
	PickedAt time.Time `json:"picked_at"`
	PackerID int       `json:"packer_id"`
	PackedAt time.Time `json:"packed_at"`

	PartnerCode string    `json:"partner_code"`
	ShipperId   int64     `json:"shipper_id"`
	ShippedAt   time.Time `json:"shipped_at"`
	UpdatedBy   int64     `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WmsOrderPayloadItem struct {
	Sku         string   `json:"sku"`
	ExpDate     []string `json:"exp_date"`
	VoucherCode []string `json:"voucher_code"`
}
