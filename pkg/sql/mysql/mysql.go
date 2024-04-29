package mysql

import (
	"butler/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MAX_OPEN_CONNECTIONS     = 60
	CONNECTION_MAX_LIFE_TIME = 120
	MAX_IDLE_CONNECTIONS     = 0
	CONNECTION_MAX_IDLE_TIME = 20
)

func InitConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Mysql.Username, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbConfig, err := db.DB()
	if err != nil {
		return nil, err
	}
	dbConfig.SetMaxOpenConns(MAX_OPEN_CONNECTIONS)
	dbConfig.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
	dbConfig.SetConnMaxIdleTime(CONNECTION_MAX_IDLE_TIME * time.Second)
	dbConfig.SetConnMaxLifetime(CONNECTION_MAX_LIFE_TIME * time.Second)
	if err = dbConfig.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
