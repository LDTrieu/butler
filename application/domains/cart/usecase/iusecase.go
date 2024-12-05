package usecase

import (
	"butler/application/domains/cart/models"
	"context"
)

type IUseCase interface {
	ResetCart(ctx context.Context, params *models.ResetCartRequest) error
	ResetCartByUser(ctx context.Context, params *models.ResetCartByUserRequest) (string, error)
}
