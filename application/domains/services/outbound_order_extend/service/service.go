package service

import (
	"butler/application/domains/services/outbound_order_extend/models"
	repo "butler/application/domains/services/outbound_order_extend/repository"
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

func (s *service) GetById(ctx context.Context, id int64) (*models.OutboundOrderExtend, error) {
	record, err := s.repo.GetById(ctx, id)
	if err != nil {
		logrus.Errorf("error when get outbound order extend by id: %v", err)
		return nil, fmt.Errorf("error when get packing by id: %v", err)
	}
	return record, nil
}

func (s *service) GetOne(ctx context.Context, params *models.GetRequest) (*models.OutboundOrderExtend, error) {
	record, err := s.repo.GetOne(ctx, params)
	if err != nil {
		logrus.Errorf("error when get outbound order extend: err: %v by params: %#v", err, params)
		return nil, fmt.Errorf("error when get outbound order extend: err: %v by params: %#v", err, params)
	}
	return record, nil
}

func (s *service) GetList(ctx context.Context, params *models.GetRequest) ([]*models.OutboundOrderExtend, error) {
	records, err := s.repo.GetList(ctx, params)
	if err != nil {
		logrus.Errorf("Error when get list outbound order extend: %v", err)
		return nil, fmt.Errorf("error when get list outbound order extend: %v", err)
	}

	return records, nil
}

func (s *service) Update(ctx context.Context, obj *models.OutboundOrderExtend) (*models.OutboundOrderExtend, error) {
	record, err := s.repo.Update(ctx, obj)
	if err != nil {
		logrus.Errorf("error when update outbound order extend: %v", err)
		return nil, fmt.Errorf("error when update outbound order extend: %v", err)
	}

	return record, nil
}
