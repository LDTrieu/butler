package lib

import (
	"butler/application/commands/cache"
	"butler/config"
	"butler/pkg/sql/mysql"
	"time"

	"bitbucket.org/hasaki-tech/zeus/package/kafka"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Lib struct {
	Db                 *gorm.DB
	Validator          *validator.Validate
	Rdb                *redis.Client
	KafkaPublisherQc   kafka.Publisher
	KafkaPublisherProd kafka.Publisher
	Cache              *cache.Cache
	DiscordBot         config.Butler
}

func InitLib(cfg *config.Config) *Lib {
	db, err := mysql.InitConnection(cfg)
	if err != nil {
		logrus.Fatalf("connect database err: %v", err)
	}

	publisherQc := kafka.NewPublisher(cfg.KafkaQc)
	publisherProd := kafka.NewPublisher(cfg.KafkaProd)

	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.RedisAddr,
		MinIdleConns: cfg.Redis.MinIdleConns,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:     cfg.Redis.Password, // no password set
		DB:           cfg.Redis.DB,       // use default DB
	})

	validate := validator.New(validator.WithRequiredStructEnabled())

	cache := cache.InitCache()

	return &Lib{
		Db:                 db,
		Validator:          validate,
		Rdb:                rdb,
		KafkaPublisherQc:   publisherQc,
		KafkaPublisherProd: publisherProd,
		Cache:              cache,
		DiscordBot:         cfg.DiscordBot.Butler,
	}
}
