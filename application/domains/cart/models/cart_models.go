package models

type ResetCartRequest struct {
	CartCode string
}
type ResetCartByUserRequest struct {
	UserId int64
}
