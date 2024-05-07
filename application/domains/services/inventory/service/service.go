package service

import (
	"butler/application/domains/services/inventory/models"
	repo "butler/application/domains/services/inventory/repository"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type service struct {
	repo repo.IRepository
}

func InitService(
	repo repo.IRepository,
) IService {
	return &service{
		repo: repo,
	}
}

func (s *service) GetById(ctx context.Context, warehouseId int64, id int64) (*models.Inventory, error) {
	record, err := s.repo.GetById(ctx, warehouseId, id)
	if err != nil {
		logrus.Errorf("error when get inventory by id: %v", err)
		return nil, fmt.Errorf("error when get inventory by id: %v", err)
	}
	return record, nil
}

func (s *service) GetOne(ctx context.Context, warehouseId int64, params *models.GetRequest) (*models.Inventory, error) {
	record, err := s.repo.GetOne(ctx, warehouseId, params)
	if err != nil {
		logrus.Errorf("error when get inventory: err: %v by params: %#v", err, params)
		return nil, fmt.Errorf("error when get inventory: err: %v by params: %#v", err, params)
	}
	return record, nil
}

func (s *service) GetList(ctx context.Context, warehouseId int64, params *models.GetRequest) ([]*models.Inventory, error) {
	records, err := s.repo.GetList(ctx, warehouseId, params)
	if err != nil {
		logrus.Errorf("Error when get list inventory: %v", err)
		return nil, fmt.Errorf("error when get list inventory: %v", err)
	}

	return records, nil
}

func (s *service) Update(ctx context.Context, warehouseId int64, obj *models.Inventory) (*models.Inventory, error) {
	record, err := s.repo.Update(ctx, warehouseId, obj)
	if err != nil {
		logrus.Errorf("error when update inventory: %v", err)
		return nil, fmt.Errorf("error when update inventory: %v", err)
	}

	return record, nil
}

func (s *service) UpdateMany(ctx context.Context, warehouseId int64, obj []*models.Inventory) error {
	err := s.repo.UpdateMany(ctx, warehouseId, obj)
	if err != nil {
		logrus.Errorf("error when update inventory: %v", err)
		return fmt.Errorf("error when update inventory: %v", err)
	}

	return nil
}
