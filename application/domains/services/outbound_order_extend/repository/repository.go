package repository

import (
	"butler/application/domains/services/outbound_order_extend/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func InitRepo(db *gorm.DB) IRepository {
	return &repo{
		DB: db,
	}
}

func (r *repo) dbWithContext(ctx context.Context) *gorm.DB {
	return r.DB.WithContext(ctx)
}

func (r *repo) GetById(ctx context.Context, id int64) (*models.OutboundOrderExtend, error) {
	record := &models.OutboundOrderExtend{}
	result := r.dbWithContext(ctx).Limit(1).Find(&record, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if record.OutboundOrderId == 0 {
		return nil, nil
	}
	return record, nil
}

func (r *repo) GetOne(ctx context.Context, params *models.GetRequest) (*models.OutboundOrderExtend, error) {
	record := &models.OutboundOrderExtend{}
	query := r.dbWithContext(ctx).Model(record)
	query = r.filter(query, params)
	result := query.Limit(1).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return record, nil
}

func (r *repo) GetList(ctx context.Context, params *models.GetRequest) ([]*models.OutboundOrderExtend, error) {
	records := []*models.OutboundOrderExtend{}
	query := r.dbWithContext(ctx).Model(&models.OutboundOrderExtend{})
	query = r.filter(query, params)

	if err := query.Scan(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r *repo) Update(ctx context.Context, obj *models.OutboundOrderExtend) (*models.OutboundOrderExtend, error) {
	if obj.OutboundOrderId == 0 {
		return nil, fmt.Errorf("id is required")
	}
	result := r.dbWithContext(ctx).Updates(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) UpdateMany(ctx context.Context, objs []*models.OutboundOrderExtend) error {
	tx := r.dbWithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, obj := range objs {
		if err := tx.Updates(obj).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) filter(query *gorm.DB, params *models.GetRequest) *gorm.DB {
	if params.OutboundOrderId != 0 {
		query = query.Where("outbound_order_id = ?", params.OutboundOrderId)
	}

	return query
}
