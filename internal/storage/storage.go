package storage

import (
	"fmt"
	"pro-link-api/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

type (
	IStorage[K ModelType] interface {
	}

	AbstractStorage[K ModelType] struct {
		db        *gorm.DB
		tableName string
	}

	Storage struct {
		db *gorm.DB
	}

	ModelType interface {
	}
)

func New(db *config.DatabaseConfig) *Storage {
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

	return &Storage{
		db: conn,
	}
}
