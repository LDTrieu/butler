package init

import (
	"butler/application/domains/services/picking_item/repository"
	"butler/application/domains/services/picking_item/service"
	"butler/config"

	"gorm.io/gorm"
)

type Init struct {
	Repository repository.IRepository
	Service    service.IService
}

func NewInit(
	db *gorm.DB,
	cfg *config.Config,
) *Init {
	repository := repository.InitRepo(db)
	service := service.InitService(repository)
	return &Init{
		Repository: repository,
		Service:    service,
	}
}
