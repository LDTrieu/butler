package usecase

import (
	"butler/application/domains/pick/models"
	"context"
)

type IUseCase interface {
	ReadyPickOutbound(ctx context.Context, params *models.ReadyPickOutboundRequest) error
}
