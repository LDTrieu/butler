package usecase

import (
	initServices "butler/application/domains/services/init"
	"butler/application/lib"
)

type usecase struct {
	lib *lib.Lib
}

func InitUseCase(
	lib *lib.Lib,
	services *initServices.Services,
) IUseCase {
	return &usecase{
		lib: lib,
	}
}
