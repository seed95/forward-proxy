package statistical

import (
	"context"
	"fmt"
	"github.com/seed95/forward-proxy/internal"
	"github.com/seed95/forward-proxy/internal/model"
	"github.com/seed95/forward-proxy/internal/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresRepo struct {
	db *gorm.DB
}

var _ repo.StatsRepo = (*postgresRepo)(nil)

func New(config *internal.PostgresConfig) repo.StatsRepo {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Address, config.Username, config.Password,
		config.DatabaseName, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// TODO handle error
		fmt.Println("gorm open", err)
	}

	err = db.AutoMigrate(&Stat{})
	if err != nil {
		// TODO handle error
		fmt.Println("auto migrate", err)
	}

	return &postgresRepo{db: db}
}

func (r *postgresRepo) SaveStat(ctx context.Context, stat model.Statistical) (err error) {
	return r.db.WithContext(ctx).Create(&stat).Error
}

func (r *postgresRepo) GetStats(ctx context.Context, from int64) (stats []model.Statistical, err error) {
	stats = make([]model.Statistical, 0)
	r.db.Model(&Stat{}).Raw("SELECT url, status_code, duration FROM stat WHERE duration > ?", from).Scan(&stats)
	return stats, nil
}
