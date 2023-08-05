package storage

import (
	"context"
	"fmt"
	"pro-link-api/internal/config"
	"pro-link-api/internal/model"

	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

type (
	IStorage[K ModelType] interface {
		FindById(ctx context.Context, id int) (data K, err error)
	}

	AbstractStorage[K ModelType] struct {
		db        *gorm.DB
		redis     *redis.Client
		tableName string
	}

	Storage struct {
		db    *gorm.DB
		redis *redis.Client
	}

	ModelType interface {
		*model.Account | *model.Education |
			*model.Experience |
			*model.Language |
			*model.Profile |
			*model.Skill |
			*model.WebsiteProfile
	}
)

func New(db *config.DatabaseConfig, redis *config.RedisConfig) *Storage {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", db.Host, db.Username, db.Password, db.Name, db.Port, db.SSLMode, db.Timezone)
	log := gormlog.Default.LogMode(gormlog.Info)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		panic(err)
	}

	database, err := conn.DB()

	if err != nil {
		panic(err)
	}

	err = database.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected...")
	}

	redisClient := NewRedis(redis)

	return &Storage{
		db:    conn,
		redis: redisClient,
	}
}

func NewRedis(redisConfig *config.RedisConfig) *redis.Client {
	dsn := redisConfig.Dsn
	if len(dsn) == 0 {
		dsn = "redis:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	return client
}

func (s *AbstractStorage[K]) FindById(ctx context.Context, id int) (data K, err error) {
	err = s.db.WithContext(ctx).Table(s.tableName).First(&data, id).Error
	return data, err
}

func (s *Storage) GetDB() *gorm.DB {
	return s.db
}

func (s *Storage) GetRedis() *redis.Client {
	return s.redis
}
