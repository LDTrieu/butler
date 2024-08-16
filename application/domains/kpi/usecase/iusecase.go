package usecase

import (
	"context"
)

type IUseCase interface {
	CountKpi(ctx context.Context, date string, env string) error
}
