package helper

import (
	"fmt"
	"github.com/seed95/forward-proxy/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresGormDB(config internal.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Address, config.Username, config.Password,
		config.DatabaseName, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
