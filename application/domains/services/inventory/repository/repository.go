package repository

import (
	"butler/application/domains/services/inventory/models"
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

func (r *repo) GetTableName(warehouseId int64) string {
	return fmt.Sprintf("wms_inventory_%d", warehouseId)
}

func (r *repo) GetById(ctx context.Context, warehouseId int64, id int64) (*models.Inventory, error) {
	record := &models.Inventory{}
	result := r.dbWithContext(ctx).Table(r.GetTableName(warehouseId)).Limit(1).Find(&record, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if record.InventoryId == 0 {
		return nil, nil
	}
	return record, nil
}

func (r *repo) GetOne(ctx context.Context, warehouseId int64, params *models.GetRequest) (*models.Inventory, error) {
	record := &models.Inventory{}
	query := r.dbWithContext(ctx).Table(r.GetTableName(warehouseId))
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

func (r *repo) GetList(ctx context.Context, warehouseId int64, params *models.GetRequest) ([]*models.Inventory, error) {
	records := []*models.Inventory{}
	query := r.dbWithContext(ctx).Table(r.GetTableName(warehouseId))
	query = r.filter(query, params)

	if err := query.Scan(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (r *repo) Update(ctx context.Context, warehouseId int64, obj *models.Inventory) (*models.Inventory, error) {
	if obj.InventoryId == 0 || warehouseId == 0 {
		return nil, fmt.Errorf("id is required")
	}
	result := r.dbWithContext(ctx).Table(r.GetTableName(warehouseId)).Where("inventory_id = ?", obj.InventoryId).Updates(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func (r *repo) UpdateMany(ctx context.Context, warehouseId int64, objs []*models.Inventory) error {
	tx := r.dbWithContext(ctx).Table(r.GetTableName(warehouseId)).Begin()
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
	if len(params.InventoryIds) > 0 {
		query = query.Where("inventory_id in ?", params.InventoryIds)
	}
	return query
}
