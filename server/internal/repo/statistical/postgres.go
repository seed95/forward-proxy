package statistical

import (
	"context"
	"github.com/seed95/forward-proxy/internal/model"
	"github.com/seed95/forward-proxy/internal/repo"
	"github.com/seed95/forward-proxy/pkg/log"
	"github.com/seed95/forward-proxy/pkg/log/keyval"
	"gorm.io/gorm"
)

type postgresRepo struct {
	db *gorm.DB
}

var _ repo.StatsRepo = (*postgresRepo)(nil)

func New(db *gorm.DB) repo.StatsRepo {
	if err := db.AutoMigrate(&Statistical{}); err != nil {
		log.Panic("postgres auto migrate", keyval.Error(err))
	}

	return &postgresRepo{db: db}
}

func (r *postgresRepo) SaveStat(ctx context.Context, stat model.Statistical) (err error) {
	dbModel := Statistical{
		Url:        stat.Url,
		StatusCode: stat.StatusCode,
		Duration:   stat.DurationResponseTime,
		ReceivedAt: stat.ReceivedAt,
	}
	return r.db.WithContext(ctx).Create(&dbModel).Error
}

func (r *postgresRepo) GetStats(ctx context.Context, from int64) (stats []model.Statistical, err error) {
	dbStats := make([]Statistical, 0)
	r.db.Raw("SELECT url, status_code, duration FROM statisticals WHERE received_at > ?", from).Scan(&dbStats)

	stats = make([]model.Statistical, 0)
	for _, s := range dbStats {
		stat := model.Statistical{
			Url:                  s.Url,
			StatusCode:           s.StatusCode,
			DurationResponseTime: s.Duration,
			ReceivedAt:           s.ReceivedAt,
		}
		stats = append(stats, stat)
	}

	return stats, nil
}
