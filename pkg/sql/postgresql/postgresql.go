package postgresql

import (
	"butler/config"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MAX_OPEN_CONNECTIONS     = 60
	CONNECTION_MAX_LIFE_TIME = 120
	MAX_IDLE_CONNECTIONS     = 0
	CONNECTION_MAX_IDLE_TIME = 20
)

func InitConnection(cfg *config.Config) (*gorm.DB, *sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgresql.Host,
		cfg.Postgresql.Port,
		cfg.Postgresql.User,
		cfg.Postgresql.Password,
		cfg.Postgresql.DBName,
		cfg.Postgresql.Sslmode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	dbConfig, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	dbConfig.SetMaxOpenConns(MAX_OPEN_CONNECTIONS)
	dbConfig.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
	dbConfig.SetConnMaxIdleTime(CONNECTION_MAX_IDLE_TIME * time.Second)
	dbConfig.SetConnMaxLifetime(CONNECTION_MAX_LIFE_TIME * time.Second)
	if err = dbConfig.Ping(); err != nil {
		return nil, nil, err
	}
	return db, dbConfig, nil
}
