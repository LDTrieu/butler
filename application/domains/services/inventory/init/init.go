package init

import (
	"butler/application/domains/services/inventory/repository"
	"butler/application/domains/services/inventory/service"
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
