package models

type ReadyPickOutboundRequest struct {
	SalesOrderNumber string
}

type PickRequest struct {
	SalesOrderNumber string `validate:"required"`
}
