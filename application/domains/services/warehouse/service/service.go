package service

import (
	"butler/application/domains/services/warehouse/models"
	repo "butler/application/domains/services/warehouse/repository"
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

func (s *service) GetById(ctx context.Context, id int64) (*models.Warehouse, error) {
	record, err := s.repo.GetById(ctx, id)
	if err != nil {
		logrus.Errorf("error when get warehouse by id: %v", err)
		return nil, fmt.Errorf("error when get warehouse by id: %v", err)
	}
	return record, nil
}

func (s *service) GetOne(ctx context.Context, params *models.GetRequest) (*models.Warehouse, error) {
	record, err := s.repo.GetOne(ctx, params)
	if err != nil {
		logrus.Errorf("error when get warehouse: err: %v by params: %#v", err, params)
		return nil, fmt.Errorf("error when get warehouse: err: %v by params: %#v", err, params)
	}
	return record, nil
}

func (s *service) GetList(ctx context.Context, params *models.GetRequest) ([]*models.Warehouse, error) {
	records, err := s.repo.GetList(ctx, params)
	if err != nil {
		logrus.Errorf("Error when get list warehouse: %v", err)
		return nil, fmt.Errorf("error when get list warehouse: %v", err)
	}

	return records, nil
}

func (s *service) Update(ctx context.Context, obj *models.Warehouse) (*models.Warehouse, error) {
	record, err := s.repo.Update(ctx, obj)
	if err != nil {
		logrus.Errorf("error when update warehouse: %v", err)
		return nil, fmt.Errorf("error when update warehouse: %v", err)
	}

	return record, nil
}

func (s *service) UpdateWithMap(ctx context.Context, warehouseId int64, obj map[string]any, specifyCol ...string) error {
	if warehouseId == 0 {
		return fmt.Errorf("warehouse id is required")
	}
	err := s.repo.UpdateWithMap(ctx, warehouseId, obj, specifyCol)
	if err != nil {
		logrus.Errorf("error when update warehouse: %v", err)
		return fmt.Errorf("error when update warehouse: %v", err)
	}

	return nil
}
