package service

import (
	"butler/application/domains/services/picking_item/models"
	repo "butler/application/domains/services/picking_item/repository"
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

func (s *service) GetById(ctx context.Context, id int64) (*models.PickingItem, error) {
	record, err := s.repo.GetById(ctx, id)
	if err != nil {
		logrus.Errorf("error when get picking item by id: %v", err)
		return nil, fmt.Errorf("error when get picking item by id: %v", err)
	}
	return record, nil
}

func (s *service) GetOne(ctx context.Context, params *models.GetRequest) (*models.PickingItem, error) {
	record, err := s.repo.GetOne(ctx, params)
	if err != nil {
		logrus.Errorf("error when get picking item: err: %v by params: %#v", err, params)
		return nil, fmt.Errorf("error when get picking item: err: %v by params: %#v", err, params)
	}
	return record, nil
}

func (s *service) GetList(ctx context.Context, params *models.GetRequest) ([]*models.PickingItem, error) {
	records, err := s.repo.GetList(ctx, params)
	if err != nil {
		logrus.Errorf("Error when get list picking item: %v", err)
		return nil, fmt.Errorf("error when get list picking item: %v", err)
	}

	return records, nil
}

func (s *service) Update(ctx context.Context, obj *models.PickingItem) (*models.PickingItem, error) {
	record, err := s.repo.Update(ctx, obj)
	if err != nil {
		logrus.Errorf("error when update picking item: %v", err)
		return nil, fmt.Errorf("error when update picking item: %v", err)
	}

	return record, nil
}
